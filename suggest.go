package kevin

import (
	"sort"
)

type sug struct {
	w string
	d float64
}

// Suggest a number of spelling corrections using the Distance scoring function.
func (l Layout) Suggest(s string, num int, words []string) []string {
	return l.SuggestFn(s, num, words, l.Distance)
}

// Suggest a number of spelling corrections using the DistanceBiased scoring function.
func (l Layout) SuggestBiased(s string, num int, words []string) []string {
	return l.SuggestFn(s, num, words, l.DistanceBiased)
}

// Suggest a number of spelling corrections using a given scoring function.
func (l Layout) SuggestFn(s string, num int, words []string, fn func(string, string) float64) []string {
	sugs := make([]sug, 0)

	for _, w := range words {
		sugs = append(sugs, sug{
			w: w,
			d: fn(s, w),
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
