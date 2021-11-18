package types

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
