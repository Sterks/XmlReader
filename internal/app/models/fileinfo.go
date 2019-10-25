package model

import "time"

// FileInfo ...
type FileInfo struct {
	ID           int
	FileParent   string
	FileName     string
	FilePath     string
	FileSize     int64
	FileIsDir    bool
	FileDateMod  time.Time
	FileCreateAt time.Time
	FileArea     string
	FileType     int
	FileHash     string
}
