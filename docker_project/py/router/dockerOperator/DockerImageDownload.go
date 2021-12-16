package dockerOperator

import (
	"com.py/common"
	"com.py/router"
	"com.py/router/dockerCmd"
	"github.com/gin-gonic/gin"
	"sync"
)

type Image_req_parm struct {
	ImageName   string `json:"ImageName"`
	CreatorName string `json:"CreatorName"`
}

var DownloadMark map[string]int
var DownloadImageMapLock sync.Mutex

func ImageDownload(r *gin.Context) {
	parm := &Image_req_parm{}
	res := &router.Res{}
	res.Code = 200
	defer func() {
		r.JSON(200, gin.H{
			"code": res.Code,
			"msg":  res.Msg,
		})
	}()
	if err := r.ShouldBind(parm); err != nil {
		res.Msg = err.Error()
		res.Code = common.Status_parrm_err
		return
	}
	if parm.ImageName == "" {
		res.Msg = "镜像名称不能为空"
		res.Code = common.Status_ImageNameNull_err
		return
	}
	res.Msg = "镜像准备中"
	dockerCmd.DownloadImage(parm.ImageName, parm.ImageName)
	return
}
