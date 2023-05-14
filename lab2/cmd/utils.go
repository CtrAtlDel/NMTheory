package main

import (
	"math"
	"sort"

	"gonum.org/v1/gonum/mat"
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

func GetDivider(number float64, divider float64) float64 {
	if divider < 0 {
		if number < 0 {
			return math.Abs(number) / math.Abs(divider)
		} else {
			return -1 * (math.Floor(-number / divider))
		}
	} else {
		return math.Floor(number / divider)
	}
}

func Devide(matrix [][]float64) [][]float64 {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for !IsZeros(matrix[i][i+1:]) {
			sort.Slice(matrix[i:], func(k, l int) bool {
				return matrix[i+k][i] < matrix[i+l][i]
			})
			dividers := GetDividers(matrix[i][i+1:], matrix[i][i])
			for j, div := range dividers {
				for k := 0; k < n; k++ {
					matrix[k][i+1+j] -= matrix[k][i] * div
				}
			}
		}
	}
	return matrix
}

func Sort(matrix [][]float64) [][]float64 {
	k := len(matrix[0]) - 1
	for k > 0 {
		ind := 0
		for j := 1; j <= k; j++ {
			if math.Abs(matrix[0][j]) > math.Abs(matrix[0][ind]) && matrix[0][ind] > 0 {
				ind = j
			}
		}
		N := len(matrix)
		for i := 0; i < N; i++ {
			b := matrix[i][ind]
			matrix[i][ind] = matrix[i][k]
			matrix[i][k] = b
		}
		k--
	}
	nulls := 0
	N := len(matrix[0])
	for i := 0; i < N; i++ {
		if matrix[0][i] != 0 {
			break
		}
		nulls++
	}
	if nulls > 0 {
		for i := 0; i < len(matrix); i++ {
			b := matrix[i][nulls]
			matrix[i][nulls] = matrix[i][0]
			matrix[i][0] = b
		}
	}
	return matrix
}

func CanResolve(aValues [][]int, m int) bool {
	isCouldResolve := true
	for j := 0; j < m; j++ {
		isCouldResolve = isCouldResolve && (aValues[j][0] == 0)
	}
	return isCouldResolve
}

func GetRank(matrix *mat.Dense) int {
	_, numCols := matrix.Dims()
	rank := 0
	for i := 0; i < numCols; i++ {
		if matrix.At(i, i) > 0 {
			rank++
		} else {
			break
		}
	}
	return rank
}
