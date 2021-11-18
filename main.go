package main

import (
	"fmt"

	"github.com/jonkerj/go-iec62056/pkg/ast"
	"github.com/jonkerj/go-iec62056/pkg/samples"
	parsec "github.com/prataprc/goparsec"
)

func pp(q parsec.Queryable, indent string, last bool) {
	fmt.Print(indent)

	if last {
		fmt.Print(`\-`)
		indent += `  `
	} else {
		fmt.Print(`|-`)
		indent += `| `
	}
	fmt.Printf("%s: %s\n", q.GetName(), q.GetValue())
	for i, child := range q.GetChildren() {
		pp(child, indent, i == (len(q.GetChildren())-1))
	}
}

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
