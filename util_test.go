package opencage

import (
	"strings"
	"testing"
)

func TestFormatFloat32Slice(t *testing.T) {
	result := formatFloat64Slice([]float32{-0.563160, 51.280430, 0.278970, 51.683979})
	actual := strings.Join(result, ",")
	if actual != "-0.5631600,51.2804298,0.2789700,51.6839790" {
		t.Fatalf("incorrect float slice format")
	}
}
