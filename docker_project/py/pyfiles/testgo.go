package main

import (
	"fmt"
	"os/exec"
)

func main() {
	c := exec.Command("/bin/bash", "-c", "python3 test.py")
	b, err := c.Output()
	if err != nil {
		fmt.Println("err :", err)
	}
	fmt.Println(string(b))
}
