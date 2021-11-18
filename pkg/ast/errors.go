package ast

import (
	"errors"
	"fmt"
)

var (
	ErrNilParse = errors.New("nil parse tree")
)

type ErrPartialParse struct {
	Position int
	Rest     []byte
}

func (p *ErrPartialParse) Error() string {
	return fmt.Sprintf("tokens left after parse: pos %d, text=%v", p.Position, p.Rest)
}
