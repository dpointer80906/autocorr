package autocorr

import (
	"testing"
)

func TestInitConfig(t *testing.T) {
	config, err := initConfig("config.json")
	if err != nil {
		t.Log("err is not nil")
		t.Fail()
	}
	if config.Total != 3000 {
		t.Log("invalid Total")
		t.Fail()
	}
	if config.Window != 30 {
		t.Log("invalid Window")
		t.Fail()
	}
	config, err = initConfig("invalid.json")
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Log("err is nil")
		t.Fail()
	}

}

func TestInitData(t *testing.T) {
	config := Config{Total: 10, Window: 1}
	yData := initData(config, 1)
	if len(yData) != config.Total {
		t.Fail()
	}
}

func TestInitalSum(t *testing.T) {
	config := Config{Total: 10, Window: 2}
	yData := initData(config, 1)
	actual := sumWindow(yData[0:config.Window])
	expected := yData[0] + yData[1]
	if !floatEquals(expected, actual) {
		t.Fail()
	}
}

func TestRunningSum(t *testing.T) {
	config := Config{Total: 10, Window: 3}
	yData := make([]float64, config.Total)
	for i := range yData {
		yData[i] = float64(i)
	}
	ySum := sumWindow(yData[0:config.Window])
	for i := 0; i < config.Total-config.Window-1; i++ {
		ySum = runningSum(yData[i:i+config.Window+1], ySum)
		expected := sumWindow(yData[i+1 : i+config.Window+1])
		if expected != ySum {
			t.Fail()
		}
	}

}

// compare actual to expected autocorrelation value with default random seed
func TestLag1Autocorr(t *testing.T) {
	var expected = -0.029174179033668285

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
