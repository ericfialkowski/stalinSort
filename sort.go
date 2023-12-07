package stalinSort

import "cmp"

func Sort[V cmp.Ordered](s []V) []V {
	l := len(s)
	if l == 0 {
		return []V{}
	}
	r := make([]V, l)
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

func SortDesc[V cmp.Ordered](s []V) []V {
	l := len(s)
	if l == 0 {
		return []V{}
	}
	r := make([]V, l)
	r[0] = s[0]
	si := 1
	ri := 0
	for si < l {
		if s[si] <= r[ri] {
			ri = ri + 1
			r[ri] = s[si]
		}
		si = si + 1
	}
	return r[0 : ri+1]
}
