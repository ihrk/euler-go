package problems

func P75() int {
	ps := []int{}
	const lim = 1500000
	genPer(3, 4, 5, lim, &ps)
	nums := make([]int, lim+1)
	for i := 0; i < len(ps); i++ {
		for j := ps[i]; j < len(nums); j += ps[i] {
			nums[j]++
		}
	}
	s := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			s++
		}
	}
	return s
}

func genPer(a, b, c, n int, ps *[]int) {
	if p := a + b + c; p <= n {
		*ps = append(*ps, p)
	} else {
		return
	}
	genPer(a-2*b+2*c, 2*a-b+2*c, 2*a-2*b+3*c, n, ps)
	genPer(a+2*b+2*c, 2*a+b+2*c, 2*a+2*b+3*c, n, ps)
	genPer(-a+2*b+2*c, -2*a+b+2*c, -2*a+2*b+3*c, n, ps)
}
