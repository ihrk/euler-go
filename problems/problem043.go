package problems

import "github.com/ihrk/euler-go/utils"

func P43() int {
	d := utils.Permutations([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	s := 0
	for _, v := range d {
		if utils.DigitsToNum(v[1:4])%2 != 0 {
			continue
		}
		if utils.DigitsToNum(v[2:5])%3 != 0 {
			continue
		}
		if utils.DigitsToNum(v[3:6])%5 != 0 {
			continue
		}
		if utils.DigitsToNum(v[4:7])%7 != 0 {
			continue
		}
		if utils.DigitsToNum(v[5:8])%11 != 0 {
			continue
		}
		if utils.DigitsToNum(v[6:9])%13 != 0 {
			continue
		}
		if utils.DigitsToNum(v[7:10])%17 != 0 {
			continue
		}
		s += utils.DigitsToNum(v)
	}
	return s
}
