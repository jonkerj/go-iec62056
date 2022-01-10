package parser

import (
	"fmt"
	"regexp"
	"strconv"

	parsec "github.com/prataprc/goparsec"
)

var idSplitter *regexp.Regexp

func init() {
	idSplitter = regexp.MustCompile(`[-:.*]`)
}

func nodifyTelegram(ns []parsec.ParsecNode) parsec.ParsecNode {
	objs := make([]Object, 0)
	for _, obj := range ns[1].([]parsec.ParsecNode) {
		objs = append(objs, obj.(Object))
	}

	return Telegram{
		Identification: ns[0].(*parsec.Terminal).GetValue(),
		Objects:        objs,
	}
}

func nodifyID(idS string) ID {
	parts := idSplitter.Split(idS, 7)
	ints := []byte{}

	for idx, str := range parts {
		if idx > 5 { // should not be more than 6. If so, it's an unknown format
			break
		}
		i, err := strconv.ParseUint(str, 10, 8)
		if err != nil {
			panic(fmt.Errorf("error converting ID to byte: %w", err))
		}
		ints = append(ints, byte(i))
	}

	switch len(parts) {
	case 2:
		return ID{
			A: nil,
			B: nil,
			C: &ints[0],
			D: &ints[1],
			E: nil,
			F: nil,
		}
	case 5:
		return ID{
			A: &ints[0],
			B: &ints[1],
			C: &ints[2],
			D: &ints[3],
			E: &ints[4],
		}
	case 6:
		return ID{
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
	return Object{
		ID:        nodifyID(ns[0].(*parsec.Terminal).GetValue()),
		Value:     Value{Value: nil, Unit: nil},
		Timestamp: nil,
	}
}

func nodifyCosem(ns []parsec.ParsecNode) parsec.ParsecNode {
	return Object{
		ID:        nodifyID(ns[0].(*parsec.Terminal).GetValue()),
		Value:     ns[1].(Value),
		Timestamp: nil,
	}
}

func nodifyCosemEmpty(ns []parsec.ParsecNode) parsec.ParsecNode {
	return Value{
		Value: nil,
		Unit:  nil,
	}
}

func nodifyCosemValue(ns []parsec.ParsecNode) parsec.ParsecNode {
	val := ns[1].(*parsec.Terminal).GetValue()
	return Value{
		Value: &val,
		Unit:  nil,
	}
}

func nodifyCosemValueUnit(ns []parsec.ParsecNode) parsec.ParsecNode {
	val := ns[1].(*parsec.Terminal).GetValue()
	unit := ns[3].(*parsec.Terminal).GetValue()
	return Value{
		Value: &val,
		Unit:  &unit,
	}
}

func nodifyDSMR3Gas(ns []parsec.ParsecNode) parsec.ParsecNode {
	ts := ns[2].(*parsec.Terminal).GetValue()

	u := ns[17].(*parsec.Terminal).GetValue()
	v := ns[20].(*parsec.Terminal).GetValue()

	val := Value{
		Value: &v,
		Unit:  &u,
	}

	return Object{
		ID:        nodifyID(ns[0].(*parsec.Terminal).GetValue()),
		Value:     val,
		Timestamp: &ts,
	}
}

func nodifyMBus(ns []parsec.ParsecNode) parsec.ParsecNode {
	ts := ns[2].(*parsec.Terminal).GetValue()
	return Object{
		ID:        nodifyID(ns[0].(*parsec.Terminal).GetValue()),
		Value:     ns[4].(Value),
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
