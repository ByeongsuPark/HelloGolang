package main

import (
	"fmt"
	"log"
	"os"
)

/*
	1.Go 에러
	   Go는 내장 타입으로 error라는 Interface 탕ㅂ을 갖는다.
	   Go에러는 이 error 인터페이스를 통해서 주고 받는다.
	   이 interface는 다음과 같은 하나의 메서들ㄹ 가즌ㄴ다.
	   개발자는 이 인터페이스를 구현하는 커스텀 에러 타입을 만들 수 있다.
*/

type error interface {
	Error() string
}

/*
	2.Go 에러처리
	   Go함수가 결과와 에러를 함께 리턴한다면, 이 에러가 nil인지를 체크해서
	   에러가 없는지를 체크할 수 있다. 예를 들어, os.Open() 함수는
	   func Open(name string) (file *File, err error)과 같은 함수 원형을
	   갖는 것으로 첫번째는 File포인터를 두번째는 error 인터페이스를 리턴한다.
	   그래서 이경우는 두번째 error를 체크해서 nil이면 에러가 없는 것이고,
	   nil이 아니면 err.Error()로부터 해당 에러를 알 수 있다.
	   아래 예쩨는 파일을 오픈하는데 에러가 발생하면 에러메시지를 출력하고 빠져나가는
	   예이다.
*/

func main() {
	f, err := os.Open("~/abc.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(f.Name())

}
