package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	f, e := ioutil.TempFile("tmp", "tmp-r-*")
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(filepath.Join(f.Name()))

}
