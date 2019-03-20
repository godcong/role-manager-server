package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/godcong/role-manager-server/config"
	"github.com/godcong/role-manager-server/model"
	"github.com/godcong/role-manager-server/proto"
	"github.com/google/uuid"
	"github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

// OrgMediaUpdate ...
/**
* @api {post} /v0/org/media/:id 视频更新(OrgMediaUpdate)
* @apiName OrgMediaUpdate
* @apiGroup OrgMedia
* @apiVersion  0.0.1
*
* @apiHeader {string} token user token
*
* @apiParam  {string} 	block           下架:true,false
*
* @apiParam  {string}	vip_free        Vip免费
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
* @apiParam  {string}	key_address     key地址
* @apiParam  {string}	price           价格
* @apiParam  {string}	play_type       播放类型(单次,多次)
* @apiParam  {string}	expire_date     过期时间(48H,24H,0H)
* @apiParam  {string}	censor_result   审核(pass)
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
* @apiSuccessExample {json} Success-Response:
*		{
*		}
*
* @apiUse Failed
* @apiSampleRequest /v0/org/media/:id
 */
func OrgMediaUpdate(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		media := model.NewMedia()
		media.ID = model.ID(id)
		e := media.Find()
		if e != nil {
			failed(ctx, e.Error())
			return
		}
		media.Block, _ = strconv.ParseBool(ctx.PostForm("block"))
		media.VIPFree = ctx.DefaultPostForm("vip_free", media.VIPFree)

		media.Name = ctx.DefaultPostForm("name", media.Name)
		media.Type = ctx.DefaultPostForm("type", media.Type)
		media.Language = ctx.DefaultPostForm("language", media.Language)
		media.Output3D = ctx.DefaultPostForm("output_3d", media.Output3D)
		media.VR = ctx.DefaultPostForm("vr", media.VR)
		media.Thumb = ctx.DefaultPostForm("thumb", media.Thumb)
		media.Introduction = ctx.DefaultPostForm("introduction", media.Introduction)
		media.Starring = ctx.DefaultPostForm("starring", media.Starring)
		media.Director = ctx.DefaultPostForm("director", media.Director)
		media.Episode = ctx.DefaultPostForm("episode", media.Episode)
		media.TotalNumber = ctx.DefaultPostForm("total_number", media.TotalNumber)
		//media.IPNSAddress = ctx.DefaultPostForm("ipns_address", media.IPNSAddress)
		//media.IPFSAddress = ctx.DefaultPostForm("ipfs_address", media.IPFSAddress)
		//media.VideoOSSAddress = ctx.DefaultPostForm("video_oss_address", media.VideoOSSAddress)
		//media.PictureOSSAddress = []string{ctx.DefaultPostForm("picture_oss_address", media.PictureOSSAddress[0])}
		//media.Photo = media.PictureOSSAddress[0]
		media.KeyAddress = ctx.DefaultPostForm("key_address", media.KeyAddress)
		media.Price = ctx.DefaultPostForm("price", media.Price)
		media.PlayType = ctx.DefaultPostForm("play_type", media.PlayType)
		media.ExpireDate = ctx.DefaultPostForm("expire_date", media.ExpireDate)
		media.CensorResult = ctx.DefaultPostForm("censor_result", media.CensorResult)
		cfg := config.Config()
		media.KeyAddress = replaceRuleAddress(cfg, media.ID.Hex())
		e = media.Update()
		if e != nil {
			failed(ctx, e.Error())
			return
		}

		if media.Block == false || media.CensorResult == "pass" {
			e = SendToNodeProcessGRPC(cfg, media)
			if e != nil {
				failed(ctx, e.Error())
				return
			}

		}

		success(ctx, media)
	}
}

func replaceRuleAddress(cfg *config.Configure, target string) string {
	uri := strings.Replace(cfg.Manager.KeyAddressRule, ":id", target, -1)
	return cfg.Manager.Host + uri
}

