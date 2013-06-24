package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"net/url"
// 	"strconv"
// 	"time"
// )

// // 7023233
// // 7024270

// func getResult(rollno int, ch chan []byte) {
// 	if resp, err := http.PostForm("http://jee.iitd.ac.in/resultstatus.php",
// 		url.Values{"regno": {strconv.Itoa(rollno)}, "submit": {"Submit"}}); err == nil {
// 		defer resp.Body.Close()
// 		body, _ := ioutil.ReadAll(resp.Body)
// 		ch <- body
// 	} else {
// 		ch <- nil
// 	}
// }

// func main() {
// 	ch := make(chan []byte, 1)
// 	for i := 0; i < 1; i++ {
// 		go getResult(7023233+i, ch)
// 		time.Sleep(3 * time.Millisecond)
// 	}
// 	for i := 0; i < cap(ch); i++ {
// 		// n, _ := strconv.ParseInt(string(<-ch), 10, 32)
// 		// if int(n-7023233) != i {
// 		// 	fmt.Println(n-7023233, i)
// 		// }
//     fmt.Println(string(<-ch))
// 		// fmt.Println(n-7023233, i)
// 		// strconv.ParseInt(s, base, bitSize)
// 	}
// }
