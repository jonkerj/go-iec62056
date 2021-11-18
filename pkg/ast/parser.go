package ast

import (
	"bytes"

	"github.com/jonkerj/go-iec62056/pkg/types"
	parsec "github.com/prataprc/goparsec"
)

func Parse(raw []byte) (*types.Telegram, error) {
	y := makey()
	s := parsec.NewScanner(raw)
	s.SetWSPattern(`^[ \t\r\n\x02\x03]+`)

	root, rest := y(s)

	if root == nil {
		return nil, NilParse
	}

	if !rest.Endof() && !bytes.Equal(raw[rest.GetCursor():], []byte{'\r', '\n'}) {
		return nil, &PartialParse{Position: rest.GetCursor(), Rest: raw[rest.GetCursor():]}
	}

	tg := root.(types.Telegram)

	return &tg, nil
}
