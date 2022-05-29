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

func isShipInfo(message string, group_id int64, qq_id string) {
	var Select = 0
	log.Printf("QQ：%v在群：%v查询了舰船信息", qq_id, group_id)
	message = regexp.MustCompile("^wws ").ReplaceAllString(message, "")
	message = regexp.MustCompile(" ship").ReplaceAllString(message, "")

	if regexp.MustCompile(" [0-9]$").FindAllStringSubmatch(message, -1) != nil {
		Id, err := strconv.Atoi(strings.Replace(regexp.MustCompile(" [0-9]$").FindAllStringSubmatch(message, -1)[0][0], " ", "", -1))
		if err != nil {
			WriteMessage("参数有误", group_id)
			return
		}
		Select = Id
	}
	log.Printf(strconv.Itoa(Select))
	log.Printf(message)
	message = regexp.MustCompile(" [0-9]$").ReplaceAllString(message, "")
	log.Printf(message)

	if regexp.MustCompile("^ me").FindAllStringSubmatch(message, -1) != nil {
		message = regexp.MustCompile("^ me").ReplaceAllString(message, "")
		wwsShipInfoByQQ(group_id, qq_id, strings.Replace(message, " ", "", -1), Select)
		return
	}

	server := ""
	switch {
	case regexp.MustCompile("^asia |亚服 |asian |亚洲 |Asia ").FindAllStringSubmatch(message, -1) != nil:
		server = "asia"
		message = regexp.MustCompile("^asia |亚服 |asian |亚洲 |Asia ").ReplaceAllString(message, "")
	case regexp.MustCompile("^eu |欧服 |europe |欧洲 ").FindAllStringSubmatch(message, -1) != nil:
		server = "eu"
		message = regexp.MustCompile("^eu|欧服|europe|欧洲").ReplaceAllString(message, "")
	case regexp.MustCompile("^na |美服 |America |NorthAmerican |美洲 |南美 |美国 ").FindAllStringSubmatch(message, -1) != nil:
		server = "na"
		message = regexp.MustCompile("^na |美服 |America |NorthAmerican |美洲 |南美 |美国 ").ReplaceAllString(message, "")
	case regexp.MustCompile("^ru |俄服 |russia |Russia |毛子 |俄罗斯 |毛服 ").FindAllStringSubmatch(message, -1) != nil:
		server = "ru"
		message = regexp.MustCompile("^ru |俄服 |russia |Russia |毛子 |俄罗斯 |毛服 ").ReplaceAllString(message, "")
	case regexp.MustCompile("^cn |国服 |china |China ").FindAllStringSubmatch(message, -1) != nil:
		//server = "cn"
		//message = regexp.MustCompile("^.cn |国服 |china |China ").ReplaceAllString(message, "")
		WriteMessage("暂不支持国服查询", group_id)
	default:
		WriteMessage("请求参数有误 请检查您的参数", group_id)
	}
	log.Printf(message)
	usernameGet := regexp.MustCompile("^.* ").FindAllStringSubmatch(message, -1)
	username := usernameGet[0][0]
	log.Printf(username)
	log.Printf(message)
	message = regexp.MustCompile("^.* ").ReplaceAllString(message, "")
	log.Printf(message)
	wwsShipInfoByName(server, strings.Replace(username, " ", "", -1), group_id, strings.Replace(message, " ", "", -1), Select)
	return
}

//func GetShipIdByID(shipName string, groupId int64, id int) (shipId string) {
//	r, _ := serverApi.ShipSearch(shipName)
//	if len(r.Data) == 0 {
//		WriteMessage(r.Message, groupId)
//		return
//	}
//	if len(r.Data) > 1 {
//		shipId = strconv.FormatInt(r.Data[id].Id, 10)
//	}
//
//	return
//}

func GetShipId(shipName string, groupId int64, id int) (shipId string) {
	r, _ := serverApi.ShipSearch(shipName)

	if len(r.Data) == 0 {
		WriteMessage("参数有误 请检查", groupId)
		return
	}

	if len(r.Data) == 1 {
		shipId = strconv.FormatInt(r.Data[0].Id, 10)
		return
	}

	if len(r.Data) > 1 {
		msg := "该名称还有如下舰艇，在同样的语句之后加上空格跟数字即可选择\n"
		for k, x := range r.Data {
			msg += strconv.Itoa(k) + ":  " + x.ShipNameCn + "\n"
		}
		WriteMessage(msg, groupId)
	}
	shipId = strconv.FormatInt(r.Data[id].Id, 10)

	return
}

func wwsShipInfoByQQ(groupId int64, qqId string, shipName string, id int) {
	shipId := GetShipId(shipName, groupId, id)
	if shipId == "" {
		return
	}
	image, err := wwsShipInfo("QQ", qqId, shipId)
	util.ErrorHandle(err)

	cqmsg := "[CQ:image,file=file:///" + image + ",id=40000,subType=0]"
	WriteMessage(cqmsg, groupId)
}

func wwsShipInfoByName(server string, userName string, groupId int64, shipName string, id int) {
	accountId := serverApi.GetAccountId(server, userName)
	if accountId == "Error" {
		return
	}
	shipId := GetShipId(shipName, groupId, id)
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

	createPath := filepath.Join(global.CurrentPath, "/temp/", r.Data.UserName) + strconv.FormatInt(time.Now().Unix(), 10) + r.Data.ShipInfo.ShipInfo.NameCn + ".html"
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

	imagePath = ImageRender("file:///" + createPath)

	log.Printf(createPath)
	log.Printf(filepath.Join(global.CurrentPath, "/template/go-wws-ship.html"))
	log.Printf(imagePath)
	return
}
