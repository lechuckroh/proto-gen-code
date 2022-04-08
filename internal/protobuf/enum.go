package protobuf

import (
	"github.com/lechuckroh/protogencode/internal/util/fp"
	parser "github.com/yoheimuta/go-protoparser/v4/parser"
)

type Enum interface {
	Fields() []EnumField
	Name() string
	SetName(string)
}

type EnumField interface {
	Comments() []string
	Enum() Enum
	Name() string
	Number() string
}

func NewEnum(parserEnum *parser.Enum) Enum {
	enum := &EnumImpl{
		source: parserEnum,
		name:   parserEnum.EnumName,
		fields: nil,
	}

	var enumFields []EnumField
	for _, visitee := range parserEnum.EnumBody {
		switch visitee.(type) {
		case *parser.EnumField:
			enumFields = append(enumFields, NewEnumField(enum, visitee.(*parser.EnumField)))
		}
	}
	enum.fields = enumFields

	return enum
}

func NewEnumField(enum Enum, enumField *parser.EnumField) EnumField {
	return &EnumFieldImpl{
		enum:   enum,
		source: enumField,
		comments: fp.Map(enumField.Comments, func(c *parser.Comment) string {
			return c.Raw
		}),
	}
}

// -------------------------------------------------
// EnumImpl
// -------------------------------------------------

type EnumImpl struct {
	source *parser.Enum
	name   string
	fields []EnumField
}

func (e *EnumImpl) Name() string {
	return e.name
}

func (e *EnumImpl) SetName(name string) {
	e.name = name
}

func (e *EnumImpl) Fields() []EnumField {
	return e.fields
}

// -------------------------------------------------
// EnumFieldImpl
// -------------------------------------------------

type EnumFieldImpl struct {
	comments []string
	enum     Enum
	source   *parser.EnumField
}

func (f *EnumFieldImpl) Comments() []string {
	return f.comments
}

func (f *EnumFieldImpl) Enum() Enum {
	return f.enum
}

func (f *EnumFieldImpl) Name() string {
	return f.source.Ident
}

func (f *EnumFieldImpl) Number() string {
	return f.source.Number
}
