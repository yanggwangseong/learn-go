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

func main(){
	v := Vertex{1, 2}
	v.x = 10
	fmt.Println(v.x) // 10
}
