package main

func main() {

	/*
	   Go는 반복문이 for하나 밖에 없다.
	   다른 언어와 비슷하게 
	      for 초기값 ; 조건식 ; 증감 
	   형식을 따른다. 그러나 초기값,조건식,증감을 둘러싸는 괄호를 생략한다.
	   괄호를 쓰면 에러가 난다.
	*/
	
	//1 부터 100까지의 합
	
	sum := 0
	for i :=1 ; i <= 100 ; i++ {
		sum += i
	}

	println(sum)
	
	// 무한 루프
	//for {
	//	println("Infinite loop")
	//}

	/*
	   for range 문은 컬렉션으로부터 한 요소(element)씩 가져와 차례로 for 블럭의
	   문장들을 실행한다.

	   			for 인덱스,요소값 := range 컬렉션

	   루프문을 구성한다.
	   range 키워드 다음의 컬렉션으로 하나씩 요소를 리턴해서 그 요소의 위치인덱스와
	   값을 for 키워드 다음의 2개의 변수에 각각 할당한다.
	*/

	name := []string{"홍길동", "이순신", "강감찬"}
	for index,name := range name {
		println( index, name)
	}

}
