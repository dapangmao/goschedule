package main

import (
	"net/http"
	"fmt"
	"time"
	"strings"
)


type command struct {
	j *Job
	cmd string
}

var ui2sched = make(chan command)
var sched2ui = make(chan ErrorMsg)


//func helloHandle(w http.ResponseWriter, r *http.Request) {
//	key, ok := r.URL.Query()["key"]
//	if ok {
//		ui2sched  <- strings.Join(key, "")
//	}
//	fmt.Fprintf(w, "Hello World, you reached %s \n", r.URL.Path)
//}


func main() {

	//go func(){
	//	for x := range ui2sched  {
	//		tick := time.NewTicker(2 * time.Second)
	//		go func(x string, ticker *time.Ticker){
	//			for range ticker.C {
	//				fmt.Println( x)
	//			}
	//		}(x, tick)
	//	}
	//}()
	//
	//http.HandleFunc("/", helloHandle)
	http.ListenAndServe(":8011", nil)
}