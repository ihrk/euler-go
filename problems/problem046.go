package problems

import "github.com/kavaliouIhar/euler-project-go/utils"

func P46() int {
	for i := 35; ; i += 2 {
		if utils.IsPrime(i) {
			continue
		}
		b := false
		for j := 1; j*j*2 < i; j++ {
			if utils.IsPrime(i - 2*j*j) {
				b = true
				break
			}
		}
		if b {
			continue
		}
		return i
	}
}
