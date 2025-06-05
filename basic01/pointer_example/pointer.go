package main

import "fmt"

func main() {
	i, j := 42, 2701

	p:= &i // i의 pointer
	fmt.Println(*p) // p는 i의 pointer가 담겨 있으니까 *를 통해서 해당 포인터의 값 출력 
	*p = 21 // 포인터 p의 값이 21 &p->i , *p = 21 -> i = 21
	fmt.Println(i) // 21

	p = &j
	*p = *p / 37 // *p = 2701 /37 이 결과값을 다시 포인터p에 할당하니까 j도 같이 바뀐다.
	fmt.Println(j) // 73

	fmt.Println("test") // test
	
}
