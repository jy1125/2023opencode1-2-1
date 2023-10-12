package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	println("Hello Go!")
	fmt.Println("오픈소스 프로그래밍", "go")
	//변수 선언 방식 1
	var a int
	a = 1
	var b float32 = 1.5
	var c string = "안녕"
	//변수 선언 방식 2
	d := 5 // 얘는 타입명 안 써도 자동으로 함

	//변수 선언 _ 제로변수
	var e int    // 변수 선언 후 값을 넣지 않으면 0
	var f string //string 타입은 0도 안나오고 공백처리
	println(a, b, c)
	println(d)
	println(e)
	println(f)

	//타입 변환
	// 타입이 다르면 연산, 비교가 X
	hi := 9
	bye := 9.5
	fmt.Println(hi > int(bye)) // 일시적으로 이 라인만 bye가 int형으로 바뀜
	//fmt.Println(hi > bye) 오류

	//자료형 확인 함수
	fmt.Println(reflect.TypeOf(hi))
	fmt.Println(reflect.TypeOf(bye))

	//수학 함수
	fmt.Println(math.Floor(2.75)) //버림 함수
	fmt.Println(math.Ceil(2.75)) // 올림 함수
    
}
