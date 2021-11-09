package ast

import (
	"fmt"
	"io"
	"strings"
)

type Node interface {
	Data() interface{}
	Children() []Node
}

func PrettyPrint(n Node, w io.Writer) error {
	return recursivePrint(w, 0, n)
}

func recursivePrint(w io.Writer, indent int, node Node) error {
	_, err := fmt.Fprintf(w, "%s%v\n", strings.Repeat("  ", indent), node.Data())
	if err != nil {
		return err
	}

	for _, child := range node.Children() {
		err = recursivePrint(w, indent+1, child)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t Telegram) Children() []Node {
	children := []Node{}

	for _, object := range t.Objects {
		children = append(children, object)
	}

	return children
}

func (t Telegram) Data() interface{} {
	return t.Identification
}

func (o Object) Children() []Node {
	switch {
	case o.Cosem != nil:
		return o.Cosem.Children()
	case o.DSMR3Gas != nil:
		return o.DSMR3Gas.Children()
	case o.MBus != nil:
		return o.MBus.Children()
	case o.ProfileGeneric != nil:
		return o.ProfileGeneric.Children()
	}
	return []Node{}
}

func (o Object) Data() interface{} {
	switch {
	case o.Cosem != nil:
		return o.Cosem.Data()
	case o.DSMR3Gas != nil:
		return o.DSMR3Gas.Data()
	case o.MBus != nil:
		return o.MBus.Data()
	case o.ProfileGeneric != nil:
		return o.ProfileGeneric.Data()
	}
	return "object without name"
}

func (c Cosem) Children() []Node {
	return []Node{}
}

func (c Cosem) Data() interface{} {
	switch {
	case c.Unit != nil && c.Value != nil:
		return fmt.Sprintf("%s: %s %s", c.ID, *c.Value, *c.Unit)
	case c.Unit == nil && c.Value != nil:
		return fmt.Sprintf("%s: %s", c.ID, *c.Value)
	case c.Unit == nil && c.Value == nil:
		return c.ID
	}
	return fmt.Sprintf("Unknown Cosem %s", c.ID)
}

func (d DSMR3Gas) Children() []Node {
	return []Node{}
}

func (d DSMR3Gas) Data() interface{} {
	return fmt.Sprintf("%s: %s %s", d.ID, d.Value, d.Unit)
}

func (m MBus) Children() []Node {
	return []Node{}
}

func (m MBus) Data() interface{} {
	switch {
	case m.Unit != nil && m.Value != nil:
		return fmt.Sprintf("%s: %s %s", m.ID, *m.Value, *m.Unit)
	case m.Unit == nil && m.Value != nil:
		return fmt.Sprintf("%s: %s", m.ID, *m.Value)
	case m.Unit == nil && m.Value == nil:
		return m.ID
	}
	return fmt.Sprintf("Unknown Cosem %s", m.ID)
}

func (p ProfileGeneric) Children() []Node {
	children := []Node{}

	for _, logentry := range p.LogEntries {
		children = append(children, logentry)
	}

	return children
}

func (p ProfileGeneric) Data() interface{} {
	return fmt.Sprintf("%s: %s", p.ID, p.Code)
}

func (l LogEntry) Children() []Node {
	return []Node{}
}

func (l LogEntry) Data() interface{} {
	switch {
	case l.Unit != nil && l.Value != nil:
		return fmt.Sprintf("%s: %s %s", l.Timestamp, *l.Value, *l.Unit)
	case l.Unit == nil && l.Value != nil:
		return fmt.Sprintf("%s: %s", l.Timestamp, *l.Value)
	case l.Unit == nil && l.Value == nil:
		return l.Timestamp
	}
	return fmt.Sprintf("Unknown Cosem %s", l.Timestamp)
}
