package main

import (
	"net/http"
)

//func helloHandle(w http.ResponseWriter, r *http.Request) {
//	key, ok := r.URL.Query()["key"]
//	if ok {
//		ui2sched  <- strings.Join(key, "")
//	}
//	fmt.Fprintf(w, "Hello World, you reached %s \n", r.URL.Path)
//}

func parse() {

}

func gowatch() {
	//http.HandleFunc("/", helloHandle)
	http.ListenAndServe(":8011", nil)

}
