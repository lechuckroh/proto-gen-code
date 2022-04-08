package util

import (
	"path/filepath"
	"strings"
)

// GetBaseFilename 파일경로에서 디렉토리와 확장자를 제외한 파일 이름만 반환합니다.
func GetBaseFilename(path string) string {
	return strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
}
