package main

import (
	"fmt"
	bootcheck "golang-studies/general/bootcheck"
)

func main() {
	// test := 25

	fmt.Println(bootcheck.EnvironmentOK("1.1.1"))

	fmt.Println("go studies ready")
}
