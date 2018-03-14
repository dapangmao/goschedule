package main

var ui2sched = make(chan command)
var sched2ui = make(chan ErrorMsg)


//func helloHandle(w http.ResponseWriter, r *http.Request) {
//	key, ok := r.URL.Query()["key"]
//	if ok {
//		ui2sched  <- strings.Join(key, "")
//	}
//	fmt.Fprintf(w, "Hello World, you reached %s \n", r.URL.Path)
//}

//
//func main() {
//
//
//	//http.HandleFunc("/", helloHandle)
//	http.ListenAndServe(":8011", nil)
//
//}





func main() {
	crawl()
}