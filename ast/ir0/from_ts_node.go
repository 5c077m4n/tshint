package ir0

import (
	"errors"
	"log/slog"
	"strconv"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
)

var ErrConvertTSToIR0Node = errors.New("could not convert this TS node to an AST one")

func From(node *tree_sitter.Node, sourceCode []byte) (Node, error) {
	start, end := node.ByteRange()

	byteRange := [2]uint{start, end}
	content := string(sourceCode[start:end])

	switch node.Kind() {
	case "identifier":
		return &Identifier{
			Name:      content,
			IsConst:   false,
			Type:      nil,
			Value:     nil,
			ByteRange: byteRange,
		}, nil
	case "number":
		if content == "number" {
			return &Number{ByteRange: byteRange}, nil
		}

		parsedNumber, err := strconv.ParseFloat(content, 32)
		if err != nil {
			return nil, errors.Join(ErrConvertTSToIR0Node, err)
		}

		return &NumberPrimitive{Value: float32(parsedNumber), ByteRange: byteRange}, nil
	case "boolean":
		if content == "boolean" {
			return &Boolean{ByteRange: byteRange}, nil
		}

		return &BooleanPrimitive{Value: content == "true", ByteRange: byteRange}, nil
	case "true":
		return &BooleanPrimitive{Value: true, ByteRange: byteRange}, nil
	case "false":
		return &BooleanPrimitive{Value: false, ByteRange: byteRange}, nil
	case "string":
		if content == "string" {
			return &String{ByteRange: byteRange}, nil
		}

		return &StringPrimitive{Value: content, ByteRange: byteRange}, nil
	default:
		slog.Warn(
			"Unhandled node",
			"kind", node.Kind(),
			"content", sourceCode[start:end],
		)
		return nil, ErrConvertTSToIR0Node
	}
}
