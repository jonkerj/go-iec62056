package ast

import (
	"testing"

	"github.com/jonkerj/go-iec62056/pkg/samples"
	parsec "github.com/prataprc/goparsec"
)

func TestAllSamples(t *testing.T) {
	ast, y := GetParser()
	for name, telegram := range samples.All {
		t.Run(name, func(st *testing.T) {
			s := parsec.NewScanner([]byte(telegram))
			s.SetWSPattern(`^[ \t\r\n\x02\x03]+`)

			root, rest := ast.Parsewith(y, s)

			if root == nil {
				st.Fatalf("nil parse tree")
			}

			if !rest.Endof() && telegram[rest.GetCursor():] != "\r\n" {
				st.Errorf("%d runes left that are not CRLF", len(telegram)-rest.GetCursor())
			}
		})
	}
}
