package ftpdownloader

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"sync/atomic"

	"github.com/Sterks/XmlReader/internal/app/configuration"
	"github.com/secsy/goftp"

	"github.com/sirupsen/logrus"
)

// FtpDownloader ...
type FtpDownloader struct {
	config *configuration.Configuration
	logger *logrus.Logger
	ftp    *goftp.Client
}

//New ...
func New(con *configuration.Configuration) *FtpDownloader {
	return &FtpDownloader{
		config: con,
		logger: logrus.New(),
	}
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
func (f *FtpDownloader) GetFiles(client *goftp.Client) {
	rootPath := f.config.RootDir
	docType := f.config.DocType
	listFiles, err := client.ReadDir(rootPath)
	if err != nil {
		logrus.Errorf("Не возможно прочитать директорию - %s", err)
	}

	// маммив директорий внутри которых нужен поиск
	var lister []os.FileInfo

	for _, value := range listFiles {
		if value.IsDir() == true {
			lister = append(lister, value)
		}
	}

	for _, value := range lister {
		Walk(client, rootPath+"/"+value.Name()+"/"+docType, func(fullpath string, info os.FileInfo, err error) error {
			if err != nil {
				if err.(goftp.Error).Code() == 550 {
					return nil
				}
				return err
			}
			fmt.Println(fullpath)
			return nil
		})
	}

}

//Walk Рекурсивный перебор
func Walk(client *goftp.Client, root string, walkFn filepath.WalkFunc) (ret error) {
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
				if err = walkFn(path.Join(dir, file.Name()), file, nil); err != nil {
					if file.IsDir() && err == filepath.SkipDir {
						continue
					}
					ret = err
					close(dirsToCheck)
					return
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
