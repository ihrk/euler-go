package problems

import "github.com/kavaliouIhar/euler-go/utils"

//very bad written, since we cant proove that result is less 100
func P77() int {
	primes := utils.PrimeSieve(100)
	ways := make([]int, 101)
	ways[0] = 1
	for i := 0; i < len(primes); i++ {
		for j := primes[i]; j < len(ways); j++ {
			ways[j] += ways[j-primes[i]]
		}
	}
	for i, v := range ways {
		if v > 5000 {
			return i
		}
	}
	return 0
}
