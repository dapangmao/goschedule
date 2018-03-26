package main

import "sync"

var Entries = struct {
	data map[int]Feedback
	sync.Mutex
}{
	make(map[int]Feedback),
	sync.Mutex{},
}




func getStats() {
	for fb := range sched2ui {
		Entries.Lock()
		Entries.data[fb.id] = fb
		Entries.Unlock()
	}
}


func main() {


}
