package main





var stats = make(map[int]Feedback)


func getStats() {
	for fb := range sched2ui {
		stats[fb.id] = fb
	}
}


func main() {


}
