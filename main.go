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
	ast, y := ast.GetParser()

	s := parsec.NewScanner([]byte(samples.Kamstrup_Multical66))
	s.SetWSPattern(`^[ \t\r\n\x02\x03]+`)

	root, rest := ast.Parsewith(y, s)

	if root == nil {
		panic("Nil parse!")
	}
	pp(root, ``, true)
	fmt.Printf("endof? %v\n", rest.Endof())
}
