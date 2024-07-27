package main

import (
	"github.com/lechuckroh/protogencode/internal/pkg/protobuf"
)

type GenContext struct {
	Proto         protobuf.Proto
	ConstFile     string
	MsgVarPostfix string
	MsgVarPrefix  string
}
