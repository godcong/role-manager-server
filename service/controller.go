package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/godcong/role-manager-server/model"
	"github.com/godcong/role-manager-server/util"
	"github.com/mongodb/mongo-go-driver/bson"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"net/http"
	"strings"
	"time"
)

const globalKey = ""
const globalSalt = ""

// GenesisGET ...
func GenesisGET(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := model.NewGenesis()
		err := model.FindOne(role, bson.M{
			"slug": role.Slug,
		})
		ru := model.NewRoleUser()
		if model.IsExist(ru, bson.M{
			"roleid": role.ID,
		}) {
			failed(ctx, "genesis is exist")
			return
		}

		passwd := "123456"
		user := model.NewUser()
		user.Name = "genesis"
		user.SetPassword(PWD(passwd))
		err = makeUser(user, role)
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, gin.H{
			"Name":     user.Name,
			"Password": passwd,
		})
		return
	}
}

// RegisterPOST ...
func RegisterPOST(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.PostForm("applyName")     //商户名称
		ctx.PostForm("applyCode")     //社会统一信用代码
		ctx.PostForm("applyContact")  //商户联系人
		ctx.PostForm("applyPosition") //联系人职位
		ctx.PostForm("applyPhone")    //联系人手机号
		ctx.PostForm("applyMailbox")  //联系人邮箱
	}
}

// AddUserPOST ...
func AddUserPOST(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := addUser(ctx)
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, gin.H{
			"Name":     user.Name,
			"Password": user.Password,
		})
	}
}

// User ...
func User(ctx *gin.Context) *model.User {
	if v, b := ctx.Get("user"); b {
		if v0, b := v.(*model.User); b {
			return v0
		}
	}
	return nil
}

// LoginPOST ...
func LoginPOST(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := ValidateUser(ctx)
		if err != nil {
			failed(ctx, err.Error())
			return
		}

		token, err := ToToken(user)

		if err != nil {
			failed(ctx, err.Error())
			return
		}

		success(ctx, gin.H{
			"token": token,
		})

	}
}

func addUser(ctx *gin.Context) (*model.User, error) {
	user := model.NewUser()
	user.Name = ctx.PostForm("name")
	user.Username = ctx.PostForm("username")
	user.Email = ctx.PostForm("email")
	user.Mobile = ctx.PostForm("mobile")
	user.IDCardFacade = ctx.PostForm("idCardFacade")
	user.IDCardObverse = ctx.PostForm("idCardObverse")
	user.Organization = ctx.PostForm("organization")
	user.SetPassword(PWD(ctx.PostForm("password")))
	err := user.Create()
	user.Password = ctx.PostForm("password")
	return user, err
	//nextRole := model.NewRole()
	//if my, b := ctx.Get("user"); b {
	//	if myUser, b := my.(*model.User); b {
	//		role, _ := myUser.Role()
	//		switch role.Slug {
	//		case model.SlugGenesis:
	//			nextRole = model.NewAdmin()
	//		case model.SlugAdmin:
	//			nextRole = model.NewOrg()
	//		}
	//	}
	//}
	//err := makeUser(user, nextRole)
	//return user, err

}

func makeUser(user *model.User, role *model.Role) error {
	err := model.Transaction(func() error {
		err := user.Create()
		if err != nil {
			return err
		}
		ru := model.NewRoleUser()
		ru.SetUser(user)
		ru.SetRole(role)
		err = ru.CreateIfNotExist()
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

// AccessControlAllow ...
func AccessControlAllow(ctx *gin.Context) {
	origin := ctx.Request.Header.Get("origin")
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", origin)
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, XMLHttpRequest, "+
		"Accept-Encoding, X-CSRF-Token, Authorization")
	if ctx.Request.Method == "OPTIONS" {
		ctx.String(200, "ok")
		return
	}
	ctx.Next()
}

// DecryptJWT ...
func DecryptJWT(key []byte, token string) (string, error) {
	cl := jwt.Claims{}
	signed, err := jwt.ParseSigned(token)
	if err != nil {
		return "", err
	}

	err = signed.Claims(key, &cl)
	if err != nil {
		return "", err
	}

	expected := jwt.Expected{
		Issuer: "godcong",
		Time:   time.Now(),
	}

	err = cl.Validate(expected)
	if err != nil {
		return "", err
	}

	return cl.Subject, nil
}

// EncryptJWT ...
func EncryptJWT(key []byte, sub []byte) (string, error) {
	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: key}, (&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		return "", nil
	}
	cl := jwt.Claims{
		Subject:   string(sub),
		Issuer:    "godcong",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Expiry:    jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 14)),
		NotBefore: jwt.NewNumericDate(time.Now()),
		ID:        util.GenerateRandomString(16),
	}

	raw, err := jwt.Signed(sig).Claims(cl).CompactSerialize()
	return raw, err
}

// ValidateUser ...
func ValidateUser(ctx *gin.Context) (*model.User, error) {
	user := ctx.PostForm("username")
	pass := ctx.PostForm("password")
	u := model.NewUser()
	err := model.FindOne(u, bson.M{
		"name": user,
	})
	if err != nil {
		failed(ctx, err.Error())
		return nil, err
	}

	if !u.ValidatePassword(PWD(pass)) {
		return nil, errors.New("password not validated")
	}
	return u, err
}

// PWD ...
func PWD(pwd string) string {
	m := hmac.New(sha256.New, []byte(globalKey))
	m.Write([]byte(pwd))
	m.Write([]byte(globalSalt))
	return strings.ToUpper(fmt.Sprintf("%x", m.Sum(nil)))
}

func result(ctx *gin.Context, code int, message string, detail interface{}) {
	h := gin.H{
		"code":    code,
		"message": message,
		"detail":  detail,
	}
	ctx.JSON(http.StatusOK, h)
}

func success(ctx *gin.Context, detail interface{}) {
	result(ctx, 0, "success", detail)
}

func failed(ctx *gin.Context, message string) {
	result(ctx, -1, message, nil)
}
