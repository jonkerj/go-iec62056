package ast

import (
	"bytes"

	parsec "github.com/prataprc/goparsec"
)

func Parse(telegram []byte) (parsec.Queryable, error) {
	ast := parsec.NewAST("iec62056", 100)
	y := makey(ast)
	s := parsec.NewScanner(telegram)
	s.SetWSPattern(`^[ \t\r\n\x02\x03]+`)

	root, rest := ast.Parsewith(y, s)

	if root == nil {
		return nil, &NilParse{}
	}

	if !rest.Endof() && !bytes.Equal(telegram[rest.GetCursor():], []byte{'\r', '\n'}) {
		return nil, &PartialParse{Position: rest.GetCursor(), Rest: telegram[rest.GetCursor():]}
	}

	return root, nil
}
