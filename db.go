package main

// import (
// 	"fmt"
// 	"labix.org/v2/mgo"
// 	"labix.org/v2/mgo/bson"
// )

// type Person struct {
// 	Name  string
// 	Phone string
// }

// func main() {
// 	session, err := mgo.Dial("mongodb://localhost")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer session.Close()

// 	// Optional. Switch the session to a monotonic behavior.
// 	session.SetMode(mgo.Monotonic, true)

// 	c := session.DB("jee").C("students")

// 	// result := Student{}
// 	n, err := c.Find(bson.M{}).Count()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(n)
// 	// fmt.Println("Phone:", result.Phone)
// }
