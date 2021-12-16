package dockerOperator

import (
	"com.py/common"
	"com.py/router"
	"com.py/router/dockerCmd"
	"github.com/gin-gonic/gin"
)

type image_check_parm struct {
	ownImage dockerCmd.Owner
}

func ImageDownloadCheck(r *gin.Context) {
	parm := &image_check_parm{}
	res := &router.Res{}
	res.Code = 200
	defer func() {
		r.JSON(200, gin.H{
			"code": res.Code,
			"msg":  res.Msg,
		})
	}()
	var code int
	var msg string
	for code == common.Status_ImageDownloading {
		msg, code = dockerCmd.CheckDownloading(parm.ownImage)
	}
	res.Code = code
	res.Msg = msg
}
