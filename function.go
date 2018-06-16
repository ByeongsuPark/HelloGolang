package main

func main() {

	/*
	   Go언어에서 함수의 파라미터는 파라미터 명 뒤에 int,string 등의
	   데이터 타입을 적어서 정의한다.
	*/

	msg := "hello"
	say(msg)

	/*
	   Go에서 파라미터 전달 방식은 크게 두가지이다.

	   	1. Pass By Value
		2. Pass By Reference
	  
	   C언어에서도 볼수 있는 개념들로 1번은 리터럴값 자체가 전달되는 데에 비해
	   2번은 메모리 주소값이 전달되기 때문에 근본적으로 변수 자체가 변한다
	*/

	msg_ref := "hello"
	say_ref(&msg_ref)
	println(msg_ref)

	/*
	   Variadic function ( 가변인자 함수 )
	     함수의 고정된 갯수의 파라미터를 전달하지 않고, 여러 갯수의 파라미터를 전달
	     하고자 할때 가변 파라미터를 나타내는 ... 를 사용한다.
	     아래와 같이 사용이 능하다.
	*/
	say_multiple("This", "is", "Sparta")

	/*
	   Go언어에서는 리턴값이 복수 개일 수도 있다.
	   C에서는 void 혹은 하나의 값만을 리턴하는 거소가는 대조적이다.
	   함수에서 리턴값이 있는 경우는 func문의 파라미터 괄호 다음 마지막에
	   리턴값의 타입을 정의해준다.
	*/
	s := sum(1, 3, 5, 7, 9)
	println(s)

	//복수 개의 값을 리턴하기 위해서는 해당 리턴 타입들을 괄호 ()안에 적어준다.

}

func say(msg string){

	println(msg)

}

func say_ref(msg *string){

	*msg = "hi"
	println(*msg)

}

func say_multiple(msg ...string){

	for _,s  := range msg {
		println(s)
	}
}

func sum(nums ...int) int {

	sum := 0
	for _, n := range nums{
		sum += n
	}


	     return sum
}
