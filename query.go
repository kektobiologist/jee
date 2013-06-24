package main

import (
	"code.google.com/p/go.net/html"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
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
			a, b := GetStudent(goquery.NewDocumentFromNode(node))
			ch1 <- a
			ch2 <- b
		}
	}
	ch1 <- Student{}
	ch2 <- false
}

func main() {
	num := 1000
	start := 7020001
	ch1 := make(chan Student, num)
	ch2 := make(chan bool, num)
	for i := 0; i < num; i++ {
		go getResult(start+i, ch1, ch2)
		time.Sleep(3 * time.Millisecond)
	}
	for i := 0; i < num; i++ {
		student, ok := <-ch1, <-ch2
		if ok {
			if student.quota > 1 {
				fmt.Println(student.quota, student.plaintext)
			}
		}
	}
}
