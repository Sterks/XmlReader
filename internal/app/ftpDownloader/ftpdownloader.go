package ftpdownloader

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"sync/atomic"
	"time"

	"github.com/Sterks/XmlReader/internal/app/configuration"
	"github.com/Sterks/XmlReader/internal/app/db"
	model "github.com/Sterks/XmlReader/internal/app/models"
	_ "github.com/lib/pq"
	"github.com/secsy/goftp"
	"github.com/sirupsen/logrus"
)

// FtpDownloader ...
type FtpDownloader struct {
	config *configuration.Configuration
	logger *logrus.Logger
	ftp    *goftp.Client
	db     *db.PgDb
}

//New ...
func New(con *configuration.Configuration) *FtpDownloader {
	return &FtpDownloader{
		config: con,
		logger: logrus.New(),
	}
}

// Start ...
func (f *FtpDownloader) Start() error {
	if err := f.ConfigureDb(); err != nil {
		return err
	}

	return nil
}

func (f *FtpDownloader) configureLogger() {
	level, err := logrus.ParseLevel(f.config.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	f.logger.SetLevel(level)
}

//Connect ...
func (f *FtpDownloader) Connect() (*goftp.Client, error) {
	con := goftp.Config{
		User:     "free",
		Password: "free",
	}
	ftp, err := goftp.DialConfig(con, f.config.FtpConnect)
	if err != nil {
		return nil, err
	}
	return ftp, nil
}

//GetFiles ...
func (f *FtpDownloader) GetFiles(client *goftp.Client, from time.Time, to time.Time) []model.FileInfo {
	rootPath := f.config.RootDir
	docType := f.config.DocType

	listFiles, err := client.ReadDir(rootPath)
	if err != nil {
		logrus.Errorf("Не возможно прочитать директорию - %s", err)
	}

	// массив директорий внутри которых нужен поиск
	var lister []os.FileInfo

	for _, value := range listFiles {
		if value.IsDir() == true {
			lister = append(lister, value)
		}
	}

	var fileinfo model.FileInfo
	var fileinfolist []model.FileInfo
	for _, value := range lister {
		Walk(client, rootPath+"/"+value.Name()+"/"+docType, func(fullpath string, info os.FileInfo, err error) error {
			if err != nil {
				if err.(goftp.Error).Code() == 550 {
					return nil
				}
				return err
			}
			// fmt.Println(fullpath)
			fileinfo.FileName = info.Name()
			fileinfo.FilePath = fullpath
			fileinfo.FileSize = info.Size()
			fileinfo.FileDateMod = info.ModTime()
			fileinfo.FileArea = value.Name()
			fileinfo.FileIsDir = info.IsDir()
			fileinfolist = append(fileinfolist, fileinfo)
			// massStr = append(massStr, fullpath)
			return nil
		}, from, to)
	}
	return fileinfolist
}

//Walk Рекурсивный перебор
func Walk(client *goftp.Client, root string, walkFn filepath.WalkFunc, from time.Time, to time.Time) (ret error) {
	dirsToCheck := make(chan string, 100)

	var workCount int32 = 1
	dirsToCheck <- root

	for dir := range dirsToCheck {
		go func(dir string) {
			files, err := client.ReadDir(dir)

			if err != nil {
				if err = walkFn(dir, nil, err); err != nil && err != filepath.SkipDir {
					ret = err
					close(dirsToCheck)
					return
				}
			}

			for _, file := range files {
				if file.ModTime().After(from) && file.ModTime().Before(to) && file.IsDir() == false {
					if err = walkFn(path.Join(dir, file.Name()), file, nil); err != nil {
						if file.IsDir() && err == filepath.SkipDir {
							continue
						}
						ret = err
						close(dirsToCheck)
						return
					}
				}

				if file.IsDir() {
					atomic.AddInt32(&workCount, 1)
					dirsToCheck <- path.Join(dir, file.Name())
				}
			}

			atomic.AddInt32(&workCount, -1)
			if workCount == 0 {
				close(dirsToCheck)
			}
		}(dir)
	}

	return ret
}

// ConfigureDb ...
func (f *FtpDownloader) ConfigureDb() error {
	st := db.New(f.config)
	if err := st.Open(); err != nil {
		return err
	}
	f.db = st
	return nil
}
