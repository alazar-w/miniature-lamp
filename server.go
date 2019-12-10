// package main

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type MyFloat float64

// //can not define new methods on non-local type

// func (f MyFloat) Abs() float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }

// func main() {
// 	f := MyFloat(-math.Sqrt2)
// 	fmt.Println(f.Abs())
// }

// // methods

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Vertex struct {
// 	X, Y float64
// }

// func (v Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// // func (v Vertex) Scale(f float64) {
// // 	v.X = v.X * f
// // 	v.Y = v.Y * f
// // }

// func (v *Vertex) Scale(f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// func main() {
// 	v := Vertex{3, 4}
// 	v.Scale(10)
// 	fmt.Println(v)
// }

//pointer reciver

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Abser interface {
// 	Abs() float64
// }

// func main() {
// 	var a Abser
// 	f := MyFloat(-math.Sqrt2)
// 	v := Vertex{3, 4}

// 	a = f  // a MyFloat implements Abser
// 	a = &v // a *Vertex implements Abser
// 	//a = v  // a vertex does not implement abser

// 	fmt.Println(a.Abs())
// }

// type MyFloat float64

// func (f MyFloat) Abs() float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }

// type Vertex struct{ X, Y float64 }

// func (v *Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// //interface

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func say(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

// func main() {
// 	go say("world")
// 	say("hello")
// }

// go routin

// package main

// import "fmt"

// func sum(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	c <- sum // send sum to c
// }

// func main() {
// 	s := []int{7, 2, 8, -9, 4, 0}

// 	c := make(chan int)
// 	go sum(s[:len(s)/2], c)
// 	go sum(s[len(s)/2:], c)
// 	x, y := <-c, <-c // receive from c

// 	fmt.Println(x, y, x+y)
// }

// // channels

// package main

// import "fmt"

// func fibonacci(c, quit chan int) {
// 	x, y := 0, 1
// 	for {
// 		select {
// 		case c <- x:
// 			x, y = y, x+y
// 		case <-quit:
// 			fmt.Println("quit")
// 			return
// 		}
// 	}
// }

// func main() {
// 	c := make(chan int)
// 	quit := make(chan int)
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			fmt.Println(<-c)
// 		}
// 		quit <- 0
// 	}()
// 	fibonacci(c, quit)
// }

// // select

// package main

// import "net/http"

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("<h1>hello world</h1>"))
// }

// func contactHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("<h1>hii</h1>"))
// }

// func main() {
// 	// mux := http.NewServeMux()
// 	// mux.HandleFunc("/", indexHandler)
// 	// mux.HandleFunc("/contact", contactHandler)
// 	// http.ListenAndServe(":8080", mux)

// 	//alternatively
// 	http.HandleFunc("/index", indexHandler)
// 	http.HandleFunc("/contact", contactHandler)
// 	http.ListenAndServe(":8080", nil)
// 	//it defaults to DefaultServeMux.
// }

// ------ 01 start server

// package main

// import (
// 	"html/template"
// 	"net/http"
// )

// var templ = template.Must(template.ParseFiles("index.html"))

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	templ.Execute(w, nil)
// }

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", indexHandler)
// 	http.ListenAndServe(":8080", mux)
// }

// --------02 using template

// import (
// 	"html/template"
// 	"net/http"
// )

// // Course struct
// type Course struct {
// 	Title, Code, Description string
// 	ECTS                     int
// }

// var templ = template.Must(template.ParseFiles("class.html"))

// func classHandler(w http.ResponseWriter, r *http.Request) {
// 	crs := Course{"Web programing", "ITSE-3182", "Lorem Ipsum", 7}
// 	templ.Execute(w, crs)
// }

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/class", classHandler)
// 	http.ListenAndServe(":8080", mux)
// }

// // --------03 passing data to template

// package main

// import (
// 	"html/template"
// 	"net/http"
// )

// var templ = template.Must(template.ParseFiles("indexstyle.html"))

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	templ.Execute(w, nil)
// }

// func main() {
// 	mux := http.NewServeMux()
// 	fs := http.FileServer(http.Dir("assets"))
// 	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
// 	mux.HandleFunc("/", indexHandler)
// 	http.ListenAndServe(":8080", mux)
// }

// // ----- 04 serving static files

package main

import (
	"html/template"
	"net/http"
)

var templ = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templ.Execute(w, nil)
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", mux)
}

// ----- 05 using bootstrap
