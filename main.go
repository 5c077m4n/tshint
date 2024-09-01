// Package main
package main

import (
	"flag"
	"tshint/walker"
)

func main() {
	source := flag.String("eval", "", "Source code to evaluate")
	flag.Parse()

	if err := walker.Walk([]byte(*source)); err != nil {
		panic(err)
	}
}
