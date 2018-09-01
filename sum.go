package util

import "fmt"

func Sum(values interface{}) (interface{}, error) {

	switch t := values.(type) {
	case []int:
		return sumInt(t), nil
	case []float64:
		return sumFloat(t), nil
	default:
		return nil, fmt.Errorf("type of argument should be int or float64.")
	}
}

func sumInt(values []int) int {

	var sum int
	for _, val := range values {
		sum += val
	}
	return sum
}

func sumFloat(values []float64) float64 {
	var sum float64
	for _, val := range values {
		sum += val
	}
	return sum
}
