package common

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Sterks/XmlReader/internal/app/configuration"
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

// GetIDLastDB ...
func GetIDLastDB(db *sql.DB, config *configuration.Configuration) int {
	var number int
	_ = db.QueryRow(`select currval('public."Files_f_id_seq"')`).Scan(&number)
	return number
}

//GenerateID ...
func GenerateID(ident int) string {

	word := strconv.Itoa(ident)
	ch := len(word)
	nool := 12 - ch
	var ap string
	ap = word
	for i := 0; i < nool; i++ {
		ap = fmt.Sprintf("0%s", ap)
	}
	return ap
}

// CreateFolder ...
func CreateFolder(config *configuration.Configuration, ident int) string {
	saveDir := config.FileDir
	if err := os.MkdirAll(saveDir, 0755); err != nil {
		log.Fatal(err)
	}
	stringID := GenerateID(ident)
	lv1 := fmt.Sprint(stringID[0:3])
	lv2 := fmt.Sprint(stringID[3:6])
	lv3 := fmt.Sprint(stringID[6:9])
	// lv4 := fmt.Sprintln(stringID[9:12])
	if err := os.MkdirAll(saveDir+"/"+lv1, 0755); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(saveDir+"/"+lv1+"/"+lv2, 0755); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(saveDir+"/"+lv1+"/"+lv2+"/"+lv3, 0755); err != nil {
		log.Fatal(err)
	}
	// if err := os.MkdirAll(lv4, 0755); err != nil {
	// 	log.Fatal(err)
	// }
	path := fmt.Sprintf("%s/%s/%s/", lv1, lv2, lv3)
	return path
}

// GetLocalPath ...
func GetLocalPath(config *configuration.Configuration, ident int) string {
	rootPath := config.FileDir
	word := strconv.Itoa(ident)
	stringID := GenerateID(ident)
	lv1 := fmt.Sprint(stringID[0:3])
	lv2 := fmt.Sprint(stringID[3:6])
	lv3 := fmt.Sprint(stringID[6:9])
	s := fmt.Sprint(rootPath + "/" + lv1 + "/" + lv2 + "/" + lv3 + "/" + word)
	return s
}

// Hash ...
func Hash(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	h := md5.New()
	if _, err := io.Copy(h, file); err != nil {
		log.Fatal(err)
	}
	hash := hex.EncodeToString(h.Sum(nil))

	return hash
}
