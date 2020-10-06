package problems

import "github.com/ihrk/euler-go/utils"

func P33() int {
	nmr := 1
	dmr := 1
	for i := 10; i < 99; i++ {
		for j := i + 1; j < 100; j++ {
			t1, t11 := i/10, i%10
			t2, t22 := j%10, j/10
			if i*t2 == j*t1 && t11 == t22 {
				nmr *= i
				dmr *= j
			}
		}
	}
	return dmr / utils.Gcd(nmr, dmr)
}
