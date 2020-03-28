package something

import "fmt"

//다른 파일에서 사용할 수 있는 펑션을 만들려면 이름 제일 앞스펠링을 대문자로 만들어야 한다.
func SayBye() {
	fmt.Println("Bye")
}

func SayHello() {
	fmt.Println("Hello")
}
