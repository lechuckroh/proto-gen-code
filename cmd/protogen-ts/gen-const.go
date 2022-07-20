package main

import (
	"github.com/lechuckroh/protogencode/internal/util"
	"io"
	"text/template"
)

// generateMessages 상수들을 정의한 *.ts 파일을 생성합니다.
func generateConstants(wr io.Writer, ctx *GenContext) error {
	funcMap := template.FuncMap{}
	tplText := `{{- range .Enums}}export const {{.Name}} = {
  {{- range .Fields}}
  {{.Name}}: {{.Number}},{{end}}
} as const;

{{end}}
`

	tpl, err := util.NewTemplate("constant", tplText, funcMap)
	if err != nil {
		return err
	}

	return tpl.Execute(wr, ctx.Proto)
}
