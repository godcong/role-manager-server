package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/godcong/role-manager-server/model"
	"github.com/json-iterator/go"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"log"
	"net/http"
	"runtime"
	"strings"
)

/**
 * @apiDefine Success
 * @apiSuccess {string} msg 返回具体消息
 * @apiSuccess {int} code 返回状态码：【正常：0】，【失败，-1】
 * @apiSuccess {json} [detail]  如正常则返回detail
 */
/**
 * @apiDefine Failed
 * @apiErrorExample {json} Error-Response:
 *     {
 *       "code":-1,
 *       "msg":"error message",
 *     }
 */
const globalKey = ""
const globalSalt = ""

// OrganizationApply ...
/**
* @api {post} /v0/apply 组织申请(OrganizationApply)
* @apiName OrganizationApply
* @apiGroup Default
* @apiVersion  0.0.1
*
* @apiParam  {string} applyName           			商户名称
* @apiParam  {string} applyCode       				社会统一信用代码
* @apiParam  {string} applyContact          		商户联系人
* @apiParam  {string} applyBusinessLicense			营业执照
* @apiParam  {string} applyCorporate				企业法人
* @apiParam  {string} applyPosition         		联系人职位
* @apiParam  {string} applyPhone 					联系人手机号
* @apiParam  {string} applyMailbox					联系人邮箱
* @apiParam  {string} applyCorporateIDCardFacade	法人身份证(正)
* @apiParam  {string} applyCorporateIDCardObverse	法人身份证(反)
* @apiParam  {string} applyIDCardFacade				联系人身份证(正)
* @apiParam  {string} applyIDCardObverse			联系人身份证(反)
* @apiParam  {string} [applyDescription]			描述
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
* @apiSuccessExample {json} Success-Response:
*		{
*		    "code": 0,
*		    "detail": {
*		        "ID": "5c35b06daad2d1c5eb7292bd",
*		        "CreatedAt": "2019-01-09T16:27:25.9038177+08:00",
*		        "UpdatedAt": "2019-01-09T16:27:25.9038177+08:00",
*		        "DeletedAt": null,
*		        "Version": 1,
*		        "IsDefault": false,
*		        "Verify": "application",
*		        "Name": "商户名称",
*		        "Code": "社会统一信用代码",
*		        "Contact": "商户联系人",
*		        "Position": "联系人职位",
*		        "Phone": "联系人手机号",
*		        "Mailbox": "联系人邮箱",
*		        "Description": ""
*		    },
*		    "message": "success"
*		}
*
* @apiUse Failed
* @apiSampleRequest /v0/apply
 */
func OrganizationApply(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		org := model.NewOrganization()
		org.Name = ctx.PostForm("applyName")
		org.Code = ctx.PostForm("applyCode")
		org.Contact = ctx.PostForm("applyContact")
		org.Position = ctx.PostForm("applyPosition")
		org.Phone = ctx.PostForm("applyPhone")
		org.Mailbox = ctx.PostForm("applyMailbox")

		org.Corporate = ctx.PostForm("applyCorporate")
		org.CorporateIDCardFacade = ctx.PostForm("applyCorporateIDCardFacade")
		org.CorporateIDCardObverse = ctx.PostForm("applyCorporateIDCardObverse")
		org.BusinessLicense = ctx.PostForm("applyBusinessLicense")
		org.IDCardFacade = ctx.PostForm("applyIDCardFacade")
		org.IDCardObverse = ctx.PostForm("applyIDCardObverse")
		org.Description = ctx.PostForm("applyDescription")

		org.Verify = model.VerifyApplication //申请中
		err := org.CreateIfNotExist()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, org)
	}
}

// UserRegister ...
/**
* @api {post} /v0/register 用户注册(UserRegister)
* @apiName UserRegister
* @apiGroup Default
* @apiVersion  0.0.1
*
* @apiParam  {string} name            名称
* @apiParam  {string} username        用户名
* @apiParam  {string} email           邮件
* @apiParam  {string} mobile          移动电话
* @apiParam  {string} id_card_facade  身份证(正)
* @apiParam  {string} id_card_obverse 身份证(反)
* @apiParam  {string} organization_id 组织ID
* @apiParam  {string} password        密码
* @apiParam  {string} certificate     证书
* @apiParam  {string} private_key     私钥
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
* @apiSuccessExample {json} Success-Response:
*		{
*		    "code": 0,
*		    "detail": {
*		        "ID": "5c35b8355f262f9b85b765a2",
*		        "CreatedAt": "2019-01-09T17:00:37.6693418+08:00",
*		        "UpdatedAt": "2019-01-09T17:00:37.6693418+08:00",
*		        "DeletedAt": null,
*		        "Version": 1,
*		        "Name": "godcong",
*		        "Username": "godcong",
*		        "Email": "godcong@163.com",
*		        "Mobile": "13058750423",
*		        "IDCardFacade": "",
*		        "IDCardObverse": "",
*		        "OrganizationID": "5c35b06daad2d1c5eb7292bd",
*		        "Password": "123456",
*		        "Certificate": "noacc",
*		        "PrivateKey": "noacc",
*		        "Token": ""
*		    },
*		    "message": "success"
*		}
*
* @apiUse Failed
* @apiSampleRequest /v0/register
 */
