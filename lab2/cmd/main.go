package main

import (
	"fmt"
	"math"
)

func GetRangMatrix(matrix [][]int) int {
	r := 0
	end := len(matrix) - len(matrix[0])
	for i := 0; i < end; i++ {
		if matrix[i][i] > 0 {
			r = i
		} else {
			break
		}
	}
	return r + 1
}

func ResolveEquation(aMatrix [][]int, bVec [][]int) (bool, [][]int) {
	n := len(aMatrix[0])
	m := len(aMatrix)

	notBVec := make([][]int, len(bVec))
	for i, row := range bVec {
		notBVec[i] = []int{-row[0]}
	}

	zeroVec := make([][]int, n)
	for i := 0; i < n; i++ {
		zeroVec[i] = []int{0}
	}

	eMatrix := make([][]int, n)
	for i := 0; i < n; i++ {
		eMatrix[i] = make([]int, n)
		eMatrix[i][i] = 1
	}

	aEMatrixes := append(aMatrix, eMatrix...)
	cKMatrixes := make([][]int, len(aEMatrixes))
	for i := 0; i < len(aEMatrixes); i++ {
		cKMatrixes[i] = make([]int, len(aEMatrixes[i]))
		copy(cKMatrixes[i], aEMatrixes[i])
	}
	for i := 0; i < len(cKMatrixes); i++ {
		for j := i + 1; j < len(cKMatrixes[i]); j++ {
			cKMatrixes[i][j] = 0
		}
	}

	r := GetRangMatrix(cKMatrixes)
	aValues := append(notBVec, zeroVec...)

	k := 0
	for k < r && cKMatrixes[k][k] != 0 {
		kI := int(math.Floor(float64(aValues[k][0]) / float64(cKMatrixes[k][k])))
		for i := 0; i < len(aValues); i++ {
			aValues[i][0] -= kI * cKMatrixes[i][k]
		}
		k++
	}

	if CanResolve(aValues, m) {
		s := n - r
		x0 := aValues[m:]
		kVectors := make([][]int, s)
		for i := 0; i < s; i++ {
			kVectors[i] = cKMatrixes[m+i][r-1:]
		}
		kVectorsSumStr := ""
		for i := 0; i < s; i++ {
			vecStr := "t_" + fmt.Sprint(i+1) + "*" + fmt.Sprintf("%v", kVectors[i])
			kVectorsSumStr += vecStr
			if i != s-1 {
				kVectorsSumStr += " + "
			}
		}
		// fmt.Println("Particular solution is ", x0, " + ", kVectorsSumStr)
		return true, x0
	} else {
		return false, aValues[m:]
	}
}

func main() {
	fmt.Println("go")
}
