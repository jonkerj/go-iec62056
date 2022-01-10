package parser

import "fmt"

func (t Telegram) String() string {
	return fmt.Sprintf("ID %s, %d objects, checksum=%v", t.Identification, len(t.Objects), t.Checksum)
}

func (o Object) String() string {
	switch {
	case o.Value.Value != nil && o.Value.Unit != nil:
		return fmt.Sprintf("ID %s %s %s", o.ID, *o.Value.Value, *o.Value.Unit)
	case o.Value.Value != nil:
		return fmt.Sprintf("ID %s %s", o.ID, *o.Value.Value)
	}
	return fmt.Sprintf("ID %s", o.ID)
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
