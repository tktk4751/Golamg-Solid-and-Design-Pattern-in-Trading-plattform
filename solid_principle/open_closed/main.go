package main

import (
	"fmt"
	"solid/solid_principle/open_closed/bad"
	"solid/solid_principle/open_closed/good"
)

func main() {
	fmt.Println("BAD")
	bad.Run()

	fmt.Println("GOOD")
	good.Run()
}
