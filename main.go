package main

import (
	"github.com/gin-gonic/gin"
	"go-wows-bot/global"
	"go-wows-bot/util"
	"go-wows-bot/wows"
	"log"
	"path"
	"runtime"
)

func main() {
	global.CurrentPath = getCurrentAbPathByCaller()
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

func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

// TODO 内部启动go-cqhttp
func start_go_cqhttp() {

}
