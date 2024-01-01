package opencage

import "strconv"

func formatFloat64Slice(v []float64) []string {
	var stringBounds []string
	for idx := range v {
		stringBounds = append(stringBounds, FormatFloat(v[idx]))
	}

	return stringBounds
}

func FormatFloat(v float64) string {
	return strconv.FormatFloat(v, 'f', 14, 32)
}
