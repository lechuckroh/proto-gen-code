package util

import (
	"strings"
	"text/template"
)

func NewTemplate(tplName, tplText string, funcMap template.FuncMap) (*template.Template, error) {
	return template.New(tplName).Funcs(funcMap).Parse(tplText)
}

// ParseEnumeratedKeyValues parses "key1=value1,key2=value2,..." formatted string
func ParseEnumeratedKeyValues(s string) map[string]string {
	result := make(map[string]string)

	for _, token := range strings.Split(s, ",") {
		kv := strings.SplitN(token, "=", 2)
		key := kv[0]
		value := kv[1]
		result[key] = value
	}

	return result
}

func KeyValueSlicesToMap(slices []string) map[string]string {
	result := make(map[string]string)

	for _, token := range slices {
		kv := strings.SplitN(token, "=", 2)
		key := kv[0]
		value := kv[1]
		result[key] = value
	}

	return result
}
