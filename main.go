package main

<<<<<<< HEAD



func main() {


=======
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
>>>>>>> f55b5065742085b2978640078bbe76342b8994e9
}
