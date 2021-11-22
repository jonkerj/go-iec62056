package main

import (
	"fmt"

	"github.com/jonkerj/go-iec62056/pkg/ast"
	"github.com/jonkerj/go-iec62056/pkg/samples"
)

func main() {
	telegram, err := ast.Parse([]byte(samples.IskraMT382_1000_DSMRv5))

	if err != nil {
		panic(fmt.Sprintf("Parse error: %v", err))
	}

	fmt.Printf("telegram: %v\n", telegram)
	for _, obj := range telegram.Objects {
		fmt.Printf("obj: %v\n", obj)
	}
}
