package main

import "fmt"

/**
* var
**/
// var i, j int = 1, 2 선언과 초기화를 짧게 할 수 있다.

func main() {
	ab := 1; cd :=2
	var c, python, java = true, false, "no!"
	fmt.Println(ab, cd, c, python, java) // 1 2 false false false

	// zero values
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// const
	const golang bool = true
	
}

