package discontinuities

import (
	"testing"
)

func TestDetectDiscontinuities(t *testing.T) {
	tests := []struct {
		name     string
		f        Function
		start    float64
		end      float64
		step     float64
		expected []Discontinuity
	}{
		{
			name: "Function with jump discontinuity",
			f: func(x float64) float64 {
				if x < 0 {
					return -1
				}
				return 1
			},
			start: -1,
			end:   1,
			step:  0.1,
			expected: []Discontinuity{
				{Point: 0, Type: Jump},
			},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DetectDiscontinuities(tt.f, tt.start, tt.end, tt.step)
			if len(got) != len(tt.expected) {
				t.Fatalf("expected %d discontinuities, got %d", len(tt.expected), len(got))
			}
			for i, discontinuity := range got {
				if discontinuity.Point != tt.expected[i].Point || discontinuity.Type != tt.expected[i].Type {
					t.Errorf("expected discontinuity at %f of type %s, got discontinuity at %f of type %s",
						tt.expected[i].Point, tt.expected[i].Type, discontinuity.Point, discontinuity.Type)
				}
			}
		})
	}
}
