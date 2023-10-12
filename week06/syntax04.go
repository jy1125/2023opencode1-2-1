package main
import (
	"fmt"
	"math/rand"
	"time"
)
// 난수 추출된 수의 소수 판정 프로그램 v0.6
func main() {
	seed := time.Now().Unix() //현재 시간을 기준으로 시드값 설정
	rand.Seed(seed) // 이게 없으면 매번 같은 랜덤 수만 뜸

	isPrime := true
	number := rand.Intn(150) + 2
	//number = 21
	fmt.Println("임의로 추출된 수 : ", number)
	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
			break // 첫 번째 약수가 발견되면 반복문 즉시 종료
		}
		//fmt.Print(i, " ")
	}
	if isPrime { // 비교 연산자 제거
		fmt.Println(number, "는(은) 소수입니다!")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다~")
	}
}
// 난수 추출된 수의 소수 판정 프로그램 v0.4
//ackage main
//
//mport (
//	"fmt"
//	"math/rand"
//	"time"
//
//
//unc main() {
//	seed := time.Now().Unix()
//	rand.Seed(seed)
//	isPrime := true
//	number := rand.Intn(150) + 2
//	fmt.Println("임의로 추출된 수 : ", number)
//
//	for i := 2; i < number; i++ {
//		if number%i == 0 {
//			isPrime = false
//			fmt.Println("number의 약수: ", i, " ")
//		}
//
//		//fmt.Print(i, " ") for문 밖에다 쓰면 2부터 number수 까지 다 나옴
//	}
//	if isPrime { // 비교 연산자 제거 //어짜피 true나 false이니까
//		fmt.Println(number, "는(은) 소수입니다!")
//	} else {
//		fmt.Println(number, "는(은) 소수가 아닙니다~")
//	}
//
//
// 난수 추출된 수의 소수 판정 프로그램 v0.3
//package main
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//func main() {
//	seed := time.Now().Unix()
//	rand.Seed(seed)
//	isPrime := true
//	number := rand.Intn(150) + 2
//	fmt.Println("임의로 추출된 수 : ", number)
//	for i := 2; i < number; i++ {
//		if number%i == 0 { //2부터 본인 전까지니까 소수면 무조건 나머지가 생김
//			isPrime = false //나머지가 2로 나눴을 때 0이면 false 다음 숫자인 3으로 나눴을 때 0이면 false
//			// 최소 소수가 아니면 0으로 나눠짐
//			//count = count + 1
//		}
//	}
//	if isPrime == true {
//		fmt.Println(number, "는(은) 소수입니다!")
//	} else {
//		fmt.Println(number, "는(은) 소수가 아닙니다~")
//	}
//}

// 난수 추출된 수의 소수 판정 프로그램 v0.2
//package main
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//ackage main
//
//mport (
//	"fmt"
//	"math/rand"
//	"time"
//
//
//unc main() {
//	seed := time.Now().Unix()
//	rand.Seed(seed)
//	count := 0
//	number := rand.Intn(150) + 2 // 0과 1제외, 2 ~ 151 사이의 수
//	fmt.Println("임의로 추출된 수 : ", number)
//	for i := 2; i < number; i++ { // 1과 number일때 loop 돌지 않음
//		if number%i == 0 {
//			//count++
//			count = count + 1
//		}
//	}
//	if count == 0 {
//		fmt.Println(number, "는(은) 소수입니다!")
//	} else {
//		fmt.Println(number, "는(은) 소수가 아닙니다~")
//	}
//

//package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// // 난수 추출된 수의 소수 판정 프로그램 v0.1
// // 소수 : 1과 자기자신외에는 나누어 떨어지지 않는 수 (0과 1은 제외)
// func main() {
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	count := 0
// 	number := rand.Intn(150) + 2 // 0과 1제외, 2 ~ 151 사이의 수
// 	fmt.Println("임의로 추출된 수 : ", number)

// 	for i := 1; i <= number; i++ {
// 		if number%i == 0 {
// 			count++
// 		}
// 	}

// 	if count == 2 { //1과 자기 자신으로 인한 count 증가 => 2번
// 		fmt.Println(number, "는(은) 소수입니다!")
// 	} else { //2번 이상이라는건 1과 본인 말고도 다른 숫자로 쪼개지기에 소수가 아님
// 		fmt.Println(number, "는(은) 소수가 아닙니다~")
// 	}
// }

//package main

//import (
//	"fmt"
//	"math/rand"
//	"time" // seed 생성용 패키지
//)

//func main() {
//	// seed 설정
//	seed := time.Now().Unix()
//	rand.Seed(seed)
//
//	for i := 1; i < 6; i++ { //5번 반복
//		dice := rand.Intn(6) + 1 //정수형 1-6까지 랜덤으로 나옴
//		fmt.Println(dice)
//	}
//}
