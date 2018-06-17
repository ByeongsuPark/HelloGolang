package main

func main(){

	/*
	   1.익명함수
	   익명함수는 함수명이 없는 함수이다.
	   일반적으로 익명함수는 그 함수 전체를 변수에 할당하거나, 다른 함수의 파라미터에
	   직접 정의되어 사용되곤 한다.
	   익명함수가 변수에 할당된 이후에는 변수명이 함수명과 같이 취급된다.
	*/

	sum := func( n ...int) int{
		
		s := 0
		for _, i := range n {
			s += i
		}

		   return s
	     }

	result := sum(1, 2, 3, 4, 5)
	println(result)
	
	/*
	   2.일급함수
	    Go언어에서 함수는 일급함수로서 Go의 기본 타입과 동일하게 취급된다.
	    따라서 다른 함수의 파라미터로 전달하거나, 다른 함수의 리턴값으로 사용가능하다.
	    함수를 다른 함수의 파라미터로 전달하기 위해서는 익명함수를 변수에 할당한 후
	    이 변수를 전달하는 방법과 직접 다른 함수 호출 파라미터에 함수를 적는 방법이 있다.
	*/

	// 변수 add에 함수 전달
	add := func(i int, j int) int{
		
		    total := 0
		    total = i + j
		    return total
	}
	
	// add 함수 전달
	r1 := calc( add, 10, 20)
	println(r1)
}

	/*
	   3.type문을 사용한 함수 원형 정의
	     type문은 구조체(struct), 인터페이스 등 Custom Type을 정의하기 위해 사용된다.
	     type문은 또한 함수 원형을 정의하는데 사용이 가능하다.
	     아래 예제에서 func(int, int) int 가 반복되는 경우 type문을 사용하여
	     간단하게 해당 함수의 원형을 표현할 수 있다.
	     이렇게 함수의 원형을 정의하고 함수를 타 메서드에 전달하고 리턴받는 기능을
	     타 언어에서는 흔히 델리게이트(Delegate)라 부른다.
	*/
type calculator func(int, int) int

func calc(f calculator, a int, b int) int{

	result := f(a, b)
	return result

}
