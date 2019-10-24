package model

import "time"

// FileInfo ...
type FileInfo struct {
	ID          int
	FileParent  string
	FileName    string
	FilePath    string
	FileSize    int64
	FileIsDir   bool
	FileDateMod time.Time
	FileArea    string
}
