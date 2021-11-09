package ast

import (
	"strings"
	"testing"

	"github.com/jonkerj/go-iec62056/pkg/samples"
)

func TestPretty(t *testing.T) {
	tg := Telegram{}
	err := IECParser.ParseString("kaifa_ma105", samples.Kaifa_MA105, &tg)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	var sb strings.Builder
	err = PrettyPrint(tg, &sb)
	if err != nil {
		t.Errorf("error during pretty print: %v", err)
	}
}
