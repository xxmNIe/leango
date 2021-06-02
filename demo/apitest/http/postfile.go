package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {

	//f,err := os.Create("demo/")
	//if err!=nil {
	//	fmt.Println("111111   ",err)
	//	return
	//}
	//s,err := f.Stat()
	//if err!=nil {
	//	fmt.Println("111111   ",err)
	//	return
	//}
	//
	//fmt.Println(s.Name())
	r ,err := postfilefunc("demo/apitest/http/postfile.txt","abc.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r.Header.Get("Content-Type"))


}

func postfilefunc(filepath,filename string)( *http.Request ,error){
	srcFile,err := os.Open(filepath)
	if err !=nil {
		return nil,err
	}
	defer srcFile.Close()

	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	fileWriter, err := writer.CreateFormFile("uploadFile", filename)
	if err != nil{
		return nil ,err
	}

	if _,err = io.Copy(fileWriter,srcFile);err !=nil {
		return nil,err
	}
	contentType := writer.FormDataContentType()
	writer.Close()
	r := &http.Request{}
	//
	r.Header = http.Header{}
	r.Header.Set("Content-Type",contentType)
	r.Body = ioutil.NopCloser(buf)
	return r,nil

}