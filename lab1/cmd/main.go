package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Gcd(x, y int64) int64 {
	if x == 0 {
		return y
	} else if y == 0 {
		return x
	} else if x > y {
		return Gcd(x%y, y)
	} else {
		return Gcd(y%x, x)
	}
}

func Lcm(x, y int64) int64 {
	return (x * y) / Gcd(x, y)
}

func readChar(v string) int64 {
	fmt.Printf("Input \"%s\":\n", v)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	trimmed := strings.TrimSpace(input)
	result, err := strconv.ParseInt(trimmed, 10, 64)

	if err != nil {
		fmt.Printf("Bad symbol use only int: %s\n", trimmed)
		return 0
	} else {
		fmt.Printf("Your integer input: %d\n", result)
		return result
	}
}

func main() {
	a := math.Abs(float64(readChar("a")))
    b := math.Abs(float64(readChar("b")))

    fmt.Printf("Least Common Multiple: %d\n", Lcm(int64(a), int64(b)))
    fmt.Printf("Greatest Common Divisor: %d\n", Gcd(int64(a), int64(b)))
}