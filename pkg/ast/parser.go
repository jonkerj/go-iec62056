package ast

import (
	participle "github.com/alecthomas/participle/v2"
)

type Telegram struct {
	Identification string    `"/" @IDString EOL EOL`
	Objects        []*Object `@@*`
	Checksum       *string   `"!" @Data* EOL`
}

type Object struct {
	Cosem          *Cosem          ` @@ EOL`
	MBus           *MBus           `|@@ EOL`
	DSMR3Gas       *DSMR3Gas       `|@@ EOL`
	ProfileGeneric *ProfileGeneric `|@@ EOL`
}

type Cosem struct {
	ID    string  `@Code`
	Value *string `"(" @String?`
	Unit  *string `("*" @String)? ")"`
}

type MBus struct {
	ID        string  `@Code`
	Timestamp string  `"(" @String ")"`
	Value     *string `"(" @String`
	Unit      *string `("*" @String)? ")"`
}

type DSMR3Gas struct {
	ID        string `@Code`
	TimeStamp string `"(" @String ")"`
	Type      string `"(" @String ")"`
	Id        string `"(" @String ")"`
	Switch    string `"(" @String ")"`
	Code      string `"(" @String ")"`
	Unit      string `"(" @String ")" EOL`
	Value     string `"(" @String ")"`
}

type ProfileGeneric struct {
	ID         string    `@Code`
	Length     string    `"(" @String ")"`
	Code       string    `"(" @String ")"`
	LogEntries *LogEntry `@@*`
}

type LogEntry struct {
	Timestamp string `"(" @String ")"`
	Value     string `"(" @String?`
	Unit      string `("*" @String)? ")"`
}

var (
	IECParser = participle.MustBuild(&Telegram{},
		participle.Lexer(IECLexer),
		participle.UseLookahead(20))
)
