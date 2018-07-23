package main

import "fmt"
import "math"

/*
   구조체(struct)가 필드들의 집합체라면, interface는 메서드들의 집합체이다.
   interface는 타입(type)이 구현해야하는 메서드 원형(Prototype)들을 정의한다.
   하나의 사용자 정의 타입이 interface를 구현하기 위해서는 단순히 그 인터페이스가
   갖는 모든 메서드들을 구현하면 된다.

   인터페이스는 struct와 마찬가지로 type문을 사용하여 정의한다.
*/

type Shape interface {
	area() float64
	perimeter() float64
}

// Rect 정의
type Rect struct {
	width, height float64
}

//Circle 정의
type Circle struct {
	radius float64
}

// Rect 타입에 대한 Shape 인터페이스 구현.( java와 달리 implements 같은 키워드 쓸 필요 x)
func (r Rect) arae() float64 {
	return r.width * r.height
}
func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Circle 타입에 대한 Shape 인터페이스 구현
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

/*
   Go 프로그래밍을 하다보면 흔히 빈 인터페이스(empty interface)를 자주 접하게 되는데,
   흔히 인터페이스 타입(interface type)으로도 불리운다. 예를 들어, 여러 표준패키지들의 함수
   prototype을 살펴 보면, 아래와 같이 빈 interface 가 자주 등장함을 볼 수 있따.
   빈 interface는 interface{} 와 같이 표현한다

   func Marshal(v interface{}) ([]byte, error);

   Empty interface는 메서드를 전혀 갖지 않는 빈 인터페이스로서, Go의 모든 Type은 적어도 0개의
   메서드를 구현하므로, 흔히 Go에서 모든 Type을 나타내기 위해 빈 인터페이스를 사용한다.
   즉, 빈 인터페이스는 어떠한 타입도 담을 수 있는 컨테이너라고 볼 수 있으며,
   여러 다른 언어에서 흔히 일컫는 Dynamic Type이라 볼 수 있다.
   ( 예를 들면, java에서는 object )

*/

/*
   Type Assertion
      Interface type의 x와 타입 T에 대하여 x.(T)로 표현했을 때, 이는 x가 nil이 아니며,
      x는 T타입에 속한다는 점을 확인(assert) 하는 것으로
      이러한 표현을 "Type Assertion"이라 부른다.
      만약 x가 nil이거나 x의 타입이 T가 아니라면, 런타임에러가 발생할 것이고,
      x가 T 타입인 경우는 T타입의 x를 리턴한다.
      즉 아래 예제에서 변수 j는 a.(int)로 부터 int 형 변수 j 가 된다.
*/

func main() {

	var a interface{} = 1

	i := a       // a와 i는 dynamic type, 값은 1
	j := a.(int) // j는 int 타입, 값은 1

	fmt.Println(i) // 포인터 주소 출력
	fmt.Println(j) // 1 출력
}
