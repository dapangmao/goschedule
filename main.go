package main

import "sync"

var stats = make(map[int]Feedback)


func getStats() {
	var mu sync.Mutex
	for fb := range sched2ui {
		mu.Lock()
		stats[fb.id] = fb
		mu.Unlock()
	}
}


func main() {


}
