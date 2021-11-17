package ast

import "fmt"

type NilParse struct{}

func (n *NilParse) Error() string {
	return "nil parse tree"
}

type PartialParse struct {
	Position int
	Rest     []byte
}

func (p *PartialParse) Error() string {
	return fmt.Sprintf("tokens left after parse: pos %d, text=%v", p.Position, p.Rest)
}
