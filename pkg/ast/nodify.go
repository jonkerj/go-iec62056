package ast

import (
	"github.com/jonkerj/go-iec62056/pkg/types"
	parsec "github.com/prataprc/goparsec"
)

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

func nodifyIDOnly(ns []parsec.ParsecNode) parsec.ParsecNode {
	return types.Object{
		ID:        ns[0].(*parsec.Terminal).GetValue(),
		Value:     types.Value{Value: nil, Unit: nil},
		Timestamp: nil,
	}
}

func nodifyCosem(ns []parsec.ParsecNode) parsec.ParsecNode {
	return types.Object{
		ID:        ns[0].(*parsec.Terminal).GetValue(),
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
		ID:        ns[0].(*parsec.Terminal).GetValue(),
		Value:     val,
		Timestamp: &ts,
	}
}

func nodifyMBus(ns []parsec.ParsecNode) parsec.ParsecNode {
	ts := ns[2].(*parsec.Terminal).GetValue()
	return types.Object{
		ID:        ns[0].(*parsec.Terminal).GetValue(),
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
