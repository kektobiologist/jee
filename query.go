package main

import (
	"code.google.com/p/go.net/html"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"labix.org/v2/mgo"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// 7023233
// 7024270

func getResult(rollno int, ch1 chan Student, ch2 chan bool) {
	if resp, err := http.PostForm("http://jee.iitd.ac.in/resultstatus.php",
		url.Values{"regno": {strconv.Itoa(rollno)}, "submit": {"Submit"}}); err == nil {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		if node, e := html.Parse(strings.NewReader(string(body))); e != nil {
			fmt.Println(e)
			ch1 <- Student{}
			ch2 <- false
		} else {
			a, b := GetStudent(goquery.NewDocumentFromNode(node), rollno)
			ch1 <- a
			ch2 <- b
		}
	}
	ch1 <- Student{}
	ch2 <- false
}

var genFunc func() int = (func() func() int {
	codelist := make([]int, 0, len(regionMap))
	for key, _ := range regionMap {
		codelist = append(codelist, key*10000)
	}
	i := 0
	n4 := 1
	num := codelist[0] + 1
	abort := false
	return func() int {
		defer (func() {
			n4 += 1
			if n4 > 9999 {
				i += 1
				n4 = 1
				if i >= len(codelist) {
					fmt.Println("exhausted rolno!!")
					abort = true
				}
			}
			if abort {
				num = -1
			} else {
				num = codelist[i] + n4
			}
		})()
		return num
	}
})()

func main() {
	// fmt.Println(genFunc(), genFunc(), genFunc())
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("jee").C("students")
	num := 959903
	ch1 := make(chan Student, num)
	ch2 := make(chan bool, num)
	for i := 0; i < num; i++ {
		go getResult(genFunc(), ch1, ch2)
		time.Sleep(3 * time.Millisecond)
	}
	fmt.Println("started all threads!")
	for i := 0; i < num; i++ {
		student, ok := <-ch1, <-ch2
		if ok {
			// fmt.Println("got 1")
			// if student.quota > 3 && student.selected {
			// 	fmt.Println(student.Q, student.plaintext)
			err = c.Insert(&student)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
