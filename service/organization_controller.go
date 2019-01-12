package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/role-manager-server/model"
	"github.com/json-iterator/go"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// CensorServer ...
var CensorServer = "localhost:7789"

// OrgMediaAdd ...
/**
* @api {post} /v0/org/media 视频添加
* @apiName OrgMediaAdd
* @apiGroup OrgMedia
* @apiVersion  0.0.1
*
* @apiHeader {string} token user token
*
* @apiParam  {string}	vip_free        Vip免费
* @apiParam  {string}	photo           照片
* @apiParam  {string}	name            名称
* @apiParam  {string}	type            类别
* @apiParam  {string}	language        语言
* @apiParam  {string}	output_3d       3D
* @apiParam  {string}	vr              VR
* @apiParam  {string}	thumb           缩略图
* @apiParam  {string}	introduction    简介
* @apiParam  {string}	starring        主演
* @apiParam  {string}	director        导演
* @apiParam  {string}	episode         集数
* @apiParam  {string}	total_number    总集数
* @apiParam  {string}	ipns_address    ipns地址
* @apiParam  {string}	ipfs_address    ipfs地址
* @apiParam  {string}	key_address     key地址
* @apiParam  {string}	price           价格
* @apiParam  {string}	play_type       播放类型(单次,多次)
* @apiParam  {string}	expire_date     过期时间(48H,24H,0H)
*
* @apiParam  {string}	video_object_key  视频OSS地址
* @apiParam  {string}	pic_object_key    图片OSS地址
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
* @apiSuccessExample {json} Success-Response:
*		{
*		    "code": 0,
*		    "detail": [
*
* @apiUse Failed
* @apiSampleRequest /v0/org/media
 */
func OrgMediaAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := User(ctx)
		media := model.NewMedia()
		media.VIPFree = ctx.PostForm("vip_free")
		media.Photo = ctx.PostForm("photo")
		media.Name = ctx.PostForm("name")
		media.Type = ctx.PostForm("type")
		media.Language = ctx.PostForm("language")
		media.Output3D = ctx.PostForm("output_3d")
		media.VR = ctx.PostForm("vr")
		media.Thumb = ctx.PostForm("thumb")
		media.Introduction = ctx.PostForm("introduction")
		media.Starring = ctx.PostForm("starring")
		media.Director = ctx.PostForm("director")
		media.Episode = ctx.PostForm("episode")
		media.TotalNumber = ctx.PostForm("total_number")
		media.IPNSAddress = ctx.PostForm("ipns_address")
		media.IPFSAddress = ctx.PostForm("ipfs_address")
		media.KEYAddress = ctx.PostForm("key_address")
		media.Price = ctx.PostForm("price")
		media.PlayType = ctx.PostForm("play_type")
		media.ExpireDate = ctx.PostForm("expire_date")

		media.OrganizationID = user.OrganizationID

		vrd := make(chan *model.ResultData)
		vid := ctx.PostForm("video_object_key")
		go ThreadRequest(vrd, "",
			url.Values{"name": []string{vid}})

		prd := make(chan *model.ResultData)
		pic := ctx.PostForm("pic_object_key")
		go ThreadRequest(prd, "",
			url.Values{"name": []string{pic}})

		mc := model.NewMediaCensor()

		select {
		case v := <-vrd:
			if v != nil {
				mc.ResultData = append(mc.ResultData, v)
			}
		case p := <-prd:
			if p != nil {
				mc.ResultData = append(mc.ResultData, p)
			}
			//TODO:
		}

		err := mc.Create()
		if err != nil {
			failed(ctx, err.Error())
			return
		}

		media.CensorID = mc.ID
		err = media.Create()
		if err != nil {
			failed(ctx, err.Error())
			return
		}

		success(ctx, mc)
	}
}

// OrgMediaList ...
/**
* @api {post} /v0/org/media 视频列表
* @apiName OrgMediaList
* @apiGroup OrgMedia
* @apiVersion  0.0.1
*
* @apiHeader {string} token user token
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
* @apiSuccessExample {json} Success-Response:
*		{
*		    "code": 0,
*		    "detail": [
*
* @apiUse Failed
* @apiSampleRequest /v0/org/media
 */
func OrgMediaList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := User(ctx)
		media := model.NewMedia()
		media.OrganizationID = user.OrganizationID
		medias, err := media.FindByOrg()
		if err != nil {
			log.Println(err)
			failed(ctx, err.Error())
			return
		}
		success(ctx, medias)
		return
	}
}

// ThreadRequest ...
func ThreadRequest(data chan<- *model.ResultData, uri string, values url.Values) {
	resp, err := http.PostForm(CensorServer+uri, values)
	if err != nil {
		log.Println(uri, values.Encode(), err.Error())
		data <- nil
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(uri, values.Encode(), err.Error())
		data <- nil
		return
	}

	var json JSON
	err = jsoniter.Unmarshal(bytes, &json)
	if err != nil {
		log.Println(uri, values.Encode(), err.Error())
		data <- nil
		return
	}

	data <- json.Detail
}
