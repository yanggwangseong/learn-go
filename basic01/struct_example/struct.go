package main

import "fmt"

// struct
type Vertex struct {
	x int
	y int
}

// func main(){
// 	fmt.Println(Vertex{1, 2})
// }

// func main(){
// 	v := Vertex{1, 2}
// 	v.x = 10
// 	fmt.Println(v.x) // 10
// }

// ****Pointers to structs****
// func main(){
// 	v := Vertex{1, 2}
// 	p := &v
// 	p.x = 1e9
// 	fmt.Println(v) // {1000000000 2}
// }

// ****Struct Literals****
var (
	v1 = Vertex{1, 2}
	v2 = Vertex{x: 1}
	v3 = Vertex{}
	p = &Vertex{1, 2}
)
func main(){
	fmt.Println(v1, p, v2, v3)
}
