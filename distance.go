package kevin

import "math"

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

func (l Layout) Distance(a, b string) float64 {
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

			tmp2 := mat[y-1][x-1] + l.RuneDistance(ra[y-1], rb[x-1])*float64(sx-x)/float64(sy-y)

			mat[y][x] = math.Min(
				tmp1, tmp2,
			)
		}
	}

	return mat[sy-1][sx-1]
}
