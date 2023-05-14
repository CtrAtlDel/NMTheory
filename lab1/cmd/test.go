package main

import "testing"

func Test(t *testing.T) {
	var x int64 = 3
	var y int64 = 5

	var expectedGcd int64 = 3
	var expectedLcm int64 = 15

	resultGcd := Gcd(x, y)
	resultLcm := Lcm(x, y)

	if resultGcd != expectedGcd {
		t.Errorf("MyFunction(2, 3) = %d; expected %d", resultGcd, expectedGcd)
	}

	if resultLcm != expectedLcm {
		t.Errorf("MyFunction(2, 3) = %d; expected %d", resultLcm, expectedLcm)
	}
}
