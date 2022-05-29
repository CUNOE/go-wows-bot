package global

import (
	"github.com/gorilla/websocket"
)

var API_Server = "https://api.wows.linxun.link"

var Conn *websocket.Conn

var Token string

var ServerStartPort string
var ServerNumber int

var CurrentPath string
