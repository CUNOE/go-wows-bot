package wows

import (
	"errors"
	"fmt"
	"go-wows-bot/global"
	"go-wows-bot/serverApi"
	"go-wows-bot/util"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var flag = false

var queue = make(chan interface{})

func isShipInfo(message string, group_id int64, qq_id string) {
	log.Printf("QQ：%v在群：%v查询了舰船信息", qq_id, group_id)
	message = regexp.MustCompile("^wws").ReplaceAllString(message, "")
	message = regexp.MustCompile(" ship").ReplaceAllString(message, "")

	if regexp.MustCompile("^ me").FindAllStringSubmatch(message, -1) != nil {
		message = regexp.MustCompile("^ me").ReplaceAllString(message, "")
		wwsShipInfoByQQ(group_id, qq_id, strings.Replace(message, " ", "", -1))
		return
	}

	server := ""
	switch {
	case regexp.MustCompile("^.asia |亚服 |asian |亚洲 ").FindAllStringSubmatch(message, -1) != nil:
		server = "asia"
		message = regexp.MustCompile("^.asia|亚服|asian|亚洲").ReplaceAllString(message, "")
	case regexp.MustCompile("^.eu |欧服 |europe |欧洲 ").FindAllStringSubmatch(message, -1) != nil:
		server = "eu"
		message = regexp.MustCompile("^.eu|欧服|europe|欧洲").ReplaceAllString(message, "")
	case regexp.MustCompile("^.na |美服 |America |NorthAmerican |美洲 |南美 |美国 ").FindAllStringSubmatch(message, -1) != nil:
		server = "na"
		message = regexp.MustCompile("^.na|美服|America|NorthAmerican|美洲|南美|美国").ReplaceAllString(message, "")
	case regexp.MustCompile("^.ru|俄服|russia|Russia|毛子|俄罗斯|毛服").FindAllStringSubmatch(message, -1) != nil:
		server = "ru"
		message = regexp.MustCompile("^.ru |俄服 |russia |Russia |毛子 |俄罗斯 |毛服 ").ReplaceAllString(message, "")
	case regexp.MustCompile("^.cn |国服 |china |China ").FindAllStringSubmatch(message, -1) != nil:
		//server = "cn"
		//message = regexp.MustCompile("^.cn |国服 |china |China ").ReplaceAllString(message, "")
		WriteMessage("暂不支持国服查询", group_id)
	default:
		WriteMessage("请求参数有误 请检查您的参数", group_id)
	}
	usernameGet := regexp.MustCompile(" .* ").FindAllStringSubmatch(message, -1)
	username := usernameGet[0][0]

	message = regexp.MustCompile(" .* ").ReplaceAllString(message, "")

	wwsShipInfoByName(server, strings.Replace(username, " ", "", -1), group_id, strings.Replace(message, " ", "", -1))
	return
}

func GetShipId(shipName string, groupId int64) (shipId string) {
	r, _ := serverApi.ShipSearch(shipName)
	if len(r.Data) == 0 {
		WriteMessage("请求参数有误", groupId)
		return
	}
	if len(r.Data) > 1 {
		msg := ""
		for k, x := range r.Data {
			msg += strconv.Itoa(k) + ":  " + x.ShipNameCn
		}
		msg += "请在10s内选择"
		WriteMessage("舰船选择功能W未完善， 将默认选择No.0", groupId)
	}
	shipId = strconv.FormatInt(r.Data[0].Id, 10)

	return
}

func wwsShipInfoByQQ(groupId int64, qqId string, shipName string) {
	shipId := GetShipId(shipName, groupId)
	if shipId == "" {
		return
	}
	image, err := wwsShipInfo("QQ", qqId, shipId)
	util.ErrorHandle(err)

	cqmsg := "[CQ:image,file=file:///" + image + ",id=40000,subType=0]"
	WriteMessage(cqmsg, groupId)
}

func wwsShipInfoByName(server string, userName string, groupId int64, shipName string) {
	accountId := serverApi.GetAccountId(server, userName)
	if accountId == "Error" {
		return
	}
	shipId := GetShipId(shipName, groupId)
	if shipId == "" {
		return
	}

	image, err := wwsShipInfo(server, accountId, shipId)
	util.ErrorHandle(err)
	//if msg != "" {
	//	WriteMessage(msg, groupId)
	//	return
	//}

	cqmsg := "[CQ:image,file=file:///" + image + ",id=40000,subType=0]"
	WriteMessage(cqmsg, groupId)
}

func wwsShipInfo(server string, accountId string, shipId string) (imagePath string, err error) {
	r, err := serverApi.GetShipInfo(server, accountId, shipId)
	if err != nil {
		return "", errors.New("Error")
	}

	var htmlData map[string]interface{}
	htmlData = make(map[string]interface{})

	htmlData["clanColor"] = r.Data.ClanInfo.ColorRgb
	htmlData["guild"] = r.Data.ClanInfo.Tag
	htmlData["userName"] = r.Data.UserName
	htmlData["serverName"] = r.Data.ServerName
	htmlData["shipNameCn"] = r.Data.ShipInfo.ShipInfo.NameCn
	htmlData["shipNameEn"] = r.Data.ShipInfo.ShipInfo.NameEnglish
	htmlData["damageTopColor"] = setUpinfoColor(r.Data.DwpDataVO.Damage)
	htmlData["damageTop"] = r.Data.DwpDataVO.Damage
	htmlData["winsTopColor"] = setUpinfoColor(int(r.Data.DwpDataVO.Wins))
	htmlData["winsTop"] = r.Data.DwpDataVO.Wins
	htmlData["prTopColor"] = setUpinfoColor(r.Data.DwpDataVO.Pr)
	htmlData["prTop"] = r.Data.DwpDataVO.Pr
	htmlData["prValueColor"] = r.Data.ShipInfo.Pr.Color
	htmlData["prValue"] = r.Data.ShipInfo.Pr.Value
	htmlData["lastTime"] = time.Unix(int64(r.Data.ShipInfo.LastBattlesTime), 0).Format("2006-01-02 15:04:05")
	htmlData["battles"] = r.Data.ShipInfo.Battles
	htmlData["wins"] = r.Data.ShipInfo.Wins
	htmlData["damage"] = r.Data.ShipInfo.Damage
	htmlData["xp"] = r.Data.ShipInfo.Xp
	htmlData["kda"] = r.Data.ShipInfo.Kd
	htmlData["hit"] = r.Data.ShipInfo.Hit
	htmlData["maxDamage"] = r.Data.ShipInfo.ExtensionDataInfo.MaxDamage
	htmlData["maxDamageScouting"] = r.Data.ShipInfo.ExtensionDataInfo.MaxDamageScouting
	htmlData["maxTotalAgro"] = r.Data.ShipInfo.ExtensionDataInfo.MaxTotalAgro
	htmlData["maxXp"] = r.Data.ShipInfo.ExtensionDataInfo.MaxXp
	htmlData["maxFragsBattle"] = r.Data.ShipInfo.ExtensionDataInfo.MaxFrags
	htmlData["maxPlanesKilled"] = r.Data.ShipInfo.ExtensionDataInfo.MaxPlanesKilled

	createPath := filepath.Join(global.CurrentPath, "/data/") + "\\" + r.Data.UserName + strconv.FormatInt(time.Now().Unix(), 10) + r.Data.ShipInfo.ShipInfo.NameCn + ".html"
	create, err := os.Create(createPath)
	if err != nil {
		return
	}

	t, err := template.ParseFiles(filepath.Join(global.CurrentPath, "/template/go-wws-ship.html"))

	if err != nil {
		fmt.Println("create error")
	}

	err = t.Execute(create, htmlData)
	if err != nil {
		fmt.Println("execute error")
	}

	imagePath = ImageRender(createPath)
	return
}
