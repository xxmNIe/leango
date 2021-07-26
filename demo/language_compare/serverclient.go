package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)
var  basepath string ="./"
var  fileend string =".upload"
func main() {
	http.HandleFunc("/upload", uploadHanlder)
	http.HandleFunc("/get", getfile)
	http.ListenAndServe(":8080", nil)
}
func uploadHanlder(w http.ResponseWriter, r *http.Request) {
	reader ,err := r.MultipartReader()
	if err !=nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for {
		part ,err :=reader.NextPart()
		if err == io.EOF {
			break
		}
		fmt.Printf("FileName=[%s]", part.FileName())
		if part.FileName() == "" {  // this is FormData
			data, _ := ioutil.ReadAll(part)
			fmt.Printf("FormData=[%s]\n", string(data))
		} else {    // This is FileData
			dst, _ := os.Create(basepath + part.FileName() + fileend)
			defer dst.Close()
			io.Copy(dst, part)
		}
	}
}

func getfile(w http.ResponseWriter, r *http.Request){
	fileName := r.URL.Query().Get("file")
	fullpath := basepath+fileName+fileend
	file ,err := os.Open(fullpath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	fileNames := url.QueryEscape(file.Name())
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Content-Disposition", "attachment; filename=\""+fileNames+"\"")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Write(content)
	}
}