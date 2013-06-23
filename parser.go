package main

import (
	"code.google.com/p/go.net/html"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	base, er := os.Getwd()
	content, er := ioutil.ReadFile(base + "/data/source1.html")
	if er != nil {
		fmt.Println(er)
		return
	}
	rawdoc, _ := html.Parse(strings.NewReader(string(content)))
	doc := goquery.NewDocumentFromNode(rawdoc)
	fmt.Println(doc.Text())
}
