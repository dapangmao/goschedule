package main





var stats = make(map[int]Feedback)


func getStats() {
	for x := range sched2ui {
		stats[x.id] = x
	}
}


func main() {


}
