package protobuf

import (
	fp2 "github.com/lechuckroh/protogencode/internal/pkg/util/fp"
	"github.com/yoheimuta/go-protoparser/v4/parser"
)

// Proto *.proto 파일의 내용을 보관하는 인터페이스
type Proto interface {
	Enums() []Enum
}

type ProtoOption struct {
	Includes []string
	Excludes []string
	NameMap  map[string]string
}

// NewProto 함수는 *.proto 에서 읽어온 내용을 사용해 Proto 인스턴스를 생성합니다.
func NewProto(p *parser.Proto, opt ProtoOption) Proto {
	var enums []Enum
	for _, visitee := range p.ProtoBody {
		switch visitee.(type) {
		case *parser.Enum:
			enums = append(enums, NewEnum(visitee.(*parser.Enum)))
		default:
			// 필요시 추가
		}
	}

	// post process
	enums = filterEnums(enums, opt.Includes, opt.Excludes)
	renameEnums(enums, opt.NameMap)

	return &ProtoImpl{enums: enums, option: opt}
}

type ProtoImpl struct {
	enums  []Enum
	option ProtoOption
}

func (p *ProtoImpl) Enums() []Enum {
	return p.enums
}

// filterEnums 조건에 맞는 enum 목록을 필터링합니다.
func filterEnums(enums []Enum, includes []string, excludes []string) []Enum {
	acceptNotExcluded := len(includes) == 0

	return fp2.Filter(enums, func(enum Enum) bool {
		predicate := func(s string) bool {
			return s == enum.Name()
		}
		if fp2.Any(includes, predicate) {
			return true
		}
		if fp2.Any(excludes, predicate) {
			return false
		}
		return acceptNotExcluded
	})
}

// renameEnums Enum의 이름을 변경합니다.
func renameEnums(enums []Enum, nameMap map[string]string) {
	for _, enum := range enums {
		if newName, exists := nameMap[enum.Name()]; exists && newName != "" {
			enum.SetName(newName)
		}
	}
}
