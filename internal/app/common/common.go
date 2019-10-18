package common

import (
	"path/filepath"
	"time"
)

//FileExt возвращает расширение файла
func FileExt(path string) string {
	g := filepath.Ext(path)
	return g
}

//DateTimeNowString Дата в формате
func DateTimeNowString() string {
	t := time.Now().Local()
	s := t.Format("2006-01-02")
	return s
}
