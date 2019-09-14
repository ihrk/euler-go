package problems

func P31() int {
	coinSizes := []int{1, 2, 5, 10, 20, 50, 100, 200}
	ways := make([]int, 201)
	ways[0] = 1
	for i := 0; i < len(coinSizes); i++ {
		for j := coinSizes[i]; j < len(ways); j++ {
			ways[j] += ways[j-coinSizes[i]]
		}
	}
	return ways[200]
}
