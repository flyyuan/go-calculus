// discontinuities/detect.go

package discontinuities

import (
	"github.com/Knetic/govaluate"
	"gonum.org/v1/gonum/diff/fd"
	"math"
)

type DiscontinuityType string

const (
	Removable        DiscontinuityType = "Removable"
	Jump             DiscontinuityType = "Jump"
	Infinite         DiscontinuityType = "Infinite"
	NonRemovable     DiscontinuityType = "Non-Removable"
	NotDiscontinuous                   = "Not Discontinuous"
)

type Discontinuity struct {
	Point float64
	Type  DiscontinuityType
}

type Function func(float64) float64

// CheckDiscontinuity 检查函数在给定点的间断性
func CheckDiscontinuity(f Function, x float64) Discontinuity {
	epsilon := 1e-9
	leftLimit := f(x - epsilon)
	rightLimit := f(x + epsilon)

	// Derivative
	settings := &fd.Settings{
		Formula: fd.Central,
		Step:    1e-6,
	}
	derivative := fd.Derivative(f, x, settings)

	if math.IsNaN(derivative) {
		if leftLimit != rightLimit {
			return Discontinuity{x, Jump}
		}
		return Discontinuity{x, NonRemovable}
	}

	if math.IsInf(leftLimit, 0) || math.IsInf(rightLimit, 0) {
		return Discontinuity{x, Infinite}
	}

	if leftLimit != rightLimit {
		return Discontinuity{x, Jump}
	}

	return Discontinuity{x, NotDiscontinuous}
}

func ExpressionToFunction(expr string) (func(float64) float64, error) {
	expression, err := govaluate.NewEvaluableExpression(expr)
	if err != nil {
		return nil, err
	}

	return func(x float64) float64 {
		parameters := make(map[string]interface{}, 8)
		parameters["x"] = x
		result, err := expression.Evaluate(parameters)
		if err != nil {
			return math.NaN()
		}
		return result.(float64)
	}, nil
}

// DetectDiscontinuities 返回函数的所有间断点
func DetectDiscontinuities(expr string) ([]Discontinuity, error) {
	f, err := ExpressionToFunction(expr)
	if err != nil {
		return nil, err
	}

	// 默认的范围和步长
	start := -10.0
	end := 10.0
	step := 0.1

	var discontinuities []Discontinuity
	for x := start; x <= end; x += step {
		discontinuity := CheckDiscontinuity(f, x)
		if discontinuity.Type != NotDiscontinuous {
			discontinuities = append(discontinuities, discontinuity)
		}
	}
	return discontinuities, nil
}
