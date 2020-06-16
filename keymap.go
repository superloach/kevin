package kevin

import "math"

type KeyMap map[rune][2]float64

func MakeKeyMap(skip rune, rowsets ...[]string) KeyMap {
	k := KeyMap{}

	k['\x00'] = [2]float64{0, 0}
	for _, rowset := range rowsets {
		for y, row := range rowset {
			for x, r := range row {
				if r == skip {
					continue
				}

				f := k['\x00'][0]
				if k['\x00'][0] == 0 {
					f = float64(r)
				}

				k[r] = [2]float64{
					float64(x),
					float64(y),
				}
				k['\x00'] = [2]float64{
					f,
					float64(r),
				}
			}
		}
	}

	return k
}

var (
	KeyMapQWERTY = MakeKeyMap('…',
		[]string{"qwertyuiop", "asdfghjkl…", "zxcvbnm………"},
		[]string{"QWERTYUIOP", "ASDFGHJKL…", "ZXCVBNM………"},
	)
	KeyMapDvorak = MakeKeyMap('…',
		[]string{"………pyfgcrl", "aoeuidhtns", "…qjkxbmwvz"},
		[]string{"………PYFGCRL", "AOEUIDHTNS", "…QJKXBMWVZ"},
	)
	KeyMapColemak = MakeKeyMap('…',
		[]string{"qwfpgjluy…", "arstdhneio", "zxcvbkm………"},
		[]string{"QWFPGJLUY…", "ARSTDHNEIO", "ZXCVBKM………"},
	)
)

func (k KeyMap) Dist(a, b rune) float64 {
	if k == nil {
		return 1
	}

	ap, aok := k[a]
	bp, bok := k[b]

	if aok && bok {
		return math.Sqrt(math.Pow(ap[0]-bp[0], 2) + math.Pow(ap[1]-bp[1], 2))
	} else {
		return 1
	}
}

func (k KeyMap) First() rune {
	return rune(k['\x00'][0])
}

func (k KeyMap) Last() rune {
	return rune(k['\x00'][1])
}

func (k KeyMap) MaxDist() float64 {
	return k.Dist(k.First(), k.Last())
}

func (k KeyMap) RatioDist(a, b rune) float64 {
	return k.Dist(a, b) / k.MaxDist()
}
