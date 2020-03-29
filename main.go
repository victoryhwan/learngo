package main

import (
	"fmt"
	"strings"
)

type person struct {
	name    string
	age     int
	favFood []string
}

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

	// result := superAdd(1, 2, 3, 4, 5, 6)
	// fmt.Println(result)

	// fmt.Println(canIDrink(16))

	// fmt.Println(canIDrink_switch(16))
	// fmt.Println(canIDrink_switch2(15))

	// pointer()

	// arrayAndSlices()

	// map_ex()

	struct_ex()
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

func canIDrink(age int) bool {
	if korAge := age + 2; korAge < 18 {
		return false
	}
	return true
}

func canIDrink_switch(age int) bool {
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age < 50:
		return false
	}
	return false
}

func canIDrink_switch2(age int) bool {
	switch korAge := age + 2; korAge {
	case 1:
		fmt.Println("1")
		return false
	case 18:
		fmt.Println("18")
		return true
	case 50:
		fmt.Println("50")
		return false
	}
	return false
}

func pointer() {
	a := 2
	b := &a

	fmt.Println(a, b)  // b는 a의 메모리 주소를 보고있다
	fmt.Println(a, *b) // b는 a의 메모리 주소의 값을 보고있다.

	*b = 20 //b를 통해 a를 변경시킬 수 있다.
	fmt.Println(a)
}

func arrayAndSlices() {
	names := [5]string{"hwan", "nico", "dal"}
	names[3] = "alala"
	names[4] = "alala"

	fmt.Println(names)

	sliceArr := []string{"hwan", "nico", "dal"} //배열 크기를 선언하지 않고 사용하기
	sliceArr = append(sliceArr, "alal")
	sliceArr = append(sliceArr, "alal")
	sliceArr = append(sliceArr, "alal")
	sliceArr = append(sliceArr, "alal")

	fmt.Println(sliceArr)
}

func map_ex() {
	nico := map[string]string{"name": "hwan", "age": "18"} //map[key타입]value타입

	fmt.Println(nico)

	for key, _ := range nico {
		fmt.Println(key)
		fmt.Println(nico[key])
	}
}

func struct_ex() {
	favFood := []string{"boolgogi", "zun"}
	hwan := person{"hwan", 32, favFood}
	nico := person{name: "hwan", age: 32, favFood: favFood}

	fmt.Println(hwan)
	fmt.Println(hwan.age)

	fmt.Println(nico)
}
