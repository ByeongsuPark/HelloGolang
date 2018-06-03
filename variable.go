package main

func main() {
	// 변수는 Go 키워드 var 사용하여 선언
	// var [변수명] [변수타입]
	var a int

	// int형 변수 선언과 동시에 초기화
	var b int = 11
	
	// 동일한 타입의 변수가 여러개일 경우 아래와 같이 선언 가능.
	var c, d, e int

	// 복수개의 변수를 선언과 동시에 초기화
	var f, g, h int = 1, 2, 3

	/*
	 * 변수를 선언하면서 초기값을 지정하지 않으면
	 * Go는 Zero Value를 기본적으로 할당한다.
	 * 즉, 숫자형은 0
	       bool타입은 false
	       String 형은 ""(빈문자열)을 할당한다.
	 */
	
	// Go는 할당되는 값을 보고 타입을 추론하는 기능이 있다.
	// 아래의 코드를 보면 i는 자동으로 int형으로
	// s는 문자열로 타입이 할당된다.
	var i = 1
	var s = "Hello"

	/* 또 다른 변수 선언 방식으로 Short Assignment Statement가 있다.
	 * 아래와 같이 := 를 사용하여 선언한다.
	 * 그러나, 이러한 표현은 함수 내에서만 사용이 가능하다.
	 * 함수 밖에서는 var를 사용해야 한다.
	 */
	i := 1

	/* 상수는 키워드 const를 사용하여 선언이 가능하며,
	 * 여러 개의 상수를 묶어서 지정할 수 있다
	 * 괄호 안에 상수들을 나열하여 선언한다.
	 */
	const (
			Visa = "Visa"
			Master = "MasterCard"
			Amex = "American Express"
	      )

	/* 상수값을 0으로부터 순차적으로 부여하기 위해 ( c의 enum 같은..)
	 * iota 라는 iota라는 식별자를 사용할 수 있다.
	 * 이 경우 iota가 지정된 변수는 0이 할당되구 이후 1씩 증가된
	 * 값을 부여 받는다.
	 */

	const(
			Apple = iota // 0
			Graepe // 1
			Orange // 2
	     )
}

