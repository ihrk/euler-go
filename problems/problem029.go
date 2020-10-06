package problems

import "github.com/ihrk/euler-go/utils"

func P29() int {
	const min = 2
	const max = 100
	s := 0
	nums := make([][]bool, max+1)
	for i := min; i <= max; i++ {
		if b, y, p := isPower(i); b {
			for j := min; j <= max; j++ {
				nums[y][j*p] = true
			}
		} else {
			nums[i] = make([]bool, 6*max+1)
			for j := min; j <= max; j++ {
				nums[i][j] = true
			}
		}
	}
	for i := min; i <= max; i++ {
		for j := 2; j < len(nums[i]); j++ {
			if nums[i][j] {
				s++
			}
		}
	}
	return s
}

func isPower(n int) (bool, int, int) {
	sq := utils.SqrtInt(n)
	for i := 2; i <= sq; i++ {
		p := 0
		for j := i; j <= n; j *= i {
			p++
			if j == n {
				return true, i, p
			}
		}
	}
	return false, 0, 0
}
