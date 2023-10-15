package main

import (
	"bufio"
	"fmt"
	"log" //Fatal function
	"os"
	"strconv" //TrimSpace
	"strings" //parseint
)

func main() {
	fmt.Print("단 입력:")
	rd := bufio.NewReader(os.Stdin) //rd 라는 변수명으로 입력받음
	in, err := rd.ReadString('\n')  //입력받은 rd를 str으로 읽어 in에 저장

	if err != nil { // 처음 입력 받을 때 str말고 물음표나 다른 걸 받았을 경우
		log.Fatal(err) //에러 메시지 나옴
	}

	in = strings.TrimSpace(in) //str형 변수in의 공백을 없앰
	//dan, err := strconv.Atoi(in) 이 형식을 쓰면 밑에 출력에서 int 쓸 이유가 없음
	dan, err := strconv.ParseInt(in, 10, 32) //str으로 받은 숫자 in을 10진수 int32형식으로 바꾸어 dan에 저장
	if err != nil {                          //에러가 발생하면
		log.Fatal(err) //여기 있는 err은 int형으로 바꾸다가 에러 났을 경우
	}

	for i := 1; i < 10; i++ {
		fmt.Println(dan, " * ", i, " = ", (int(dan) * i))
	}
	//fmt.Println(dan * 2) // dan *2 int int형이니 오류X
}
