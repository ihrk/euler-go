package euler

import (
	"math"
)

//this file contains functions, used in several problems
//all functions are pure, they dont modify argument's data

func sum(a []int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}

func odd(n int) bool {
	return n%2 == 1
}

func product(a []int) int {
	s := 1
	for _, v := range a {
		s *= v
	}
	return s
}

func fmap(f func(int) int, a []int) []int {
	c := make([]int, len(a))
	for i := range a {
		c[i] = f(a[i])
	}
	return c
}

func filter(f func(int) bool, a []int) []int {
	c := make([]int, 0, len(a))
	for _, v := range a {
		if f(v) {
			c = append(c, v)
		}
	}
	return c
}
func group(a []int) [][]int {
	d := [][]int{}
	for i := 0; i < len(a); {
		ta := []int{a[i]}
		j := i + 1
		for ; j < len(a); j++ {
			if a[j] == a[i] {
				ta = append(ta, a[j])
			} else {
				break
			}
		}
		i = j
		d = append(d, ta)
	}
	return d
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := make([][]int, 0, fac(len(arr)))

	helper = func(a []int, n int) {
		if n == 1 {
			res = append(res, copyArr(a))
		} else {
			for i := 0; i < n; i++ {
				helper(a, n-1)

				if n%2 == 1 {
					a[0], a[n-1] = a[n-1], a[0]
				} else {
					a[i], a[n-1] = a[n-1], a[i]
				}
			}
		}
	}
	helper(copyArr(arr), len(arr))
	return res
}

func copyArr(a []int) []int {
	r := make([]int, len(a))
	copy(r, a)
	return r
}

func elem(n int, a []int) bool {
	for _, e := range a {
		if n == e {
			return true
		}
	}
	return false
}

func reverse(a []int) []int {
	c := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = a[len(a)-i-1]
	}
	return c
}

////////////////////////////////////////////////////////////////////

func sqrt(n int) int {
	return int(math.Sqrt(float64(n)))
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	sq := sqrt(n)
	for i := 5; i <= sq; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primeGen(n int) []int { //return n first primes
	if n < 0 {
		return []int{}
	}
	a := make([]int, n)
	a[0] = 2
	for i, j := 1, 1; i < n; j++ {
		t := 2*j + 1
		if isPrime(t) {
			a[i] = t
			i++
		}
	}
	return a
}

func primeSieve(n int) []int {
	nums := make([]bool, n+1)
	sq := sqrt(n)
	for i := 2; i <= sq; i++ {
		if !nums[i] {
			for j := i * i; j <= n; j += i {
				nums[j] = true
			}
		}
	}
	r := make([]int, 0)
	for i := 2; i <= n; i++ {
		if !nums[i] {
			r = append(r, i)
		}
	}
	return r
}

func primeSieveRange(min, max int) []int {
	nums := make([]bool, max+1)
	sq := sqrt(max)
	for i := 2; i <= sq; i++ {
		if !nums[i] {
			for j := i * i; j <= max; j += i {
				nums[j] = true
			}
		}
	}
	r := make([]int, 0)
	for i := min; i <= max; i++ {
		if !nums[i] {
			r = append(r, i)
		}
	}
	return r
}

func primeFactors(n int) []int {
	if n < 2 {
		return []int{}
	}
	sq := sqrt(n)
	r := []int{}
	for i := 2; i <= sq; {
		if n%i == 0 {
			r = append(r, i)
			n /= i
		} else {
			i++
		}
	}
	if n != 1 {
		r = append(r, n)
	}
	return r
}

func primeFactorsUniq(n int) []int {
	if n < 2 {
		return []int{}
	}
	sq := sqrt(n)
	r := []int{}
	for i := 2; i <= sq; {
		if n%i == 0 {
			r = append(r, i)
			for n%i == 0 {
				n /= i
			}
		} else {
			i++
		}
	}
	if n != 1 {
		r = append(r, n)
	}
	return r
}

//effective way to calculate sum of divisors of number
func divSum(n int) int {
	m := n
	s := 1
	sq := sqrt(n)
	for i := 2; i <= sq; i++ {
		p := 1
		for n%i == 0 {
			p = p*i + 1
			n /= i
		}
		s *= p
	}
	if n > 1 {
		s *= 1 + n
	}
	return s - m
}

func numToDigits(n int) []int {
	if n <= 0 {
		return []int{}
	}
	a := []int{}
	for n > 0 {
		a = append(a, n%10)
		n /= 10
	}
	return reverse(a)
}

func digitsToNum(a []int) int {
	s := 0
	for _, v := range a {
		s *= 10
		s += v
	}
	return s
}

///////////////////////////////////////////////////////////////////////

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}

func maximum(a []int) int {
	if len(a) == 0 {
		panic("no maximum value for empty list")
	}
	s := a[0]
	for i := 1; i < len(a); i++ {
		s = max(s, a[i])
	}
	return s
}

func min(n, m int) int {
	if n < m {
		return n
	}
	return m
}

func minimum(a []int) int {
	if len(a) == 0 {
		panic("no minimum value for empty list")
	}
	s := a[0]
	for i := 1; i < len(a); i++ {
		s = min(s, a[i])
	}
	return s
}

//////////////////////////////////////////////////////////////////////////

func log(n, p int) int {
	s := 0
	for n >= p {
		n /= p
		s++
	}
	return s
}

func pow(n, p int) int {
	if n == 1 {
		return 1
	}
	if n == -1 {
		if p%2 == 0 {
			return 1
		}
		return -1
	}
	if p < 0 {
		return 0
	}
	s := 1
	for i := 0; i < p; i++ {
		s *= n
	}
	return s
}

func pow10(p int) int {
	if p < 0 {
		return 0
	}
	s := 1
	for i := 0; i < p; i++ {
		s *= 10
	}
	return s
}

func fac(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 1
	}
	return n * fac(n-1)
}

func gcd(n, m int) int {
	if n == 0 || m == 0 {
		panic("zero value")
	}

	for n != 0 {
		n, m = m%n, n
	}
	return abs(m)
}
