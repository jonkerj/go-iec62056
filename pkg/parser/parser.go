package parser

import (
	"bytes"

	parsec "github.com/prataprc/goparsec"
)

func Parse(raw []byte) (*Telegram, error) {
	y := makey()
	s := parsec.NewScanner(raw)
	s.SetWSPattern(`^[ \t\r\n\x02\x03]+`)

	root, rest := y(s)

	if root == nil {
		return nil, ErrNilParse
	}

	if !rest.Endof() && !bytes.Equal(raw[rest.GetCursor():], []byte{'\r', '\n'}) {
		return nil, &ErrPartialParse{Position: rest.GetCursor(), Rest: raw[rest.GetCursor():]}
	}

	tg := root.(Telegram)

	return &tg, nil
}
