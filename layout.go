package kevin

type Layout map[rune][2]float64

func MakeLayout(skip rune, rowsets ...[]string) Layout {
	l := Layout{}

	for _, rowset := range rowsets {
		for y, row := range rowset {
			for x, r := range row {
				if r == skip {
					continue
				}

				l[r] = [2]float64{
					float64(x),
					float64(y),
				}
			}
		}
	}

	return l
}

var (
	QWERTY = MakeLayout('…',
		[]string{"qwertyuiop", "asdfghjkl…", "zxcvbnm………"},
		[]string{"QWERTYUIOP", "ASDFGHJKL…", "ZXCVBNM………"},
	)
	Dvorak = MakeLayout('…',
		[]string{"………pyfgcrl", "aoeuidhtns", "…qjkxbmwvz"},
		[]string{"………PYFGCRL", "AOEUIDHTNS", "…QJKXBMWVZ"},
	)
	Colemak = MakeLayout('…',
		[]string{"qwfpgjluy…", "arstdhneio", "zxcvbkm………"},
		[]string{"QWFPGJLUY…", "ARSTDHNEIO", "ZXCVBKM………"},
	)
)
