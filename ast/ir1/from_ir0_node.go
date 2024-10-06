package ir1

import (
	"errors"
	"tshint/ast/ir0"
)

var (
	ErrConvertHIRToMIR = errors.New("could not convert this HIR node to a MIR one")
	ErrUnimplemented   = errors.New("unimplemented")
)

func From(_node ir0.Node) (Node, error) {
	return nil, errors.Join(ErrConvertHIRToMIR, ErrUnimplemented)
}
