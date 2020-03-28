package main

import (
	"fmt"
	"strings"
)

func main() {
	// const cName string = "chwan" //수정불가
	// var name string = "hwan"     // 수정가능
	// fmt.Println(name)
	// fmt.Println("Hello world")
	// something.SayHello()
	// something.SayBye()

	// fmt.Println(multiply(2, 2))

	// totalLength, upperName := lenAndUpper(name) // := 를 사용하면 변수 type 이 자동으로 설정된다
	// fmt.Println(totalLength, upperName)

	// totalLength2, upperName2 := lenAndUpper2(name) // := 를 사용하면 변수 type 이 자동으로 설정된다
	// fmt.Println(totalLength2, upperName2)

	// repeatMe("hwan", "lee", "dal", "hwan", "lee", "dal")

	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}

func multiply(a int, b int) int {
	return a * b
}

//리턴값이 여러개인 펑션
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

//리턴값이 여러개인 펑션
//리턴할 변수를 선언하기
func lenAndUpper2(name string) (length int, uppercase string) {
	fmt.Println("start lenAndUpper2")
	defer fmt.Println("I'm done") //해당 소스는 lenAndUpper2가 다 끝나고 리턴하기 전에 실행된다. 오호
	length = len(name)
	uppercase = strings.ToUpper(name)
	fmt.Println("end lenAndUpper2")
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

//go의 for문
func superAdd(numbers ...int) int {

	fmt.Println("case1")
	//case1
	for index, number := range numbers {
		fmt.Println(index, number)
	}

	fmt.Println("case2")
	//case2
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}

	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}
