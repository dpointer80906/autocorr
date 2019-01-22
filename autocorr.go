// lag 1 autocorr
// http://qed.econ.queensu.ca/ETM/corrections/Third-pdf/pg564-5.pdf  eqn 13.26
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var dpWindow int    // data points window from yData (N)
var dpTotal int     // data points total
var yData []float64 // all data points

func init() {
	dpWindow = 30
	dpTotal = 3000

	filename := dataFilename()
	fmt.Printf("%v", filename)

	yData = make([]float64, dpTotal)
	for i := range yData {
		yData[i] = rand.Float64()
	}
}

func main() {
	autocorr := AutocorrLag1()
	fmt.Printf("%v\n", autocorr)
}

func dataParams() (window, total int) {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Printf("number of random data points to autocorrelate? ")
	filename, _ = inputReader.R('\n')
	return
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

	ySum := initialSum(yData[0:dpWindow])
	// calculate covariance, variance terms 0..N-1
	for i := 0; i < dpTotal-dpWindow-1; i++ {
		yMean = ySum / float64(dpWindow)
		thisYSubMean := yData[i] - yMean
		nextYSubMean := yData[i+1] - yMean
		covariance += thisYSubMean * nextYSubMean
		variance += thisYSubMean * thisYSubMean
		ySum = runningSum(yData[i:i+dpWindow], ySum)
	}
	// pick up final variance term N
	thisYSubMean := yData[dpTotal-dpWindow] - yMean
	variance += thisYSubMean * thisYSubMean
	autocorr = covariance / variance
	return
}
