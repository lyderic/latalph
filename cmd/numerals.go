package cmd

import (
	"fmt"

	rom "github.com/brandenc40/romannumeral"
)

func ExampleStringToInt() {
	integer, err := rom.StringToInt("IV")
	if err != nil {
		panic(err)
	}
	fmt.Println(integer == 4) // True
}
