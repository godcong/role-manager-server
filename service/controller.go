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
	"log"
	"net/http"
	"strings"
	"time"
)

const globalKey = ""
const globalSalt = ""

// MonitorList ...
func MonitorList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// OrgVerify ...
func OrgVerify(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// AdminAdd ...
func AdminAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// DashboardAdd ...
func DashboardAdd(s string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// DashboardListGet ...
func DashboardListGet(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func roleIsExist(role *model.Role) bool {
	return model.IsExist(role, bson.M{
		"slug": role.Slug,
	})
}

// ValidateSlug ...
func ValidateSlug(my *model.User, slug string) error {
	myRole, err := my.Role()
	if err != nil {
		return err
	}

	if myRole.Slug == slug {
		return errors.New("can not add same slug")
	}

	switch myRole.Slug {
	case model.SlugGenesis:
		if slug == model.SlugAdmin || slug == model.SlugMonitor {
			return nil
		}
	case model.SlugAdmin:
		if slug == model.SlugOrg {
			return nil
		}
	}
	return errors.New("can not add slug between (" + myRole.Slug + "," + slug + ")")
}

// AddPOST ...
func AddPOST(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		my := User(ctx)
		log.Printf("%+v", *my)
		slug := ctx.PostForm("Slug")
		err := ValidateSlug(my, slug)
		log.Println(err)
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		oid := ctx.PostForm("OID")
		user := model.NewUser()
		user.ID = model.ID(oid)
		err = user.Find()
		log.Println(*user, oid)
		log.Printf("%+v", *user)
		if err != nil {
			log.Println(err)
			failed(ctx, "no_corresponding_user")
			return
		}

		role := model.NewRole()
		role.Slug = slug
		err = role.Find()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		log.Printf("%+v", *role)
		ru := model.NewRoleUser()
		ru.SetRole(role)
		ru.SetUser(user)
		err = ru.CreateIfNotExist()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, "success")
		return
	}
}

// AddOrgPOST ...
func AddOrgPOST(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// AddAdminPOST ...
func AddAdminPOST(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

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
	user.Certificate = ctx.PostForm("certificate")
	user.PrivateKey = ctx.PostForm("private_key")
	user.SetPassword(PWD(ctx.PostForm("password")))

	slug := ctx.PostForm("slug")
	role := model.NewRole()
	role.Slug = slug

	err := model.RelateMaker(func() (modeler model.Modeler, e error) {
		err := user.Create()
		user.Password = ctx.PostForm("password")
		return user, err
	}, func() (modeler model.Modeler, e error) {
		err := role.Find()
		if err != nil {
			return nil, errors.New("role is not found")
		}
		return role, nil
	}, func(a, b model.Modeler) error {
		ru := model.NewRoleUser()
		ru.SetUser(user)
		ru.SetRole(role)
		return ru.CreateIfNotExist()
	})

	ps, err := role.Permissions()
	err = model.RelateMaker(func() (modeler model.Modeler, e error) {
		return user, nil
	}, func() (modeler model.Modeler, e error) {
		return ps[0], nil
	}, func(a, b model.Modeler) error {
		for _, p := range ps {
			pu := model.PermissionUser{}
			pu.SetUser(user)
			pu.SetPermission(p)
			err := pu.CreateIfNotExist()
			if err != nil {
				return err
			}
		}
		return nil
	})
	return user, err
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
