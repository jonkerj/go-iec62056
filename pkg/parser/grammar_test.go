package parser

import (
	"bytes"
	"testing"

	"github.com/jonkerj/go-iec62056/pkg/samples"
)

func TestAllSamples(t *testing.T) {
	for name, telegram := range samples.All {
		b := *telegram
		i := bytes.LastIndexByte(b, byte('!'))

		t.Run(name, func(st *testing.T) {
			_, err := Parse(b[:i])

			if err != nil {
				st.Fatalf("parse error: %v", err)
			}
		})
	}
}
