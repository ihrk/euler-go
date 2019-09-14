package problems

import (
	"sort"

	"github.com/kavaliouIhar/euler-go/utils"
)

func P119() int {
	a := []int{}
	for i := 2; len(a) < 30; i++ { // i - number of digits in checked number
		min := utils.Pow10(i - 1)
		max := utils.Pow10(i)
		for j := 2; j <= 9*i; j++ {
			for k := 2; ; k++ {
				t := utils.Pow(j, k)
				if t < min {
					continue
				}
				if t >= max {
					break
				}
				if r := utils.Sum(utils.NumToDigits(t)); r == j {
					a = append(a, t)
				}
			}
		}
	}
	// sort, since its unknown how many numbers where added
	sort.Ints(a) // in the last iteration

	return a[29]
}