// SendToNodeProcessGRPC ...
func SendToNodeProcessGRPC(cfg *config.Configure, media *model.Media) error {
	node := NodeClient(NewGRPCClient(cfg))
	timeout, _ := context.WithTimeout(context.Background(), time.Second*5)
	response, e := node.RemoteDownload(timeout, &proto.RemoteDownloadRequest{
		ObjectKey: media.VideoOSSAddress,
		KeyURL:    media.KeyAddress,
	})
	log.Info(response, e)
	if e != nil {
		return e
	}
	ipfs := model.NewIPFS()
	ipfs.FileID = response.Detail.ID
	ipfs.MediaID = media.ID
	e = ipfs.Create()
	if e != nil {
		return e
	}
	return nil
}

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
* @apiParam  {string}	key_address     key地址
* @apiParam  {string}	price           价格
* @apiParam  {string}	play_type       播放类型(单次,多次)
* @apiParam  {string}	expire_date     过期时间(48H,24H,0H)
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
		media.VideoOSSAddress = ctx.PostForm("video_oss_address")
		media.PictureOSSAddress = []string{ctx.PostForm("picture_oss_address")}
		media.Photo = media.PictureOSSAddress[0]
		media.KeyAddress = ctx.PostForm("key_address")
		media.Price = ctx.PostForm("price")
		media.PlayType = ctx.PostForm("play_type")
		media.ExpireDate = ctx.PostForm("expire_date")

		media.OrganizationID = user.OrganizationID

		//vid := ctx.PostForm("video_oss_address")
		err := media.Create()
		if err != nil {
			failed(ctx, err.Error())
			return
		}

		reqKey := uuid.New().String()
		var rds []*model.ResultData
		rds = validateMedia(reqKey, media)

		mc := model.NewMediaCensor()
		mc.RequestKey = reqKey
		mc.MediaID = media.ID
		mc.ResultData = rds
		err = mc.Create()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, mc)
	}
}

func validateMedia(key string, media *model.Media) []*model.ResultData {
	cfg := config.Config()
	wg := &sync.WaitGroup{}
	var vrd []*model.ResultData
	var prd []*model.ResultData
	wg.Add(2)
	//if cfg.RequestType == "rest" {
	//	go httpValidate(wg, &vrd, cfg,
	//		url.Values{
	//			"object_key":    []string{media.VideoOSSAddress},
	//			"id":            []string{key},
	//			"validate_type": []string{"frame"},
	//		})
	//
	//	go httpValidate(wg, &prd, cfg,
	//		url.Values{
	//			"object_key":    media.PictureOSSAddress,
	//			"id":            []string{key},
	//			"validate_type": []string{"jpg"},
	//		})
	//
	//} else {
	go tcpValidate(wg, &vrd, cfg, &proto.ValidateRequest{
		ID:           key,
		ObjectKey:    media.VideoOSSAddress,
		ValidateType: proto.CensorValidateType_Frame,
	})

	go tcpValidate(wg, &prd, cfg, &proto.ValidateRequest{
		ID:           key,
		ObjectKey:    media.PictureOSSAddress[0],
		ValidateType: proto.CensorValidateType_JPG,
	})
	//}

	//wait for done
	wg.Wait()

	return []*model.ResultData{prd[0]}
}

func tcpValidate(group *sync.WaitGroup, data *[]*model.ResultData, cfg *config.Configure, req *proto.ValidateRequest) {
	defer group.Done()
	censor := CensorClient(NewGRPCClient(cfg))
	*data = []*model.ResultData{
		{},
	}

	timeout, _ := context.WithTimeout(context.Background(), time.Second*5)
	var rds []*model.ResultData

	if censor != nil {
		censorReply, err := censor.Validate(timeout, req)
		if err != nil {
			return
		}
		err = jsoniter.UnmarshalFromString(censorReply.Detail.Json, &rds)
		if err != nil {
			return
		}
		*data = rds
	}
}

// httpValidate ...
func httpValidate(group *sync.WaitGroup, data *[]*model.ResultData, cfg *config.Configure, values url.Values) {
	defer group.Done()

	*data = []*model.ResultData{
		{},
	}

	//host := fmt.Sprintf("%s%s/%s/%s", cfg.Censor.Addr, cfg.Censor.Port, cfg.Censor.Version, "validate")
	host := ""

	resp, err := http.PostForm(CheckPrefix(host), values)
	log.Println(host, values.Encode(), err.Error())
	if err != nil {
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	log.Println("resp:", string(bytes))
	if err != nil {
		return
	}

	if bytes == nil {
		return
	}

	var json JSON
	err = jsoniter.Unmarshal(bytes, &json)
	if err != nil {
		return
	}

	if json.Detail == nil {
		return
	}

	*data = json.Detail

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
*		            "KeyAddress": "key1",
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

// OrgMediaCensorList ...
/**
* @api {get} /v0/org/media/:id/censor 视频审核列表(群)(OrgMediaCensorList)
* @apiName OrgMediaCensorList
* @apiGroup OrgMediaCensor
* @apiVersion  0.0.1
*
* @apiHeader {string} token user token
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
* @apiSuccessExample {json} Success-Response:
*		{
*		}
*
* @apiUse Failed
* @apiSampleRequest /v0/org/media/:id/censor
 */
func OrgMediaCensorList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		media := model.NewMedia()
		media.ID = model.ID(id)
		err := media.Find()
		if err != nil {
			failed(ctx, err.Error())
			return
		}

		censors, err := media.Censors()
		if err != nil {
			failed(ctx, err.Error())
			return
		}

		success(ctx, censors)
	}
}

