package router

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"com.py/common"
	"github.com/gin-gonic/gin"
)

const dir string = "pyfiles"

type Pyfile struct {
	Name string `json:"name"`
	File string `josn:"file"`
}

type Res struct {
	Code int
	Msg  string
}

func Savepy(r *gin.Context) {

	res := &Res{}
	res.Code = 200
	defer func() {
		r.JSON(200, gin.H{
			"code": res.Code,
			"msg":  res.Msg,
		})
	}()
	req_parm := &Pyfile{}
	if err := r.ShouldBind(req_parm); err != nil {
		res.Msg = err.Error()
		res.Code = common.Status_parrm_err
		return
	}
	if req_parm.File == "" || req_parm.Name == "" {
		//r.String(http.StatusBadRequest, "parm should not empty", "")
		res.Msg = "parms should not empty"
		res.Code = http.StatusBadRequest
		return
	}
	if err := save(req_parm.Name, &req_parm.File); err != nil {
		//r.String(http.StatusInternalServerError, err.Error())
		res.Code = common.Status_file_err
		res.Msg = err.Error()
		return
	}

	filepath := filepath.Join(dir, req_parm.Name)
	msg, err := invoke(filepath)
	if err != nil {
		res.Code = common.Status_codeExec_err
	}
	res.Msg = msg
	//if res.Code == 3 {
	//	fmt.Println("python exec err:", res.Msg)
	//}
	//fmt.Println(res)
	//r.String(http.StatusOK, res.Msg)
}

func save(name string, file *string) error {

	fmt.Println("in saving")
	if _, err := os.Stat(dir); err != nil {
		if err = os.Mkdir(dir, 0777); err != nil {
			fmt.Println("set pyfile dir err: ", err)
			panic(err)
		}
	}

	path := filepath.Join(dir, name)
	pfile, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModeAppend|os.ModePerm)

	if err != nil {
		panic(err)
	}
	defer pfile.Close()
	_, err = pfile.WriteString(*file)
	if err != nil {
		fmt.Println("write py file err:", err)
		return err
	}

	return nil
}

func invoke(path string) (msg string, err error) {
	fmt.Println("in invoke")
	c := exec.Command("/bin/bash", "-c", "python "+path)
	b, e := c.CombinedOutput()
	//fmt.Println("get result")
	//res.Msg = string(b)
	//fmt.Println("result is:",res.Msg)
	//res.Code = 200
	//fmt.Println("status is",res.Code)
	//if err != nil {
	//	fmt.Println("cmd line err:", err)
	//	return msg,err
	//}
	return string(b), e
}
