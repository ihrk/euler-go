package problems

import "github.com/kavaliouIhar/euler-go/utils"

func P35() int {
	s := 0
	prs := utils.PrimeSieve(1e6)
	for _, v := range prs {
		if !noEven(v) && v > 10 {
			continue
		}
		t := rotateNum(v)
		if utils.IsPrime(t) {
			b := true
			for t != v {
				t = rotateNum(t)
				if !utils.IsPrime(t) {
					b = false
					break
				}
			}
			if b {
				s++
			}
		}
	}
	return s
}

func rotateNum(n int) int {
	if n < 10 {
		return n
	}
	a := utils.NumToDigits(n)
	t := a[0]
	for i := 1; i < len(a); i++ {
		a[i-1] = a[i]
	}
	a[len(a)-1] = t
	return utils.DigitsToNum(a)
}

func noEven(n int) bool {
	if n < 1 {
		return false
	}
	for n > 1 {
		if n%2 == 0 || n%5 == 0 {
			return false
		}
		n /= 10
	}
	return true
}
