package ast

import (
	"testing"

	"github.com/jonkerj/go-iec62056/pkg/samples"
)

func TestSamples(t *testing.T) {
	ast := &Telegram{}
	for name, telegram := range samples.All {
		t.Run(name, func(t *testing.T) {
			err := IECParser.ParseString(name, telegram, ast)

			if err != nil {
				t.Errorf("failed to parse telegram: %v", err)
			}
		})
	}
}
