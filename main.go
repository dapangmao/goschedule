package main

import "fmt"


type test struct {
	a int
	b int
}

func (t *test) sum() int {
	return t.a + t.b
}


func main() {
	a := &test{1, 2}
	fmt.Println(a.sum())
}
