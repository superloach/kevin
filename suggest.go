package kevin

import "sort"

type sug struct {
	word  string
	score float64
}

func Suggest(s string, n int, wl Wordlist, km KeyMap) []string {
	sugs := make([]sug, 0)

	for _, word := range wl {
		if len(word) >= len(inp) {
			sugs = append(sugs, sug{
				word:  word,
				score: Distance(inp, word, km),
			})
		}
	}

	sort.Slice(sugs, func(i, j int) bool {
		is, js := sugs[i].score, sugs[j].score
		return is < js
	})

	ssugs := make([]string, 0)
	for i, s := range sugs {
		if i >= n {
			break
		}
		ssugs = append(ssugs, s.word)
	}
	return ssugs
}