func UserRegister(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := addUser(ctx)
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, user)
	}
}

// UserPlay ...
func UserPlay(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		success(ctx, "")
	}
}

// UserPlayList ...
func UserPlayList(s string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		media := model.NewMedia()
		media.ALL()
		success(ctx, "")
	}
}

// MonitorList ...
func MonitorList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// ValidateSlug ...
func ValidateSlug(my *model.User, slug string) error {
	myRole, err := my.Role()
	if err != nil {
		return err
	}

	if myRole.Slug == slug {
		return errors.New("can not add same slug permission")
	}

	switch myRole.Slug {
	case model.SlugGenesis:
		//if slug == model.SlugAdmin || slug == model.SlugMonitor {
		return nil
		//}
	case model.SlugAdmin:
		if slug == model.SlugOrg {
			return nil
		}
	}
	return errors.New("can not add slug between (" + myRole.Slug + "," + slug + ")")
}

// OrgUpload ...
func OrgUpload(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		success(ctx, "")
	}
}

// GenesisGET ...
func GenesisGET(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := model.NewGenesisRole()
		err := model.FindOne(role, bson.M{
			"slug": role.Slug,
		})
		ru := model.NewRoleUser()
		if model.IsExist(ru, bson.M{
			"role_id": role.ID,
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

// User ...
func User(ctx *gin.Context) *model.User {
	if v, b := ctx.Get("user"); b {
		if v0, b := v.(*model.User); b {
			return v0
		}
	}
	return nil
}

// Logger ...
func Logger(ctx *gin.Context) *model.Log {
	if v, b := ctx.Get("logger"); b {
		if v0, b := v.(*model.Log); b {
			return v0
		}
	}
	return nil
}

// LoginPOST ...
/**
* @api {post} /v0/login 用户登录(LoginPOST)
* @apiName Login
* @apiGroup Default
* @apiVersion  0.0.1
*
* @apiParam  {string} username        用户名
* @apiParam  {string} password        密码
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
* @apiSuccessExample {json} Success-Response:
*		{
*		    "code": 0,
*		    "detail": {
*		        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDg0MDIxOTAsImlhdCI6MTU0NzE5MjU5MCwiaXNzIjoiZ29kY29uZyIsImp0aSI6IlFUejRRZ2R4T3Y3YVk3YWEiLCJuYmYiOjE1NDcxOTI1OTAsInN1YiI6IntcIm9pZFwiOlwiNWMzODQ5MDkwNzhkNGQ1YmQyMDE3N2JlXCIsXCJuYW1lXCI6XCJnZW5lc2lzXCIsXCJwd2RcIjpcIlwiLFwiZWZmZWN0aXZlVGltZVwiOjE1NDc3OTczOTB9In0.9GoC8UsbzirPnQCUVtvFQFJO5rLxXM1GNZbg4cWfKQg"
*		    },
*		    "message": "success"
*		}
*
* @apiUse Failed
* @apiSampleRequest /v0/login
 */
func LoginPOST(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := ValidateUser(ctx)
		if err != nil {
			failed(ctx, err.Error())
			return
		}

		token, err := ToToken(user)

		if err != nil {
			log.Println(err)
			failed(ctx, err.Error())
			return
		}

		success(ctx, gin.H{
			"token": token,
		})

	}
}

func updateUser(ctx *gin.Context) (*model.User, error) {
	id := ctx.Param("id")
	user := model.NewUser()
	user.ID = model.ID(id)
	err := user.Find()
	if err != nil {
		return nil, err
	}

	if oid := ctx.PostForm("organization_id"); oid != "" {
		user.OrganizationID = model.ID(oid)
	}
	if password := ctx.PostForm("password"); password != "" {
		user.SetPassword(PWD(password))
	}
	user.Name = ctx.DefaultPostForm("name", user.Name)
	user.Username = ctx.DefaultPostForm("username", user.Username)
	user.Email = ctx.DefaultPostForm("email", user.Email)
	user.Mobile = ctx.DefaultPostForm("mobile", user.Mobile)
	user.IDCardFacade = ctx.DefaultPostForm("id_card_facade", user.IDCardFacade)
	user.IDCardObverse = ctx.DefaultPostForm("id_card_obverse", user.IDCardObverse)
	user.Certificate = ctx.DefaultPostForm("certificate", user.Certificate)
	user.PrivateKey = ctx.DefaultPostForm("private_key", user.PrivateKey)
	err = user.Update()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func addUser(ctx *gin.Context) (*model.User, error) {
	org, err := checkOrganization(ctx)
	if err != nil {
		log.Println(org)
		return nil, errors.New("organization not found")
	}
	password := ctx.PostForm("password")
	user := model.NewUser()
	user.Name = ctx.PostForm("name")
	user.Username = ctx.PostForm("username")
	user.Email = ctx.PostForm("email")
	user.Mobile = ctx.PostForm("mobile")
	user.IDCardFacade = ctx.PostForm("id_card_facade")
	user.IDCardObverse = ctx.PostForm("id_card_obverse")
	user.OrganizationID = org.ID
	user.Certificate = ctx.PostForm("certificate")
	user.PrivateKey = ctx.PostForm("private_key")
	user.SetPassword(PWD(password))

	my := User(ctx)
	slug := "user"
	if my == nil {
		users, _ := org.Users()
		if users == nil {
			slug = model.SlugOrg
		}
	} else {
		slug = ctx.PostForm("slug")
		err = ValidateSlug(my, slug)
		if err != nil {
			return nil, err
		}
	}

	role := model.NewRole()
	role.Slug = slug

	err = addPermission(func() (*model.User, error) {
		err := user.CreateIfNotExist()
		return user, err
	}, role)
	user.Password = password
	if err != nil {
		return nil, err
	}
	return user, nil
}

func addPermission(f func() (*model.User, error), role *model.Role) error {
	var user *model.User
	var err error
	err = model.RelateMaker(func() (modeler model.Modeler, e error) {
		user, err = f()
		if err != nil {
			log.Println(err)
			return nil, errors.New("user is exist")
		}
		return user, nil
	}, func() (modeler model.Modeler, e error) {
		err := role.Find()
		if err != nil {
			return nil, errors.New("role not found")
		}
		return role, nil
	}, func(a, b model.Modeler) error {
		ru := model.NewRoleUser()
		ru.SetUser(user)
		ru.SetRole(role)
		return ru.CreateIfNotExist()
	})
	if err != nil {
		return err
	}

	if role.Slug == "user" {
		return nil
	}

	ps, err := role.Permissions()
	if err != nil || ps == nil {
		return errors.New("permission not found")
	}
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
	if err != nil {
		return err
	}
	return nil
}

func checkOrganization(ctx *gin.Context) (*model.Organization, error) {
	oid := ctx.PostForm("organization_id")
	org := model.NewOrganization()
	org.ID = model.ID(oid)

	err := org.Find()
	if err != nil {
		return nil, err
	}
	if org.Verify == model.VerifyPass {
		return org, nil
	}
	return org, err
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
	//origin := ctx.Request.Header.Get("origin")
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
	//ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, XMLHttpRequest, "+
	//	"Accept-Encoding, X-CSRF-Token, Authorization")
	if ctx.Request.Method == "OPTIONS" {
		ctx.String(200, "ok")
		return
	}

	ctx.Next()
}

// ValidateUser ...
func ValidateUser(ctx *gin.Context) (*model.User, error) {
	user := ctx.PostForm("username")
	pass := ctx.PostForm("password")
	u := model.NewUser()
	err := model.FindOne(u, bson.M{
		"username": user,
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

// LogOutput ...
func LogOutput(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("visit:", ctx.Request.RequestURI)
		log.Println(ctx.Request.URL)
		log.Println(ctx.Request.Host)
		log.Println(ctx.Request.Method)

	}
}

// MediaCallback ...
func MediaCallback(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mc := model.NewMediaCensor()
		mc.RequestKey = ctx.PostForm("request_key")
		err := mc.FindByKey()
		if err != nil {
			log.Println(err)
			failed(ctx, err.Error())
			return
		}
		if mc.ID == primitive.NilObjectID {
			log.Println("no media")
			failed(ctx, "no media")
			return
		}

		msg := ctx.PostForm("message")
		log.Println(msg)
		code := ctx.PostForm("code")
		log.Println(code)
		detail := ctx.PostForm("detail")
		log.Println(detail)
		if detail != "" {
			var rd []*model.ResultData
			err = jsoniter.Unmarshal([]byte(detail), &rd)
			log.Printf("%+v", rd)
			if rd != nil {
				mc.ResultData = []*model.ResultData{
					mc.ResultData[0], //picture
					rd[0],            //video
				}
				err := mc.Update()
				if err != nil {
					log.Println(err)
					failed(ctx, err.Error())
				}
			}
		}

		success(ctx, nil)
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
	logger := Logger(ctx)
	file, line := Caller(2)
	logger.Err = fmt.Sprintf("%s %d:[%s]", file, line, message)
	_ = logger.Update()
	result(ctx, -1, message, nil)
}

func nop(ctx *gin.Context, message string) {
	result(ctx, -2, message, nil)
}

// Caller ...
func Caller(depth int) (file string, line int) {
	_, file, line, ok := runtime.Caller(depth)
	if !ok {
		file = "???"
		line = 0
	}
	return
}
