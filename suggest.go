package kevin

import (
	"sort"
)

type sug struct {
	w string
	d float64
}

func (l Layout) Suggest(s string, num int, words []string) []string {
	sugs := make([]sug, 0)

	for _, w := range words {
		sugs = append(sugs, sug{
			w: w,
			d: l.Distance(s, w),
		})
	}

	sort.Slice(sugs, func(i, j int) bool {
		return sugs[i].d < sugs[j].d
	})

	sws := make([]string, 0)
	for i, s := range sugs {
		if i >= num {
			break
		}
		sws = append(sws, s.w)
	}
	return sws
}
