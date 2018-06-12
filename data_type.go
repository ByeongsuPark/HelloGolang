package main


func main(){
	
	/**
	Go언어에는 기본적으로 다음과 같은 데이터 타입이 있다.
	1. 불린 타입
	  bool
	2. 문자열 타입
	  string: 한 번 생성되면 수정 불가능한 Immutable 타입임.
	3. 정수형 타입
	  int int8 int16 int32 int64
	  uint uint8 ...
	4. Float 및 복소수 타입
	  float32 float64 complex64 complex128

	**/

	/**
	 ``


	 **/
	
	// 문자열 리터럴은 Back Quote( ` ` ) 와 이중인용부호 ( " ")로 표현가능
	rawLiteral := `test\n`
	InterpretedLiteral := "test\n"

	

	print(InterpretedLiteral)
	print(rawLiteral)

	/**
	 test
	 test\n
	 **/

	// " "로 둘러싸인 문자열안의 이스케이프 문자들은 해당 기능을 한다.
	// ` `로 둘러싸인 문자열안의 이스케이프 문자들은 기능을 못한다.






}
