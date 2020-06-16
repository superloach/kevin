package kevin

// A Layout describes the physical position of keys (runes) on a keyboard.
type Layout map[rune][2]float64

// MakeLayout allows the creation of a Layout from slices of strings.
// (see definitions of QWERTY, Dvorak, and Colemak for an example)
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

// Simple Layout definitions for standard keyboard layouts.
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
