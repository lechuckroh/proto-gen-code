package main

import (
	"fmt"
	"github.com/lechuckroh/protogencode/internal/protobuf"
	"github.com/lechuckroh/protogencode/internal/util"
	"github.com/lechuckroh/protogencode/internal/util/fp"
	"io"
	"strings"
	"text/template"
)

// generateMessages 메시지를 정의한 *.ts 파일을 생성합니다.
func generateMessages(wr io.Writer, ctx *GenContext) error {
	proto := ctx.Proto

	funcMap := template.FuncMap{
		"fieldMessage": func(f protobuf.EnumField) string {
			comments := f.Comments()
			if len(comments) > 0 {
				return strings.Join(comments, "\n")
			}
			return f.Name()
		},
		"msgVarName": func(enum protobuf.Enum) string {
			prefix := ctx.MsgVarPrefix
			postfix := ctx.MsgVarPostfix
			// 미지정시 디폴트 값 사용
			if prefix == "" && postfix == "" {
				prefix = "Msg"
			}
			return fmt.Sprintf("%s%s%s", prefix, enum.Name(), postfix)
		},
	}
	tplText := `import { {{.importTypes}} } from "./{{.constFile}}";

{{range .proto.Enums}}export const {{msgVarName .}}: { [code: number]: string } = {
  {{- range .Fields}}
  [{{.Enum.Name}}.{{.Name}}]: "{{fieldMessage .}}",{{end}}
};
{{end}}
`

	tpl, err := util.NewTemplate("message", tplText, funcMap)
	if err != nil {
		return err
	}

	return tpl.Execute(wr, map[string]interface{}{
		"importTypes": strings.Join(getMessageImportTypes(proto.Enums()), ", "),
		"constFile":   util.GetBaseFilename(ctx.ConstFile),
		"proto":       proto,
	})
}

func getMessageImportTypes(enums []protobuf.Enum) []string {
	return fp.Map(enums, func(enum protobuf.Enum) string {
		return enum.Name()
	})
}
