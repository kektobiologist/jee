package main

import (
	"code.google.com/p/go.net/html"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"strings"
)

func LoadDoc(page string) *goquery.Document {
	base, _ := os.Getwd()
	if file, e := os.Open(base + "/" + page); e != nil {
		panic(e.Error())
	} else {
		defer file.Close()
		if node, e := html.Parse(file); e != nil {
			panic(e.Error())
		} else {
			return goquery.NewDocumentFromNode(node)
		}
	}
	return nil
}

func IsSelected(str string) bool {
	if strings.Contains(str, "Congratulations!") {
		return true
	}
	return false
}

func IsInvalid(str string) bool {
	if strings.Contains(str, "Roll Number is invalid") {
		return true
	}
	return false
}
func main() {
	doc := LoadDoc("data/gen.html")
	doc2 := LoadDoc("data/invalid.html")

	fmt.Println(IsSelected(doc.Text()), IsInvalid(doc.Text()), IsSelected(doc2.Text()), IsInvalid(doc2.Text()))
}
