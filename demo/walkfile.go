package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func walkpj(path string){
	root, err := os.Stat(path)
	if err != nil {
		fmt.Println("walkpj : ",err)
	}
	if !root.IsDir() {
		fmt.Println("root not dir")
	}

	subfiles, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("open dir err :",err)
	}

	//basedir := path
	for _,f := range subfiles {
		fname := f.Name()
		if !f.IsDir() || strings.HasPrefix(fname,"."){
			continue
		}

		logfile(filepath.Dir(path)+"/copy",fname,path+"/"+fname)
	}

}

func logfile(savepath,savename,dirpath string){
	fullfime := fmt.Sprintf("%s/%s.txt",savepath,savename)
	desfile, err := os.OpenFile(fullfime,os.O_APPEND|os.O_RDWR|os.O_TRUNC,0660)
	if err != nil {
		fmt.Printf("fullfile  %s no exit,will cteate\n",fullfime)
		desfile, err = os.Create(fullfime)
		if err != nil {
			fmt.Println("create fullfile err: ",err)
		}
	}
	writer := bufio.NewWriter(desfile)
	defer desfile.Close()

	filepath.Walk(dirpath, func(path string, info fs.FileInfo, err error) error {

		if err !=nil && err == filepath.SkipDir  {
			fmt.Printf( "walk file: %s err1 %s\n",path,err.Error())
			os.Remove(fullfime)
			return err
		}
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			fmt.Printf("open file:%s err: %s\n",file,err.Error())
			return err
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for  scanner.Scan() {
			line := scanner.Text()
			_, err := fmt.Fprintln(writer, line)
			if err != nil {
				fmt.Printf("wtrire file : %s err:%s\n",path,err.Error())
			}
		}
		fmt.Fprintln(writer,path)
		fmt.Fprintln(writer,"--------------------------------------")
		//fmt.Println(path)
		return nil
	})
}


func main() {
	walkpj("/home/xxm/leango")


}