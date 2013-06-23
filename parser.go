package main

import (
	"code.google.com/p/go-html-transform/h5"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	base, er := os.Getwd()
	content, er := ioutil.ReadFile(base + "/data/source1.html")
	if er != nil {
		fmt.Println(er)
		return
	}
	fmt.Println(string(content))
}
