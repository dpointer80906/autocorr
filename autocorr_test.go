package autocorr

import (
	"testing"
)

// compare actual to expected autocorrelation value with default random seed
func TestLag1Autocorr(t *testing.T) {
	var expected = -0.02919829220151911

	actual := Example()
	if !floatEquals(expected, actual) {
		t.Fail()
	}
}

// credit: https://gist.github.com/cevaris/bc331cbe970b03816c6b (Adam Cardenas)
func floatEquals(a, b float64) bool {
	const EPSILON float64 = 0.00000001

	if (a-b) < EPSILON && (b-a) < EPSILON {
		return true
	}
	return false
}
