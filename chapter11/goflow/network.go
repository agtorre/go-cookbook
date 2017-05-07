package goflow

import flow "github.com/trustmaster/goflow"

// EncodingApp creates a flow-based
// pipeline to encode and print the
// result
type EncodingApp struct {
	flow.Graph
}

// NewEncodingApp wires together the compoents
func NewEncodingApp() *EncodingApp {
	e := &EncodingApp{}
	e.InitGraphState()

	e.Add(&Encoder{}, "encoder")
	e.Add(&Printer{}, "printer")

	e.Connect("encoder", "Res", "printer", "Line")
	e.MapInPort("In", "encoder", "Val")

	return e
}
