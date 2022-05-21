package gotesting

func cal(n1 ,n2 int, sign string) int {

	if sign == "+" {
		return n1 + n2
	}else if sign == "-"{
		return n1 - n2
	}else if sign == "*"{
		return n1 * n2
	}else if sign == "/"{
		return n1 / n2
	}else{
		return 0
	}
}