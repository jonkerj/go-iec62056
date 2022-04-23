package main

import (
	"fmt"

	"github.com/jonkerj/go-iec62056/pkg/parser"
	"github.com/jonkerj/go-iec62056/pkg/samples"
)

func main() {
	telegram, err := parser.Parse(samples.Ziv_5CTA3[:714]) // 714 is the `!` in the datagram
	if err != nil {
		panic(fmt.Sprintf("parse error: %v", err))
	}

	fmt.Printf("telegram: %v\n", telegram)
	for _, obj := range telegram.Objects {
		fmt.Printf("obj: %v (%T)\n", obj, obj.Value)
	}
}
