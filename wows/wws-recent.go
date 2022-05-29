package wows

import (
	"errors"
	"fmt"
	"go-wows-bot/global"
	"go-wows-bot/serverApi"
	"go-wows-bot/util"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func isRecentInfo(message string, group_id int64, qq_id string) {
	log.Printf("QQ：%v在群：%v查询了recent记录", qq_id, group_id)
	var times string
	log.Println("recent")
	message = regexp.MustCompile("^wws").ReplaceAllString(message, "")

	message = regexp.MustCompile(" recent").ReplaceAllString(message, "")

	timesGet := regexp.MustCompile(" -?[1-9]\\d*$").FindAllStringSubmatch(message, -1)
	if timesGet == nil {
		times = "1"
	} else {
		times = timesGet[0][0]
	}
	times = strings.Replace(times, " ", "", -1)
	message = regexp.MustCompile(" -?[1-9]\\d*$").ReplaceAllString(message, "")

	server := ""
	switch {
	case regexp.MustCompile("^ me").FindAllStringSubmatch(message, -1) != nil:
		wwsRecentInfoByQQ(group_id, qq_id, strings.Replace(times, " ", "", -1))
	case regexp.MustCompile("^.asia |亚服 |asian |亚洲 ").FindAllStringSubmatch(message, -1) != nil:
		server = "asia"
		message = regexp.MustCompile("^.asia |亚服 |asian |亚洲 ").ReplaceAllString(message, "")
		wwsRecentInfoByName(server, strings.Replace(message, " ", "", -1), group_id, times)
	case regexp.MustCompile("^.eu |欧服 |europe |欧洲 ").FindAllStringSubmatch(message, -1) != nil:
		server = "eu"
		message = regexp.MustCompile("^.eu |欧服 |europe |欧洲 ").ReplaceAllString(message, "")
		wwsRecentInfoByName(server, strings.Replace(message, " ", "", -1), group_id, times)
	case regexp.MustCompile("^.na |美服 |America |NorthAmerican |美洲 |南美 |美国 ").FindAllStringSubmatch(message, -1) != nil:
		server = "na"
		message = regexp.MustCompile("^.na |美服 |America |NorthAmerican |美洲 |南美 |美国 ").ReplaceAllString(message, "")
		wwsRecentInfoByName(server, strings.Replace(message, " ", "", -1), group_id, times)
	case regexp.MustCompile("^.ru |俄服 |russia |Russia |毛子 |俄罗斯 |毛服 ").FindAllStringSubmatch(message, -1) != nil:
		server = "ru"
		message = regexp.MustCompile("^.ru |俄服 |russia |Russia |毛子 |俄罗斯 |毛服 ").ReplaceAllString(message, "")
		wwsRecentInfoByName(server, strings.Replace(message, " ", "", -1), group_id, times)
	case regexp.MustCompile("^.cn |国服 |china |China ").FindAllStringSubmatch(message, -1) != nil:
		//server = "cn"
		//message = regexp.MustCompile("^.cn |国服 |china |China ").ReplaceAllString(message, "")
		WriteMessage("暂不支持国服查询", group_id)
	default:
		WriteMessage("请求参数有误 请检查您的参数", group_id)
	}
}

func wwsRecentInfoByQQ(groupId int64, qqId string, times string) {
	msg, image, err := wwsRecentInfo("QQ", qqId, times)
	util.ErrorHandle(err)
	if msg != "" {
		WriteMessage(msg, groupId)
		return
	}

	cqmsg := "[CQ:image,file=file:///" + image + ",id=40000,subType=0]"
	WriteMessage(cqmsg, groupId)
}

func wwsRecentInfoByName(server string, userName string, groupId int64, times string) {
	accountId := serverApi.GetAccountId(server, userName)
	if accountId == "Error" {
		WriteMessage("查询用户错误", groupId)
		return
	}
	msg, image, err := wwsRecentInfo(server, accountId, times)
	util.ErrorHandle(err)
	if msg != "" {
		WriteMessage(msg, groupId)
		return
	}

	cqmsg := "[CQ:image,file=file:///" + image + ",id=40000,subType=0]"
	WriteMessage(cqmsg, groupId)
}

func wwsRecentInfo(server string, accountId string, times string) (msg string, imagePath string, err error) {
	atoi_time, err := strconv.Atoi(times)
	if err != nil {
		return "", "", errors.New("error")
	}
	seconds := strconv.Itoa(int(time.Now().Unix()) - (atoi_time * 1 * 3600 * 24))
	r, err := serverApi.GetRecentInfo(server, accountId, seconds)

	if r.Message != "success" {
		return r.Message, "", nil
	}

	var htmlData map[string]interface{}
	htmlData = make(map[string]interface{})

	htmlData["clanColor"] = r.Data.ClanInfo.ColorRgb
	htmlData["guild"] = r.Data.ClanInfo.Tag
	htmlData["userName"] = r.Data.UserName
	htmlData["serverName"] = r.Data.ServerName
	htmlData["prValueColor"] = r.Data.Data.Pr.Color
	htmlData["prValue"] = r.Data.Data.Pr.Value
	htmlData["reTime"] = time.Unix(int64(int(time.Now().Unix())-(atoi_time*1*3600*24)), 0).Format("2006-01-02 15:04:05")
	htmlData["battles"] = r.Data.Data.Battles
	htmlData["winsColor"] = setWinColor(r.Data.Data.Wins)
	htmlData["wins"] = r.Data.Data.Wins
	htmlData["damageColor"] = setDamageColor("", r.Data.Data.Damage)
	htmlData["damage"] = strconv.Itoa(r.Data.Data.Damage)
	htmlData["xp"] = r.Data.Data.Xp
	htmlData["kd"] = r.Data.Data.Kd
	htmlData["hit"] = r.Data.Data.Hit
	htmlData["historyData"] = setHistoryData(r)

	createPath := filepath.Join(global.CurrentPath, "/temp/", r.Data.UserName) + strconv.FormatInt(time.Now().Unix(), 10) + "recent.html"
	create, err := os.Create(createPath)
	if err != nil {
		return
	}

	t, err := template.ParseFiles(filepath.Join(global.CurrentPath, "/template/go-wws-info-recent.html"))

	if err != nil {
		fmt.Println("create error")
		return
	}

	err = t.Execute(create, htmlData)
	if err != nil {
		fmt.Println("execute error")
		return
	}

	imagePath = ImageRender("file:///" + createPath)

	log.Printf(createPath)
	log.Printf(filepath.Join(global.CurrentPath, "/template/go-wws-info-recent.html"))
	log.Printf(imagePath)
	return
}
