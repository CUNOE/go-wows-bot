package global

import (
	"go-wows-bot/util"
	"os"
	"path/filepath"
)

func InitAll() {
	LoadConfig()
	delData()
}

func delData() {
	err := os.RemoveAll(filepath.Join(CurrentPath, "/data/"))
	util.ErrorHandle(err)
	err = os.Mkdir(filepath.Join(CurrentPath, "/data/"), os.ModePerm)
	util.ErrorHandle(err)
}
