package kevin

import "math"

// Calculates the physical distance between 2 keys (runes) on a Layout.
func (l Layout) RuneDistance(a, b rune) float64 {
	if l == nil {
		return 1
	}

	ap, aok := l[a]
	bp, bok := l[b]

	if aok && bok {
		return math.Sqrt(math.Pow(ap[0]-bp[0], 2) + math.Pow(ap[1]-bp[1], 2))
	} else {
		return 1
	}
}

// Calculates the modified Levenshtein edit distance between 2 strings, using the physical distance between keys.
func (l Layout) Distance(a, b string) float64 {
	return l.DistanceFn(a, b, func(c, d rune, x, y int) float64 {
		return l.RuneDistance(c, d)
	})
}

// Like Distance, but more harshly scoring letters toward the beginning of string a.
func (l Layout) DistanceBiased(a, b string) float64 {
	return l.DistanceFn(a, b, func(c, d rune, x, y int) float64 {
		return l.RuneDistance(c, d) / float64(x)
	})
}

// Calculates a modified Levenshtein edit distance between 2 strings, using a given function to determine the amount to add.
func (l Layout) DistanceFn(a, b string, fn func(rune, rune, int, int) float64) float64 {
	ra := []rune(a)
	rb := []rune(b)

	sy := len(ra) + 1
	sx := len(rb) + 1

	mat := make([][]float64, sy)
	for y := 0; y < sy; y++ {
		mat[y] = make([]float64, sx)
	}

	for y := 0; y < sy; y++ {
		mat[y][0] = float64(y)
		for x := 0; x < sx; x++ {
			mat[0][x] = float64(x)
		}
	}

	for y := 1; y < sy; y++ {
		for x := 1; x < sx; x++ {
			tmp1 := math.Min(mat[y][x-1]+1, mat[y-1][x]+1)

			tmp2 := mat[y-1][x-1]
			if ra[y-1] != rb[x-1] {
				tmp2 += fn(ra[y-1], rb[x-1], x, y)
			}

			mat[y][x] = math.Min(
				tmp1, tmp2,
			)
		}
	}

	return mat[sy-1][sx-1]
}
