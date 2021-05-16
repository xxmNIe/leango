package main

import (
	"fmt"
	"gorm.io/gorm"
)
import "gorm.io/driver/mysql"
var db *gorm.DB
func main() {

	//1
	dsn:="root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8&parseTime=True&loc=Local"
	db,err :=gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err !=nil {
		fmt.Print(err)
	}

	//2
	//db,err :=

	fmt.Println(db)


}
