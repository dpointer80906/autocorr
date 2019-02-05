// lag 1 autocorrelation of a set of random numbers
package autocorr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
)

// configuration data read from file
type Config struct {
	Total  int // total yData size
	Window int // window into yData (N)
}

func Example() (result float64) {
	result = Lag1Autocorr()
	fmt.Printf("%v\n", result)
	return
}

// Estimate autocorrelation function lag 1 for yData.
func Lag1Autocorr() (acorr float64) {
	var covariance float64 // autocorr numerator
	var variance float64   // autocorr denominator
	var yMean float64      // mean of yData window data

	// load config, init random data
	// TODO: supply config file name from command line
	config, err := initConfig("config.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// TODO: provide capability to optionally specify seed from command line
	yData := initData(config, 1)
	ySum := initialSum(yData[0:config.Window])

	// calculate covariance, variance terms 0..N-1
	for i := 0; i < config.Total-config.Window-1; i++ {
		yMean = ySum / float64(config.Window)
		thisYMean := yData[i] - yMean
		nextYMean := yData[i+1] - yMean
		covariance += thisYMean * nextYMean
		variance += thisYMean * thisYMean
		ySum = runningSum(yData[i:i+config.Window], ySum)
	}
	// pick up final variance term N
	thisYMean := yData[config.Total-config.Window] - yMean
	variance += thisYMean * thisYMean
	acorr = covariance / variance
	return
}

// initialize configuration data from file data
func initConfig(filename string) (config Config, err error) {

	// read json init data file into config struct
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	} else {
		err = json.Unmarshal(data, &config)
		if err != nil {
			return
		}
	}
	return
}

// create Total random data points
func initData(config Config, seed int64) (yData []float64) {
	rand.Seed(seed)
	yData = make([]float64, config.Total)
	for i := range yData {
		yData[i] = rand.Float64()
	}
	return
}

// calculate the initial yData window's sum.
func initialSum(yWindow []float64) (ySum float64) {
	for i := range yWindow {
		ySum += yWindow[i]
	}
	return
}

// Adjust running (sliding) sum of yData slice: subtract first slice value,
// add the last slice value to current sum.
func runningSum(yWindow []float64, oldSum float64) (newSum float64) {
	newSum = oldSum + yWindow[len(yWindow)-1] - yWindow[0]
	fmt.Printf("%v %v %v %v\n", newSum, oldSum, yWindow[len(yWindow)-1], yWindow[0])
	return
}
