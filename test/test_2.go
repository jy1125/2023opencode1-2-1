package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//reader := bufio.NewReader(os.Stdin) 이해하기!
	
	reader := bufio.NewReader(os.Stdin) //읽어오는? 입력하는 그런거 

	fmt.Print("이름을 입력하세요:")
	name, _ := reader.ReadString('\n') //name변수 값과 에러값 
	// name, _ := >>이거면 그냥 오류 무시

	// name, err :=  사용하는게 기본
	fmt.Printf("안녕하세요,%s", name)
}
