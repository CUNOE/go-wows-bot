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
	err := os.RemoveAll(filepath.Join(CurrentPath, "/temp/"))
	util.ErrorHandle(err)
	err = os.Mkdir(filepath.Join(CurrentPath, "/temp/"), os.ModePerm)
	util.ErrorHandle(err)
}