// OrgCensorList ...
/**
* @api {get} /v0/org/censor/:id 视频审核列表(单)(OrgCensorList)
* @apiName OrgCensorList
* @apiGroup OrgCensor
* @apiVersion  0.0.1
*
* @apiHeader {string} token user token
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
* @apiSuccessExample {json} Success-Response:
*		{
*		}
*
* @apiUse Failed
* @apiSampleRequest /v0/org/censor/:id
 */
func OrgCensorList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		mc := model.NewMediaCensor()
		mc.ID = model.ID(id)

		err := mc.Find()
		if err != nil {
			failed(ctx, err.Error())
			return
		}

		success(ctx, mc)
	}
}

// OrgMediaCensorUpdate ...
/**
* @api {post} /v0/org/media/:id/censor/:cid 视频审核更新(群)(OrgMediaCensorUpdate)
* @apiName OrgMediaCensorUpdate
* @apiGroup OrgMediaCensor
* @apiVersion  0.0.1
*
* @apiHeader {string} token user token
*
* @apiParam  {string} verify           	验证: 通过(pass),不通过(failed)
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
* @apiSuccessExample {json} Success-Response:
*		{
*		}
*
* @apiUse Failed
* @apiSampleRequest /v0/org/media/:id/censor
 */
func OrgMediaCensorUpdate(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("cid")

		verify := ctx.PostForm("verify")
		mc := model.NewMediaCensor()
		mc.ID = model.ID(id)

		err := mc.Find()
		if err != nil {
			failed(ctx, err.Error())
			return
		}

		mc.Verify = verify
		err = mc.Update()
		if err != nil {
			failed(ctx, err.Error())
			return
		}

		media, err := mc.Media()
		if err != nil {
			log.Println(err)
			failed(ctx, err.Error())
			return
		}
		media.CensorResult = verify
		media.CensorID = mc.ID
		err = media.Update()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		if media.CensorResult == "pass" {
			err = ReleaseIPFS(media)
			if err != nil {
				failed(ctx, err.Error())
				return
			}
		}

		success(ctx, media)
	}
}

// OrgCensorUpdate ...
/**
* @api {post} /v0/org/censor/:id 视频审核更新(单)(OrgCensorUpdate)
* @apiName OrgCensorUpdate
* @apiGroup OrgCensor
* @apiVersion  0.0.1
*
* @apiHeader {string} token user token
*
* @apiParam  {string} verify           	验证: 通过(pass),不通过(failed)
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
* @apiSuccessExample {json} Success-Response:
*		{
*		}
*
* @apiUse Failed
* @apiSampleRequest /v0/org/censor/:id
 */
func OrgCensorUpdate(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		verify := ctx.PostForm("censor_result")
		mc := model.NewMediaCensor()
		mc.ID = model.ID(id)

		err := mc.Find()
		if err != nil {
			failed(ctx, err.Error())
			return
		}

		mc.Verify = verify
		err = mc.Update()
		if err != nil {
			failed(ctx, err.Error())
			return
		}

		media, err := mc.Media()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		media.CensorResult = verify
		media.CensorID = mc.ID
		err = media.Update()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		log.Infof("%+v", media)
		if media.CensorResult == "pass" {
			err = ReleaseIPFS(media)
			if err != nil {
				failed(ctx, err.Error())
				return
			}
		}

		success(ctx, media)
	}
}

// ReleaseIPFS ...
func ReleaseIPFS(media *model.Media) error {
	censor := CensorClient(NewGRPCClient(config.Config()))
	timeout, _ := context.WithTimeout(context.Background(), time.Second*5)
	response, err := censor.Validate(timeout, &proto.ValidateRequest{
		ObjectKey: media.VideoOSSAddress,
	})
	log.Info(response, err)
	if err != nil {
		return err
	}
	//var mp NodeResult
	//err = jsoniter.UnmarshalFromString(response.Detail.Json, &mp)
	//if err != nil {
	//	log.Println(err)
	//	return err
	//}
	ipfs := model.NewIPFS()
	ipfs.FileID = response.Detail.ID
	ipfs.MediaID = media.ID
	err = ipfs.Create()
	if err != nil {
		return err
	}
	return nil
}

// NodeResult ...
type NodeResult struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Detail struct {
		ID string `json:"id"`
	} `json:"detail"`
}
