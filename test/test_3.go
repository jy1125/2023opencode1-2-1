package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // 현재 시간 기준으로 시드값 설정
	answer := rand.Intn(100) + 1 //1~100 랜덤정수 값 지정

	fmt.Println("Guess Game ( 1 ~ 100 ) :")
	//fmt.Println(answer)

	reader := bufio.NewReader(os.Stdin) //input이랑 비슷함 test_2참고

	for i := 0; i <= 10; i++ { //i가 10될때까지 반복
		if i == 10 { //기회인 i가 10이 되면 끝
			println("기회는 끝이다!!!!!!!")
			break
		}

		fmt.Println("You chance :", 10-i) //출력 9부터 하나씩 줄어들게
		fmt.Print("Guess number : ")
		inputNumberString, err := reader.ReadString('\n') //추리 숫자 입력 받기
		//err 쓴 건 에러 뜨면 어떻게 할 건지 _ 빈공간이면 에러 무시한다는 뜻
		if err != nil { // 에러가 아니라면
			log.Fatal(err) // log 패키지에서 에러를 출력하고 프로그램을 종료하는 코드
			//log는 라이브러리 Fatal은 프로그램중지 후 에러메시지 출력
		}

		inputNumberString = strings.TrimSpace(inputNumberString)
		// string.TrimSpace() 문자열을 가공하는 함수+앞뒤 공백 제거
		// TrimSpace() 는 파이선에서 strip()느낌

		inputNumber, err := strconv.Atoi(inputNumberString)
		if err != nil {
			log.Fatal("정수를 입력하시오")
		}
	
		//strconv는 문자열을 다양한 데이터로 바꾸는 함수 제공
		//Atoi 는 문자열을정수로 바꾸는 함수

		if inputNumber == answer {
			fmt.Println(("정답~"))
			break
		} else if inputNumber < answer {
			fmt.Println("그렇게 소심해서 뭐하니 얘야") //정답이 더 크다
		} else if inputNumber > answer {
			fmt.Println("그렇게 건방지면 되겠니?") //정답이 더 작다
		}

	}
}
