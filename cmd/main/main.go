package main

import (
	"fmt"
	"flag"
	"ssn"
)

var n = flag.Int("ssns", 3, "number of ssns to create")

func main() {
	flag.Parse()
	fmt.Print(ssn.Gen(*n))
}
