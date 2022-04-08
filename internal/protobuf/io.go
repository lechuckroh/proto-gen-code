package protobuf

import (
	"os"

	"github.com/yoheimuta/go-protoparser/v4"
	"github.com/yoheimuta/go-protoparser/v4/parser"
)

func LoadProtoFile(filename string) (*parser.Proto, error) {
	reader, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return protoparser.Parse(reader, protoparser.WithPermissive(true), protoparser.WithFilename(filename))
}
