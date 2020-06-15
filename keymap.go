package kevin

import "math"

type KeyMap map[rune][2]float64

func MakeKeyMap(rows ...string) KeyMap {
	m := KeyMap{}

	for y, row := range rows {
		for x, r := range row {
			m[r] = [2]float64{
				float64(x),
				float64(y),
			}
		}
	}

	return m
}

var (
	KeyMapQWERTY = MakeKeyMap("qwertyuiop", "asdfghjkl", "zxcvbnm")
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
