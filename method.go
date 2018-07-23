package main

import "fmt"

/*
   1.Go Method
     Go언어는 OOP를 고유의 방식으로 지원한다.
     타 언어의 OOP의 클래스가 필드와 메서드를 함께 갖는 것과 달리 Go언어에서는 struct
     가 필드만을 가지며, 메서드는 별도로 분리되어 정의된다.

     Go메서드는 특별한 형태의 func 함수이다. 메서드는 함수 정의에서,
     func 키워드와 함수명 사이에 "그 함수가 어떤 struct를 위한 메서드인지"를 표시하게 된다.

     흔히 receiver로 불리우는 이 부분은 메서드가 속한 struct타입과 struct 변수명을
     지정하는데, struct 변수명은 함수 내에서 마치 입력 파라미터 처럼 사용된다.
*/
type Rect struct {
	width, height int
}

func (r Rect) area() int {
	return r.width * r.height // struct 변수명은 입력 파라미터 처럼 사용...
}

func main() {

	var rect Rect

	rect = Rect{7, 8}
	area := rect.area()

	fmt.Println(area)
}
