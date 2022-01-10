package ast

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jonkerj/go-iec62056/pkg/types"
	parsec "github.com/prataprc/goparsec"
)

var idSplitter *regexp.Regexp

func init() {
	idSplitter = regexp.MustCompile(`[-:.*]`)
}

func nodifyTelegram(ns []parsec.ParsecNode) parsec.ParsecNode {
	objs := make([]types.Object, 0)
	for _, obj := range ns[1].([]parsec.ParsecNode) {
		objs = append(objs, obj.(types.Object))
	}

	c := ""
	if _, ok := ns[2].(parsec.MaybeNone); ok == false {
		c = ns[2].(string)
	}

	return types.Telegram{
		Identification: ns[0].(*parsec.Terminal).GetValue(),
		Objects:        objs,
		Checksum:       c,
	}
}

func nodifyChecksum(ns []parsec.ParsecNode) parsec.ParsecNode {
	return ns[0].(*parsec.Terminal).GetValue()
}

func nodifyID(idS string) types.ID {
	parts := idSplitter.Split(idS, 7)
	ints := []byte{}

	for idx, str := range parts {
		if idx > 5 { // should not be more than 6. If so, it's an unknown format
			break
		}
		i, err := strconv.ParseInt(str, 10, 8)
		if err != nil {
			panic(fmt.Errorf("error converting ID to byte: %w", err))
		}
		ints = append(ints, byte(i))
	}

	switch len(parts) {
	case 2:
		return types.ID{
			A: nil,
			B: nil,
			C: &ints[0],
			D: &ints[1],
			E: nil,
			F: nil,
		}
	case 5:
		return types.ID{
			A: &ints[0],
			B: &ints[1],
			C: &ints[2],
			D: &ints[3],
			E: &ints[4],
		}
	case 6:
		return types.ID{
			A: &ints[0],
			B: &ints[1],
			C: &ints[2],
			D: &ints[3],
			E: &ints[4],
			F: &ints[5],
		}
	default:
		panic(fmt.Sprintf("don't know how to process ID field with %d values", len(parts)))
	}
}

func nodifyIDOnly(ns []parsec.ParsecNode) parsec.ParsecNode {
	return types.Object{
		ID:        nodifyID(ns[0].(*parsec.Terminal).GetValue()),
		Value:     types.Value{Value: nil, Unit: nil},
		Timestamp: nil,
	}
}

func nodifyCosem(ns []parsec.ParsecNode) parsec.ParsecNode {
	return types.Object{
		ID:        nodifyID(ns[0].(*parsec.Terminal).GetValue()),
		Value:     ns[1].(types.Value),
		Timestamp: nil,
	}
}

func nodifyCosemEmpty(ns []parsec.ParsecNode) parsec.ParsecNode {
	return types.Value{
		Value: nil,
		Unit:  nil,
	}
}

func nodifyCosemValue(ns []parsec.ParsecNode) parsec.ParsecNode {
	val := ns[1].(*parsec.Terminal).GetValue()
	return types.Value{
		Value: &val,
		Unit:  nil,
	}
}

func nodifyCosemValueUnit(ns []parsec.ParsecNode) parsec.ParsecNode {
	val := ns[1].(*parsec.Terminal).GetValue()
	unit := ns[3].(*parsec.Terminal).GetValue()
	return types.Value{
		Value: &val,
		Unit:  &unit,
	}
}

func nodifyDSMR3Gas(ns []parsec.ParsecNode) parsec.ParsecNode {
	ts := ns[2].(*parsec.Terminal).GetValue()

	u := ns[17].(*parsec.Terminal).GetValue()
	v := ns[20].(*parsec.Terminal).GetValue()

	val := types.Value{
		Value: &v,
		Unit:  &u,
	}

	return types.Object{
		ID:        nodifyID(ns[0].(*parsec.Terminal).GetValue()),
		Value:     val,
		Timestamp: &ts,
	}
}

func nodifyMBus(ns []parsec.ParsecNode) parsec.ParsecNode {
	ts := ns[2].(*parsec.Terminal).GetValue()
	return types.Object{
		ID:        nodifyID(ns[0].(*parsec.Terminal).GetValue()),
		Value:     ns[4].(types.Value),
		Timestamp: &ts,
	}
}
func nodifyIdentification(ns []parsec.ParsecNode) parsec.ParsecNode {
	return ns[1]
}

func nodifyFirstItem(ns []parsec.ParsecNode) parsec.ParsecNode {
	return ns[0]
}

func nodifySecondItem(ns []parsec.ParsecNode) parsec.ParsecNode {
	return ns[1]
}
