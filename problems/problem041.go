package problems

import "github.com/kavaliouIhar/euler-go/utils"

func P41() int {
	a := []int{}
	return newL(0, a)
}

func newL(m int, a []int) int {
	for i := 1; i <= 7; i++ {
		t := utils.CopyArr(a)
		if !utils.Elem(i, t) {
			t = append(t, i)
			if len(t) == 7 || len(t) == 4 {
				r := utils.DigitsToNum(t)
				if len(t) == utils.Maximum(t) && r > m && utils.IsPrime(r) {
					m = r
				}
			}
			if len(t) < 7 {
				m = newL(m, t)
			}
		}
	}
	return m
}
