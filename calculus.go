package calculus

import "github.com/flyyuan/go-calculus/discontinuities"

func Discontinuities(f discontinuities.Function, start, end, step float64) []discontinuities.Discontinuity {
	return discontinuities.DetectDiscontinuities(f, start, end, step)
}
