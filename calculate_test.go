package gotesting

import "testing"

func TestCalculate(t *testing.T) {

	var want int
	given1 := 5
	given2 := 0
	sign := [4]string{"+", "-", "*", "P"}

	for i := 0; i < len(sign); i++ {

		if sign[i] == "+" {
			want = given1 + given2
		} else if sign[i] == "-" {
			want = given1 - given2
		} else if sign[i] == "*" {
			want = given1 * given2
		} else if sign[i] == "/" {
			if given2 == 0 {
				t.Errorf("given2 Don't value %d when sign %s", given2, sign[i])
				given2 =-1
				want = given1/given2	
			}
		}else{
			want = -1
		}

		get := cal(given1, given2, sign[i])

		if want != get {
			t.Errorf("given [%d ,%d] but want %d When sign %s", given1, given2, want, sign[i])
		}
	}

}
