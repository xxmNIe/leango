package main

import (
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)
/*
	Walk函数 遍历文件夹以及包含的文件
		遍历途中的错误会传读walkfn中的err，交给函数处理
		如果walkfn返回了错误，则遍历会终止并返回错误
 */
func main() {
	err1 :=filepath.Walk("demo", func(path string, info fs.FileInfo, err error) error {
			fmt.Println(path)
			if strings.HasSuffix(path,"html"){

				return errors.New("path err          ------")
			}
			if err!=nil {
				fmt.Println("err occr  pre")
				return err
			}else {
				return nil
			}
		})
	fmt.Println(err1)
}
