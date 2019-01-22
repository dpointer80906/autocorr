package autocorr

import (
	"fmt"

	"github.com/dpointer80906/autocorr"
)

func ExampleAutocorr() {
	result := autocorr.AutocorrLag1()
	fmt.Printf("%v\n", result)
}
