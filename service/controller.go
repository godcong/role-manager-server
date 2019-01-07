package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/godcong/role-manager-server/model"
	"github.com/godcong/role-manager-server/util"
	"github.com/mongodb/mongo-go-driver/bson"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"net/http"
	"time"
)

const key = ""

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

func addUser(ctx *gin.Context) (*model.User, error) {
	user := model.NewUser()
	user.Name = ctx.PostForm("name")
	user.Username = ctx.PostForm("username")
	user.Email = ctx.PostForm("email")
	user.Mobile = ctx.PostForm("mobile")
	user.IDCardFacade = ctx.PostForm("idCardFacade")
	user.IDCardObverse = ctx.PostForm("idCardObverse")
	user.Organization = ctx.PostForm("organization")
	user.SetPassword(ctx.PostForm("password"))
	err := user.Create()
	return user, err

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

// EncryptJWT ...
func EncryptJWT(key []byte, sub []byte) (string, error) {
	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: key}, (&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		panic(err)
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

	if u.Password != pass {
		return nil, errors.New("password not validated")
	}
	return u, err
}

// ToToken ...
func ToToken(u *model.User) (string, error) {
	t := Token{}
	t.Name = u.Name
	t.OID = u.ID.String()
	//t.Pwd = u.Password //TODO
	t.EffectiveTime = time.Now().Unix() + 3600*24*7

	sub, err := util.MarshalJSON(t)
	if err != nil {
		return "", err
	}

	token, err := EncryptJWT([]byte(key), sub)

	return token, err
}
