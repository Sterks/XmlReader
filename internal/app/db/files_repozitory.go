package db

import (
	"fmt"
	"time"

	model "github.com/Sterks/XmlReader/internal/app/models"
	"github.com/sirupsen/logrus"
)

// FilesRepository Структура репозиториев для файлов
type FilesRepository struct {
	db *PgDb
}

// Create создать запись в базе
func (r *FilesRepository) Create(f *model.FileInfo, ext string) (*model.FileInfo, error) {

	var extFile int
	_ = r.db.File().db.db.QueryRow(`select ft_id from "FilesTypes" where ft_ext = $1`, ext).Scan(&extFile)
	f.FileType = extFile

	// _, err := r.db.db.Exec(`insert into "Files" (f_name, f_type) values ($3, $5)`, f.FileName, f.FileType)
	// if err != nil {
	// 	logrus.Fatalf("Не могу добавить запись %v", err)
	// 	return nil, err
	// }
	// return f, err

	timestamp := time.Now()
	f.FileCreateAt = timestamp
	res, err := r.db.db.Exec(`insert into "Files" 
							   (f_name, f_type, f_area, f_size, f_date_create, f_date_create_from_source, f_fullpath, f_file_is_dir)
					 			values ($1, $2, $3, $4, $5, $6, $7, $8)`,
		f.FileName,
		f.FileType,
		f.FileArea,
		f.FileSize,
		f.FileCreateAt,
		f.FileDateMod,
		f.FilePath,
		f.FileIsDir,
	)
	if err != nil {
		logrus.Error(nil, err)
	}
	fmt.Println(res)
	return f, nil
}
