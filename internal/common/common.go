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

//FromDate Вернуть время начала
func FromDate() time.Time {
	now := time.Now()
	y, m, d := now.Date()
	from := time.Date(y, m, d, 0, 0, 0, 0, now.Location())
	return from
}

//ToDate Вернуть дату окончания
func ToDate() time.Time {
	to := time.Now()
	return to
}
