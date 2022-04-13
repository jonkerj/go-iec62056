package parser

import "fmt"

func (t Telegram) String() string {
	return fmt.Sprintf("ID %s, %d objects", t.Identification, len(t.Objects))
}

func (v Value) String() string {
	switch {
	case v.Value != nil && v.Unit != nil:
		return fmt.Sprintf("%s %s", *v.Value, *v.Unit)
	case v.Value != nil:
		return fmt.Sprintf("%s", *v.Value)
	}
	return "n/a"
}

func (o Object) String() string {
	return fmt.Sprintf("ID %s %s", o.ID, o.Value)
}

func (i ID) String() string {
	switch {
	case i.A != nil && i.B != nil && i.C != nil && i.D != nil && i.E != nil && i.F != nil:
		return fmt.Sprintf("%d-%d:%d.%d.%d*%d", *i.A, *i.B, *i.C, *i.D, *i.E, *i.F)
	case i.A != nil && i.B != nil && i.C != nil && i.D != nil && i.E != nil && i.F == nil:
		return fmt.Sprintf("%d-%d:%d.%d.%d", *i.A, *i.B, *i.C, *i.D, *i.E)
	case i.A == nil && i.B == nil && i.C != nil && i.D != nil && i.E == nil && i.F == nil:
		return fmt.Sprintf("%d.%d", *i.C, *i.D)
	}

	panic("Cannot convert this ID format to string")
}
