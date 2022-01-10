package parser

import (
	"testing"

	"github.com/jonkerj/go-iec62056/pkg/samples"
)

func TestAllSamples(t *testing.T) {
	for name, telegram := range samples.All {
		t.Run(name, func(st *testing.T) {
			_, err := Parse([]byte(telegram))

			if err != nil {
				st.Fatalf("parse error: %v", err)
			}
		})
	}
}
