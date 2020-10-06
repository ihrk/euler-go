package problems

import "github.com/ihrk/euler-go/utils"

func P76() int {
	return partNum(100) - 1
}

func partNum(n int) int {
	ps := make([][]int, n+1)
	ps[0] = []int{1}
	ps[1] = []int{0, 1}
	for i := 2; i <= n; i++ {
		ps[i] = make([]int, i+1)
		ps[i][1] = 1
		for j := 2; j <= i; j++ {
			ps[i][j] = ps[i][j-1] + ps[i-j][utils.Min(j, i-j)]
		}
	}

	return ps[n][n]

}
