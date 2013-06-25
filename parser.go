package main

import (
	"code.google.com/p/go.net/html"
	"github.com/PuerkitoBio/goquery"
	"os"
	"strconv"
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

func isSelected(str string) bool {
	if strings.Contains(str, "Congratulations!") {
		return true
	}
	return false
}

func isInvalid(str string) bool {
	if strings.Contains(str, "Roll Number is invalid") {
		return true
	}
	return false
}

func GetQuota(str string) Quota {
	for i, v := range quotaMap {
		if strings.Contains(str, v) {
			return i
		}
	}
	return DUNNO
}

func GetStudent(d *goquery.Document, rollno int) (s Student, ok bool) {
	//sanity on document
	// if v := d.Find(".titlehead").Children().Text(); v != "JEE (Advanced) - 2013 Result" {
	// 	return Student{}, false
	// }
	dtext := strings.Trim(d.Text(), " ")
	dfields := strings.Fields(dtext)
	for _, v := range dfields {
		s.Plaintext += v + " "
	}
	s.Plaintext = strings.Trim(s.Plaintext, " ")
	if isInvalid(dtext) {
		return s, false
	}
	ok = true
	s.Rollno = rollno
	s.Region = s.Rollno / 10000
	if !isSelected(dtext) {
		return
	}
	s.Selected = true
	s.Rank, _ = strconv.Atoi(d.Find(".style7").First().Text())
	text, _ := d.Find(".titlehead").First().Parent().Next().Children().Children().First().Html()
	tokens := strings.Split(text, "<br/>")
	nameToks := strings.Fields(tokens[1])
	nameToks = nameToks[2:len(nameToks)]
	for _, v := range nameToks {
		s.Name += v + " "
	}
	s.Name = strings.Trim(s.Name, " ")
	s.Q = GetQuota(dtext)
	return
}

// func main() {
// 	doc := LoadDoc("data/gen.html")
// 	student, _ := GetStudent(doc)
// 	fmt.Println(student)
// 	fmt.Println(regionMap[student.Region])
// }
