package problems

import "github.com/kavaliouIhar/euler-project-go/utils"

func P27() int {
	prs := utils.PrimeSieve(1000)
	s := 0
	max := 0
	for i := 0; i < len(prs); i++ {
		for j := -prs[i] + 2; j < 1000; j++ {
			r := 0
			for k := 0; ; k++ {
				t := k*k + k*j + prs[i]

				if utils.IsPrime(t) {
					r++
				} else {
					if r > max {
						max = r
						s = prs[i] * j
					}
					break
				}
			}
		}
	}
	return s
}
