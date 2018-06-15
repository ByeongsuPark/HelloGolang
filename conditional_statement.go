package main

func main(){

	k := 2

	// 1. If문
	// Go는 if 조건문의 조건식에 괄호를 둘러싸지 않는다.
	if k == 1 {
		println("One")
	}

	// 1-1. 
	//   If문과 마찬가지로 else if 혹은 else 문의
	//   블럭 시작 브레이스는 같은 라인에 있어야 한다.
	if k == 1 {
		println("One")
	} else if k ==2 { // 같은 라인
		println("Two")
	} else {	  // 같은 라인 
		println("Other")
	}

	// 2. switch문
	var name string
	var category = 1

	switch category {
		case 1:
			name = "Paper Book"
		case 2:
			name = "eBook"
		case 3, 4:
			name = "Blog"
		default:
			name = "Other"
	}
	println(name)

}
