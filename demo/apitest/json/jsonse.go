package main

import (
	"encoding/json"
	"os"
)

type info struct {
	Name string
	age int
}
func main() {
	json.NewEncoder(os.Stdout).Encode(&info{
		Name: "agv",
		age : 10,
	})

}
