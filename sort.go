package stalinSort

func Sort(s []int) []int {
	l := len(s)
	if l == 0 {
		return []int{}
	}
	r := make([]int, l)
	r[0] = s[0]
	si := 1
	ri := 0
	for si < l {
		if s[si] >= r[ri] {
			ri = ri + 1
			r[ri] = s[si]
		}
		si = si + 1
	}
	return r[0 : ri+1]
}
