package problems

import "github.com/ihrk/euler-go/utils"

func P123() int {
	for i, n := 3, 1; ; i += 2 {
		if utils.IsPrime(i) {
			n++
			if n%2 == 0 {
				continue
			}
			if 2*i*n > 1e10 {
				return n
			}
		}
	}
}
