package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/role-manager-server/model"
	"github.com/godcong/role-manager-server/util"
	"github.com/json-iterator/go"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
)

// CensorServer ...
var CensorServer = "localhost:7789"

// OrgMediaAdd ...
/**
* @api {post} /v0/org/media 视频添加(OrgMediaAdd)
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
*		    "detail": {
*		        "ID": "5c3ae6907d1ee602d7c619fd",
*		        "CreatedAt": "2019-01-13T15:19:44.5232093+08:00",
*		        "UpdatedAt": "2019-01-13T15:19:44.5232093+08:00",
*		        "DeletedAt": null,
*		        "Version": 1,
*		        "MediaID": "000000000000000000000000",
*		        "RequestKey": "LinEg1ra09YpCbnrCvCP8zNxKzUtXLCZlmp9um13GAJCTdCEcpZ98g1d25xHs1Hu",
*		        "ResultData": [
*		            {
*		                "code": 200,
*		                "data": [
*		                    {
*		                        "code": 200,
*		                        "dataId": "97cbe645-1703-11e9-8353-00155d33ca2d",
*		                        "extras": {},
*		                        "msg": "OK",
*		                        "results": null,
*		                        "taskId": "img4$kaWFQe4A97ejz2Q8O8Al-1q4okG",
*		                        "url": "https://dbipfs.oss-cn-shanghai.aliyuncs.com/2.jpg?Expires=1547450381&OSSAccessKeyId=LTAIeVGE3zRrmiNm&Signature=db%2BxXmgX08y0nSkOrpDhGC%2Fu310%3D"
*		                    }
*		                ],
*		                "msg": "OK",
*		                "requestId": "27277B1D-847F-408B-B56C-0112267440A9"
*		            }
*		        ]
*		    },
*		    "message": "success"
*		}
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
		key := util.GenerateRandomString(64)
		vid := ctx.PostForm("video_object_key")
		wg := &sync.WaitGroup{}
		wg.Add(2)
		go ThreadRequest(wg, nil, "http://127.0.0.1:7789/v0/validate/frame",
			url.Values{
				"name":        []string{vid},
				"url":         []string{"http://127.0.0.1:7788/v0/media/callback"},
				"request_key": []string{key},
			})

		var prd []*model.ResultData
		pic := ctx.PostForm("pic_object_key")
		go ThreadRequest(wg, &prd, "http://127.0.0.1:7789/v0/validate/pic",
			url.Values{"name": []string{pic}})

		mc := model.NewMediaCensor()
		mc.RequestKey = key

		//wait for done
		log.Println("waiting")
		wg.Wait()

		mc.ResultData = []*model.ResultData{
			prd[0],
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
* @api {get} /v0/org/media 视频列表(OrgMediaList)
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
*		        {
*		            "ID": "5c3ad24d725ffbe68c733d43",
*		            "CreatedAt": "2019-01-13T13:53:17.315+08:00",
*		            "UpdatedAt": "2019-01-13T13:53:17.315+08:00",
*		            "DeletedAt": null,
*		            "Version": 1,
*		            "OrganizationID": "000000000000000000000000",
*		            "CensorID": "5c3ad24d725ffbe68c733d42",
*		            "CensorResult": "",
*		            "Block": false,
*		            "VIPFree": "true",
*		            "Photo": "photo1",
*		            "Name": "name1",
*		            "Type": "type1",
*		            "Language": "language1",
*		            "Output3D": "outpu1",
*		            "VR": "vr1",
*		            "Thumb": "thumb1",
*		            "Introduction": "intro1",
*		            "Starring": "star1",
*		            "Director": "dir1",
*		            "Episode": "epis1",
*		            "TotalNumber": "total1",
*		            "IPNSAddress": "ipns",
*		            "IPFSAddress": "ipfs1",
*		            "KEYAddress": "key1",
*		            "Price": "price1",
*		            "PlayType": "play1",
*		            "ExpireDate": "ex1"
*		        },
*		        {
*		            "ID": "5c3ad28f63f6c61f001f09af",
*					...
*		   		}
*		    ],
*		    "message": "success"
*		}
*
* @apiUse Failed
* @apiSampleRequest /v0/org/media
 */
func OrgMediaList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var medias []*model.Media
		var err error
		user := User(ctx)
		media := model.NewMedia()
		role, err := user.Role()
		if err != nil {
			log.Println(err)
			failed(ctx, err.Error())
			return
		}

		if role.Slug == model.SlugGenesis {
			medias, err = media.ALL()
		} else {
			media.OrganizationID = user.OrganizationID
			medias, err = media.FindByOrg()
		}

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
func ThreadRequest(group *sync.WaitGroup, data *[]*model.ResultData, uri string, values url.Values) {
	defer group.Done()
	resp, err := http.PostForm(uri, values)

	if err != nil {
		log.Println(uri, values.Encode(), err.Error())
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	log.Println("resp:", string(bytes))
	if err != nil {
		log.Println(uri, values.Encode(), err.Error())
		return
	}

	if bytes == nil {
		return
	}

	var json JSON
	err = jsoniter.Unmarshal(bytes, &json)
	if err != nil {
		log.Println(uri, values.Encode(), err.Error())
		return
	}

	if json.Detail == nil {
		return
	}

	*data = json.Detail

}
