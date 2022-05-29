package wows

import (
	"go-wows-bot/serverApi"
	"go-wows-bot/util"
	"log"
	"strconv"
)

// wows server
func wwsServer(group_id int64) {
	log.Printf("向群：%v发送wws服务器信息", group_id)
	data, _ := serverApi.GetServerList()

	message := ""
	for _, d := range data.Data {
		message += d.Value + "  " + d.Key + "\n"
	}

	WriteMessage(message, group_id)
	return
}

// wows nation
func wwsNation(group_id int64) {
	log.Printf("向群：%v发送wws国家信息", group_id)
	data, _ := serverApi.GetNationList()

	message := ""
	for _, d := range data.Data {
		message += d.Cn + "  " + d.Nation + "\n"
	}

	WriteMessage(message, group_id)
	return
}

// wows ship.type
func wwsShipType(group_id int64) {
	log.Printf("向群：%v发送wws舰船类型信息", group_id)
	data, _ := serverApi.GetShipType()

	message := ""
	message += data.Data.AIRCARRIER.Key + "  " + data.Data.AIRCARRIER.Value + "\n"
	message += data.Data.AUXILIARY.Key + "  " + data.Data.AUXILIARY.Value + "\n"
	message += data.Data.BATTLESHIP.Key + "  " + data.Data.BATTLESHIP.Value + "\n"
	message += data.Data.CRUISER.Key + "  " + data.Data.CRUISER.Value + "\n"
	message += data.Data.DESTROYER.Key + "  " + data.Data.DESTROYER.Value + "\n"
	message += data.Data.SUBMARINE.Key + "  " + data.Data.SUBMARINE.Value + "\n"

	WriteMessage(message, group_id)
	return
}

func wwsHelp(group_id int64) {
	log.Printf("向群：%v发送帮助信息", group_id)

	message := "帮助列表\n    " +
		"wws help：获取帮助\n    " +
		"wws bind/set/绑定 服务器 游戏昵称：绑定QQ与游戏账号\n    " +
		"wws 查询绑定/查绑定/绑定列表[me]：查询指定用户的绑定账号\n    " +
		"wws 切换绑定[id]：使用查绑定中的序号快速切换绑定账号\n    " +
		"wws [服务器+游戏昵称][me]：查询账号总体战绩\n    " +
		"wws [服务器+游戏昵称][me] recent [日期]：查询账号近期战绩，默认1天\n    " +
		"wws [server][nation][ship.type]：查询战舰相关信息\n    " +
		"[待开发] wws ship recent\n    " +
		"[待开发] wws rank\n    " +
		"猴王Love❥(^_-)"

	WriteMessage(message, group_id)
	return
}

func searchBindId(message string, group_id int64, qq_id string) {
	log.Printf("QQ：%v在群：%v查询了自己的绑定信息", qq_id, group_id)
	result, err := serverApi.GetBindInfo(qq_id)
	util.ErrorHandle(err)

	msg := ""
	for num, r := range result.Data {
		msg += strconv.Itoa(num) + ":   " + r.UserName + "\n"

	}

	WriteMessage(msg, group_id)
	return
}
