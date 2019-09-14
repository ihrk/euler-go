package problems

import "github.com/kavaliouIhar/euler-go/utils"

func P37() int {
	s := 0
	for i, c := 1, 0; c < 11; i++ {
		if isTrunc(i) {
			s += i
			c++
		}
	}
	return s
}

func isTrunc(n int) bool {
	if n < 10 {
		return false
	}
	if !utils.IsPrime(n) {
		return false
	}
	lt := utils.Pow10(utils.Log(n, 10))
	for ; lt > 1; lt /= 10 {
		if !utils.IsPrime(n % lt) {
			return false
		}
	}
	for n /= 10; n > 0; n /= 10 {
		if !utils.IsPrime(n) {
			return false
		}
	}
	return true
}
