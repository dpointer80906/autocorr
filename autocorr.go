// lag 1 autocorrelation of a set of random numbers
package autocorr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
)

var yData []float64 // all data points

type Config struct {
	Total  int // total yData size
	Window int // window into yData (N)
}

var config Config // data from config.json file

func init() {

	// read jsoninit data file into config struct
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	json.Unmarshal(data, &config)

	// create Total random data points
	yData = make([]float64, config.Total)
	for i := range yData {
		yData[i] = rand.Float64()
	}
}

// Calculate the initial yData window's sum.
func initialSum(yWindow []float64) (ySum float64) {
	for i := range yWindow {
		ySum += yWindow[i]
	}
	return
}

// Adjust running (sliding) sum of yData slice: subtract first slice value,
// add the last slice value to surrent sum.
func runningSum(yWindow []float64, oldSum float64) (newSum float64) {
	newSum = oldSum + yWindow[len(yWindow)-1] - yWindow[0]
	return
}

// Estimate autocorrrelation function lag 1 for yData.
func AutocorrLag1() (autocorr float64) {
	var covariance float64 // autocorr numerator
	var variance float64   // autocorr denominator
	var yMean float64      // mean of yData window data

	ySum := initialSum(yData[0:config.Window])
	// calculate covariance, variance terms 0..N-1
	for i := 0; i < config.Total-config.Window-1; i++ {
		yMean = ySum / float64(config.Window)
		thisYSubMean := yData[i] - yMean
		nextYSubMean := yData[i+1] - yMean
		covariance += thisYSubMean * nextYSubMean
		variance += thisYSubMean * thisYSubMean
		ySum = runningSum(yData[i:i+config.Window], ySum)
	}
	// pick up final variance term N
	thisYSubMean := yData[config.Total-config.Window] - yMean
	variance += thisYSubMean * thisYSubMean
	autocorr = covariance / variance
	return
}
