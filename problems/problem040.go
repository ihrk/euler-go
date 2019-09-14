package problems

import "github.com/kavaliouIhar/euler-go/utils"

func P40() int {
	s := 1
	for i := 1; i < 1e7; i *= 10 {
		s *= getDigit(i)
	}
	return s
}

func getDigit(n int) int {
	for i, p := 1, 9; ; i++ {
		q := i * p
		t := n - q
		if t > 0 {
			n = t
		} else {
			n--
			c := n/i + p/9
			e := utils.Pow10(i - 1 - (n % i))
			return (c / e) % 10
		}

		p *= 10
	}
}
