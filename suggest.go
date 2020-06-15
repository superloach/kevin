package kevin

import (
	"sort"
)

type sug struct {
	w string
	d float64
}

func Suggest(s string, n int, wl Wordlist, km KeyMap) []string {
	return SuggestFn(s, n, wl, km, Distance)
}

func SuggestFn(s string, n int, wl Wordlist, km KeyMap, fn func(string, string, KeyMap) float64) []string {
	sugs := make([]sug, 0)

	for _, w := range wl {
		sugs = append(sugs, sug{
			w: w,
			d: fn(s, w, km),
		})
	}

	sort.Slice(sugs, func(i, j int) bool {
		return sugs[i].d < sugs[j].d
	})

	sws := make([]string, 0)
	for i, s := range sugs {
		if i >= n {
			break
		}
		sws = append(sws, s.w)
	}
	return sws
}
