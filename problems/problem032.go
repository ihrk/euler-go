package problems

import "github.com/kavaliouIhar/euler-go/utils"

func P32() int {
	a := []int{}
	for i := 2; i < 100; i++ {
		for j := 100; ; j++ {
			t := i * j
			k := []int{}
			k = append(k, utils.NumToDigits(i)...)
			k = append(k, utils.NumToDigits(j)...)
			k = append(k, utils.NumToDigits(t)...)
			if len(k) > 9 {
				break
			}
			if len(k) == 9 && noRepeat(k) && !utils.Elem(0, k) {
				if !utils.Elem(t, a) {
					a = append(a, t)
				}
			}

		}
	}
	return utils.Sum(a)
}

func noRepeat(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				return false
			}
		}
	}
	return true
}
