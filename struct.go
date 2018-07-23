package main

import "fmt"

// Go 에서 struct는 Custom Data Type을 표현하는데 사용된다. 참고롤 Go는 메서드를 갖지 않는다.

/*
   Go언어는 OOP를 고유의 방식으로 지원한다. 즉, 전통적인 OOP 언어가 가지는
   클래스, 객체, 상속 개념이 없다.
   전통적인 OOP의 클래스는 Go언어에서 struct로 표현되는데,
   전통적인 OOP의 클래스가 메서드와 데이터를 함께 갖는 것과는 달리
   Go는 데이터 필드만을 가지며, 메서드는 별도로 분리하여 정의한다.
*/

// struct를 정의하기 위해서는 type 문을 사용한다.

type person struct {
	name string
	age  int
}

/*
   생성자 함수
     때로 구조체 필드가 사용 전에 초기화되어야 하는 경우가 이다. 예를 들어, struct 필드가
     map 타입인 경우 map을 사전에 미리 초기화해 놓으면, 외부 struct 사용자가 매번 map을
     초기화 해야 된다는 것을 기억할 필요가 없다.
     이러한 목적을 위해 생성자 함수를 사용할 수 있다.
     생성자 함수는 struct를 리턴하는 함수로써 그 함수 본문에서 필요한 필드를 초기하 한다.
*/

type dict struct {
	data map[int]string
}

// 생성자 함수 정의
func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d // 포인터 전달
}

func main() {

	// struct 객체 생성 방식
	// 1. 빈 객체 생성 후 필드 값을 채워 넣는 방식
	// person 객체 생성
	p := person{}

	// 필드값 설정
	p.name = "Byeongsu Park" // struct 객체의 필드에 접근하기 위해서는 . (dot)을 사용
	p.age = 21

	fmt.Println(p)

	// 2. 객체 생성과 동시에 초기화
	var p2, p3 person

	p2 = person{"Byeongsu Park", 21}

	// 필드명을 지정해서 객체 생성하는 경우는, 순서는 상관없다.
	// 다만 생략된 필드들은 Zero value를 갖는다.
	p3 = person{age: 21, name: "Byeongsu Park"}

	fmt.Println(p2)
	fmt.Println(p3)

	dic := newDict()
	dic.data[1] = "A"

	fmt.Println(dic.data[1])
}
