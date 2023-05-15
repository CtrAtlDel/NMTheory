package main

import (
	"math"
)

func IsZeros(array []float64) bool {
	for _, elem := range array {
		if elem != 0 && elem != 0.0 {
			return false
		}
	}
	return true
}

func GetDividers(array []float64, divider float64) []float64 {
	res := make([]float64, 0)
	if divider < 0 {
		for i := 0; i < len(array); i++ {
			if array[i] < 0 {
				res = append(res, math.Abs(array[i])/math.Abs(divider))
			} else {
				res = append(res, -1*math.Floor(-array[i]/divider))
			}
		}
		return res
	} else {
		for i := 0; i < len(array); i++ {
			res = append(res, math.Floor(array[i]/divider))
		}
		return res
	}
}

func CanResolve(aValues [][]int, m int) bool {
	isCouldResolve := true
	for j := 0; j < m; j++ {
		isCouldResolve = isCouldResolve && (aValues[j][0] == 0)
	}
	return isCouldResolve
}
