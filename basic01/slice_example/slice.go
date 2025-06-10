package main

import "fmt"

// slice
// 배열의 크기가 고정되어 있습니다.
// 반면에 슬라이스는 배열의 요소에 동적으로 크기가 크고 유연한 모습입니다.
// 실제로 슬라이스는 배열보다 훨씬 일반적입니다.
// slice가 동적 배열의 역할을 하는 녀석인것 같다.

// func main(){
// 	primes := [6]int{2, 3, 5, 7, 11, 13}

// 	var s []int = primes[1:4] // [3 5 7]
// 	fmt.Println(s)
// }

// slice는 array의 대한 참조와 같다.
// slice 요소를 수정하면 원본 array 배열도 함께 수정된다.

func main(){
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names) // [John Paul George Ringo]
	
	// slice
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b) // [John Paul] [Paul George]

	b[0] = "XXX"
	fmt.Println(a, b) // [John XXX] [XXX George]
	fmt.Println(names) // [John XXX George Ringo] 원본요소도 바뀜
	
	// slice literals
	// 길이가 없는 배열 리터럴과 같다.
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	// slice deafult
	s := []int{2, 3, 5, 7, 11, 13}
	s = s[1:4]
	fmt.Println(s)

	// slice make
	// bb := make([]int, 0, 5)
	// printSlice("b", bb)

	// cc := bb[:2]
	// printSlice("c", cc)

	// slice 추가요소 append
	var ss []int
	printSlice(ss)

	// append works on nil slices.
	ss = append(ss, 0)
	printSlice(ss)

}

// func printSlice(s string, x []int) {
// 	fmt.Printf("%s len=%d cap=%d %v\n",
// 		s, len(x), cap(x), x)
// }
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
