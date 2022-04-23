package parser

import (
	"errors"
	"fmt"
	"strconv"
)

type Telegram struct {
	Identification string
	Objects        []Object
}

func (t Telegram) String() string {
	return fmt.Sprintf("ID %s, %d objects", t.Identification, len(t.Objects))
}

type Object struct {
	ID        ID
	Value     Value
	Timestamp *string
}

func (o Object) String() string {
	return fmt.Sprintf("ID %s %s", o.ID, o.Value)
}

type Value interface {
	GetNumericValue() (*float64, error)
	GetStringValue() (*string, error)
}

func NewValue(raw string, unit *string) Value {
	num, err := strconv.ParseFloat(raw, 64)
	if err == nil {
		return NumericValue{Value: num, Unit: unit}
	}
	return StringValue{Value: raw, Unit: unit}
}

type NumericValue struct {
	Value float64
	Unit  *string
}

func (n NumericValue) String() string {
	if n.Unit != nil {
		return fmt.Sprintf("%v %s", n.Value, *n.Unit)
	} else {
		return fmt.Sprintf("%v", n.Value)
	}
}

func (n NumericValue) GetNumericValue() (*float64, error) {
	return &n.Value, nil
}

func (n NumericValue) GetStringValue() (*string, error) {
	return nil, errors.New("not implemented")
}

type StringValue struct {
	Value string
	Unit  *string
}

func (s StringValue) String() string {
	if s.Unit != nil {
		return fmt.Sprintf("%s %s", s.Value, *s.Unit)
	} else {
		return s.Value
	}
}

func (s StringValue) GetNumericValue() (*float64, error) {
	return nil, errors.New("not implemented")
}

func (s StringValue) GetStringValue() (*string, error) {
	return &s.Value, nil
}

type NullValue struct {
}

func (n NullValue) String() string {
	return "n/a"
}

func (n NullValue) GetNumericValue() (*float64, error) {
	return nil, errors.New("not implemented")
}

func (n NullValue) GetStringValue() (*string, error) {
	return nil, errors.New("not implemented")
}

type ID struct {
	A *uint8
	B *uint8
	C *uint8
	D *uint8
	E *uint8
	F *uint8
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
