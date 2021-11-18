package ast

import (
	"errors"
	"fmt"
)

var (
	NilParse = errors.New("nil parse tree")
)

type PartialParse struct {
	Position int
	Rest     []byte
}

func (p *PartialParse) Error() string {
	return fmt.Sprintf("tokens left after parse: pos %d, text=%v", p.Position, p.Rest)
}
