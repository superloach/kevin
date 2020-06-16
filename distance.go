package kevin

import "math"

func Distance(s1, s2 string, km KeyMap) float64 {
	rs1 := []rune(s1)
	rs2 := []rune(s2)

	sy := len(rs1) + 1
	sx := len(rs2) + 1

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
			tmp1 := math.Min(
				mat[y][x-1]+1,
				mat[y-1][x]+1,
			)

			tmp2 := mat[y-1][x-1] + km.Dist(
				rs1[y-1],
				rs2[x-1],
			)*float64(sx-x) +
				math.Abs(float64(sx-sy)/
					float64(sx+sy))

			mat[y][x] = math.Min(tmp1, tmp2)
		}
	}

	return mat[sy-1][sx-1]
}
