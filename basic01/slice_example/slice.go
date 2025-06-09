package main

import "fmt"

// slice
// 배열의 크기가 고정되어 있습니다.
// 반면에 슬라이스는 배열의 요소에 동적으로 크기가 크고 유연한 모습입니다.
// 실제로 슬라이스는 배열보다 훨씬 일반적입니다.
// slice가 동적 배열의 역할을 하는 녀석인것 같다.

func main(){
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // [3 5 7]
	fmt.Println(s)
}
