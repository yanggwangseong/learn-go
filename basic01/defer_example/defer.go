package main

import "fmt"

func main() {
	//fmt.Println("counting")

	// defer
	/*
	* defer 키워드를 사용하면 주변 함수가 반환될때까지 실행 되지 않는 성질을 가지고 있는것 같다.
	*/
	// defer fmt.Println("world")

	// fmt.Println("hello1")
	// fmt.Println("hello2")
	// fmt.Println("hello3")
	// fmt.Println("hello4")
	/*
	* 출력 결과
	*	counting
	*	hello1
	*	hello2
	*	hello3
	*	hello4
	*	done
	*	world --> 다른 함수들이 다 출력되고 난 후 출력된다.
	*/

	//fmt.Println("done")
	stackingDefer()
}

func stackingDefer() {
	fmt.Println("counting")
	for i := range 10 {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
