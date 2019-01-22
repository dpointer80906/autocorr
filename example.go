package autocorr

import (
	"fmt"

	"./autocorr"
)

func ExampleAutocorr() {
	result := autocorr.AutocorrLag1()
	fmt.Printf("%v\n", result)
}
