package opencage

import "strconv"

func formatFloat32Slice(v []float32) []string {
	var stringBounds []string
	for idx := range v {
		stringBounds = append(stringBounds, FormatFloat(float64(v[idx])))
	}

	return stringBounds
}

func FormatFloat(v float64) string {
	return strconv.FormatFloat(v, 'f', 7, 32)
}
