package main

import (
	"go-wows-bot/global"
	"go-wows-bot/util"
	"go-wows-bot/wows"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	global.CurrentPath = getCurrentAbPathByExecutable()
	global.InitAll()

	generServer(global.ServerStartPort)
	//times := ""
	//i, _ := strconv.Atoi(times)
	//if i == 0 {
	//	fmt.Printf("yes")
	//}
	//log.Printf(strconv.FormatInt(time.Now().Unix(), 10))
}

func generServer(port string) {

	r := gin.Default()
	r.GET("/ws", wows.ReadMessage)
	err := r.Run(":" + port)
	util.ErrorHandle(err)
	log.Printf("listen in %v", port)
}

func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// TODO 内部启动go-cqhttp
func start_go_cqhttp() {

}
