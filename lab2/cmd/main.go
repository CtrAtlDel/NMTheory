package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const err1 = "Введите путь до файла как аргумент"
const err2 = "Введена некорректная матрица"
const split = "\n"

type DiophantineSystem struct {
	matrix  [][]int
	r, n, m int
	x       [][]int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(err1)
		return
	}
	file := os.Args[1]
	d, err := ReadAndSolve(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	d.divideLast()
	d.getSolution()
	d.printSolutions()
}

func ReadAndSolve(file string) (*DiophantineSystem, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	rows := make([]string, 0)
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			return nil, err
		}
		rows = append(rows, string(buf[:n]))
	}

	d := &DiophantineSystem{}
	lines := strings.Split(rows[0], split)
	dimensions := strings.Split(strings.TrimSpace(lines[0]), " ")
	d.r = 0
	d.n, _ = strconv.Atoi(dimensions[0])
	d.m, _ = strconv.Atoi(dimensions[1])
	d.x = make([][]int, 0)
	d.matrix = make([][]int, d.n)
	for i := 1; i <= d.n; i++ {
		line := strings.TrimSpace(lines[i])
		if len(strings.Split(line, " ")) != d.m {
			return nil, fmt.Errorf(err2)
		}
		d.matrix[i-1] = make([]int, d.m)
		row := strings.Split(line, " ")
		for j := 0; j < d.m; j++ {
			value, _ := strconv.Atoi(row[j])
			d.matrix[i-1][j] = value
		}
	}
	d.Trill()
	d.Devide()
	return d, nil
}

func (d *DiophantineSystem) Trill() {
	for i := 0; i < d.n; i++ {
		d.matrix[i][d.m-1] = -d.matrix[i][d.m-1]
	}
	for i := 0; i < d.m-1; i++ {
		line := make([]int, d.m)
		line[i] = 1
		d.matrix = append(d.matrix, line)
	}
	d.n += d.m - 1
}

func (d *DiophantineSystem) Devide() {
	for i := 0; i < d.n; i++ {
		for !d.isZero(i) {
			index := d.findMin(d.matrix[i], i)
			if d.matrix[i][index] == 0 {
				continue
			}
			if d.matrix[i][index] < 0 {
				d.inverse(index)
			}
			if index != i {
				d.swapColumn(i, index)
			}
			d.divide(i)
		}
	}
}

func (d *DiophantineSystem) isZero(i int) bool {
	for j := i + 1; j < d.m-1; j++ {
		if d.matrix[i][j] != 0 {
			return false
		}
	}
	return true
}

func (d *DiophantineSystem) divide(row int) {
	for i := row + 1; i < d.m-1; i++ {
		dVal := d.matrix[row][i] / d.matrix[row][row]
		d.subtract(i, row, dVal)
	}
}

func (d *DiophantineSystem) subtract(i, j, dVal int) {
	for k := 0; k < d.n; k++ {
		d.matrix[k][i] -= d.matrix[k][j] * dVal
	}
}

func (d *DiophantineSystem) inverse(i int) {
	for _, row := range d.matrix {
		row[i] = -row[i]
	}
}

func (d *DiophantineSystem) swapColumn(i, j int) {
	for k := 0; k < d.n; k++ {
		temp := d.matrix[k][i]
		d.matrix[k][i] = d.matrix[k][j]
		d.matrix[k][j] = temp
	}
}

func (d *DiophantineSystem) findMin(row []int, k int) int {
	absRow := make([]int, d.m-k-1)
	for i := k; i < d.m-1; i++ {
		absRow[i-k] = abs(row[i])
	}
	minEl := 0
	for _, x := range absRow {
		if x > 0 {
			minEl = x
			break
		}
	}
	index := -1
	for i, val := range absRow {
		if 0 < val && val < minEl {
			minEl = val
			index = i
		}
	}
	return index + k
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func (d *DiophantineSystem) divideLast() {
	for i := 0; i < d.n-d.m+1; i++ {
		if d.matrix[i][i] != 0 {
			k := d.matrix[i][d.m-1] / d.matrix[i][i]
			d.subtract(d.m-1, i, k)
		} else {
			d.r = i
			break
		}
		d.r = i + 1
	}
}

func (d *DiophantineSystem) getSolution() {
	for i := 0; i < d.n-d.m+1; i++ {
		if d.matrix[i][d.m-1] != 0 {
			panic("System is not solvable")
		}
	}
	for j := d.r; j < d.m; j++ {
		x := make([]int, 0)
		for i := d.n - d.m + 1; i < d.n; i++ {
			x = append(x, d.matrix[i][j])
		}
		d.x = append(d.x, x)
	}
}

func (d *DiophantineSystem) printSolutions() {
	n := d.m - 1
	s := len(d.x)
	for i := 0; i < n; i++ {
		fmt.Printf("%d\t", d.x[s-1][i])
		for j := 0; j < s-1; j++ {
			fmt.Printf("%d\t", d.x[j][i])
		}
		fmt.Println()
	}
	fmt.Print("x0\t")
	for i := 1; i < s; i++ {
		fmt.Printf("t%d\t", i)
	}
	fmt.Println()
}
