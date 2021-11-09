package ast

import (
	"github.com/alecthomas/participle/v2/lexer"
)

var (
	IECLexer = lexer.MustStateful(lexer.Rules{
		"Root": {
			{`Code`, `([0-9]+-[0-9]+:[0-9]+.)?[0-9]+.[0-9]+`, nil},
			{`Identification`, `/`, lexer.Push("Identification")},
			{`EOL`, `\r\n`, nil},
			{`Value`, `\(`, lexer.Push("Value")},
			{`Checksum`, `!`, lexer.Push("Checksum")},
		},
		"Identification": {
			{`IDString`, `[a-zA-Z0-9.\-_,:\x02\\ ]+`, nil},
			{`EOL`, `\r\n`, lexer.Pop()},
		},
		"Value": {
			{`ValueEnd`, `\)`, lexer.Pop()},
			{`String`, `[a-zA-Z0-9.\-_,:]+`, nil},
			{`Seperator`, `\*`, nil},
		},
		"Checksum": {
			{`EOL`, `\r\n`, nil},
			{`Data`, `.`, nil},
		},
	})
)
