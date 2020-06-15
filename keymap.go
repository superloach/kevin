package kevin

import "math"

type KeyMap map[rune][2]float64

func MakeKeyMap(rowsets ...[]string) KeyMap {
	m := KeyMap{}

	for _, rowset := range rowsets {
		for y, row := range rowset {
			for x, r := range row {
				m[r] = [2]float64{
					float64(x),
					float64(y),
				}
			}
		}
	}

	return m
}

const z rune = '\x00'

var (
	KeyMapQWERTY = MakeKeyMap(
		[]string{"qwertyuiop", "asdfghjkl", "zxcvbnm"},
		[]string{"QWERTYUIOP", "ASDFGHJKL", "ZXCVBNM"},
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
