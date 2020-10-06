package problems

import "github.com/ihrk/euler-go/utils"

func P23() int {
	s := 0
	const lim = 28123
	nums := make([]int, 0, lim/4)
	abudSums := make([]bool, lim+1)
	for i := 12; i <= lim; i++ {
		if isAbudant(i) {
			nums = append(nums, i)
		}
	}
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			t := nums[i] + nums[j]
			if t > lim {
				break
			}
			abudSums[t] = true
		}
	}
	for i := 0; i <= lim; i++ {
		if !abudSums[i] {
			s += i
		}
	}
	return s
}

func isAbudant(n int) bool {
	if n%2 != 0 && n%3 != 0 {
		return false
	}
	return utils.DivSum(n) > n
}
