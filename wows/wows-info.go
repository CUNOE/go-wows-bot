package wows

import (
	"context"
	"fmt"
	"go-wows-bot/global"
	"go-wows-bot/serverApi"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func isInfo(message string, group_id int64, qq_id string) {
	log.Printf("QQ：%v在群：%v查询了舰队信息", qq_id, group_id)
	message = regexp.MustCompile("^wws").ReplaceAllString(message, "")
	server := ""
	switch {
	case regexp.MustCompile("^ me").FindAllStringSubmatch(message, -1) != nil:
		wwsInfoByQQ(group_id, qq_id)
	case regexp.MustCompile("^.asia |亚服 |asian |亚洲 ").FindAllStringSubmatch(message, -1) != nil:
		server = "asia"
		message = regexp.MustCompile("^.asia |亚服 |asian |亚洲 ").ReplaceAllString(message, "")
		wwsInfoByName(server, strings.Replace(message, " ", "", -1), group_id)
	case regexp.MustCompile("^.eu |欧服 |europe |欧洲 ").FindAllStringSubmatch(message, -1) != nil:
		server = "eu"
		message = regexp.MustCompile("^.eu |欧服 |europe |欧洲 ").ReplaceAllString(message, "")
		wwsInfoByName(server, strings.Replace(message, " ", "", -1), group_id)
	case regexp.MustCompile("^.na |美服 |America |NorthAmerican |美洲 |南美 |美国 ").FindAllStringSubmatch(message, -1) != nil:
		server = "na"
		message = regexp.MustCompile("^.na |美服 |America |NorthAmerican |美洲 |南美 |美国 ").ReplaceAllString(message, "")
		wwsInfoByName(server, strings.Replace(message, " ", "", -1), group_id)
	case regexp.MustCompile("^.ru |俄服 |russia |Russia |毛子 |俄罗斯 |毛服 ").FindAllStringSubmatch(message, -1) != nil:
		server = "ru"
		message = regexp.MustCompile("^.ru |俄服 |russia |Russia |毛子 |俄罗斯 |毛服 ").ReplaceAllString(message, "")

		wwsInfoByName(server, strings.Replace(message, " ", "", -1), group_id)
	case regexp.MustCompile("^.cn |国服 |china |China ").FindAllStringSubmatch(message, -1) != nil:
		//server = "cn"
		//message = regexp.MustCompile("^.cn |国服 |china |China ").ReplaceAllString(message, "")
		WriteMessage("暂不支持国服查询", group_id)
	default:
		WriteMessage("请求参数有误 请检查您的参数", group_id)
	}
}

func wwsInfoByQQ(groupId int64, qqId string) {
	image := wwsInfo("QQ", qqId, groupId)

	cqmsg := "[CQ:image,file=file:///" + image + ",id=40000,subType=0]"
	WriteMessage(cqmsg, groupId)
}

func wwsInfoByName(server string, userName string, groupId int64) {
	accountId := serverApi.GetAccountId(server, userName)
	if accountId == "Error" {
		WriteMessage("查询用户失败", groupId)
		return
	}
	image := wwsInfo(server, accountId, groupId)

	cqmsg := "[CQ:image,file=file:///" + image + ",id=40000,subType=0]"
	WriteMessage(cqmsg, groupId)
}

func wwsInfo(server string, accountId string, groupId int64) (imagePath string) {

	r, err := serverApi.GetUserInfo(server, accountId)
	if err != nil {
		WriteMessage(r.Message, groupId)
		return
	}

	var htmlData map[string]interface{}
	htmlData = make(map[string]interface{})

	htmlData["clanColor"] = r.Data.ClanInfo.ColorRgb
	htmlData["guild"] = r.Data.ClanInfo.Tag
	htmlData["userName"] = r.Data.UserName
	htmlData["karma"] = strconv.Itoa(r.Data.Karma)
	htmlData["serverName"] = r.Data.ServerName
	htmlData["newDamageColor"] = setUpinfoColor(r.Data.DwpDataVO.Damage)
	htmlData["newDamage"] = strconv.Itoa(r.Data.DwpDataVO.Damage)
	htmlData["newWinsColor"] = setUpinfoColor(int(r.Data.DwpDataVO.Wins))
	htmlData["newWins"] = r.Data.DwpDataVO.Wins
	htmlData["newPrColor"] = setUpinfoColor(r.Data.DwpDataVO.Pr)
	htmlData["newPr"] = strconv.Itoa(r.Data.DwpDataVO.Pr)
	htmlData["prValueColor"] = r.Data.Pr.Color
	htmlData["prValue"] = strconv.Itoa(r.Data.Pr.Value)
	htmlData["time"] = time.Unix(int64(r.Data.LastDateTime), 0).Format("2006-01-02 15:04:05")
	htmlData["battles"] = strconv.Itoa(r.Data.Pvp.Battles)
	htmlData["winsColor"] = setWinColor(r.Data.Pvp.Wins)
	htmlData["wins"] = r.Data.Pvp.Wins
	htmlData["damageColor"] = setDamageColor("", r.Data.Pvp.Damage)
	htmlData["damage"] = strconv.Itoa(r.Data.Pvp.Damage)
	htmlData["xp"] = r.Data.Pvp.Xp
	htmlData["kd"] = r.Data.Pvp.Kd
	htmlData["hit"] = r.Data.Pvp.Hit
	htmlData["bb_battles"] = strconv.Itoa(r.Data.Type.Battleship.Battles)
	htmlData["bb_prColor"] = r.Data.Type.Battleship.Pr.Color
	htmlData["bb_pr"] = strconv.Itoa(r.Data.Type.Battleship.Pr.Value)
	htmlData["bb_winsColor"] = setWinColor(r.Data.Type.Battleship.Wins)
	htmlData["bb_wins"] = r.Data.Type.Battleship.Wins
	htmlData["bb_damageColor"] = setDamageColor("", r.Data.Type.Battleship.Damage)
	htmlData["bb_damage"] = strconv.Itoa(r.Data.Type.Battleship.Damage)
	htmlData["bb_hit"] = r.Data.Type.Battleship.Hit
	htmlData["ca_battles"] = strconv.Itoa(r.Data.Type.Cruiser.Battles)
	htmlData["ca_prColor"] = r.Data.Type.Cruiser.Pr.Color
	htmlData["ca_pr"] = strconv.Itoa(r.Data.Type.Cruiser.Pr.Value)
	htmlData["ca_winsColor"] = setWinColor(r.Data.Type.Cruiser.Wins)
	htmlData["ca_wins"] = r.Data.Type.Cruiser.Wins
	htmlData["ca_damageColor"] = setDamageColor("", r.Data.Type.Cruiser.Damage)
	htmlData["ca_damage"] = strconv.Itoa(r.Data.Type.Cruiser.Damage)
	htmlData["ca_hit"] = r.Data.Type.Cruiser.Hit
	htmlData["dd_battles"] = strconv.Itoa(r.Data.Type.Destroyer.Battles)
	htmlData["dd_prColor"] = r.Data.Type.Destroyer.Pr.Color
	htmlData["dd_pr"] = strconv.Itoa(r.Data.Type.Destroyer.Pr.Value)
	htmlData["dd_winsColor"] = setWinColor(r.Data.Type.Destroyer.Wins)
	htmlData["dd_wins"] = r.Data.Type.Destroyer.Wins
	htmlData["dd_damageColor"] = setDamageColor("", r.Data.Type.Destroyer.Damage)
	htmlData["dd_damage"] = strconv.Itoa(r.Data.Type.Destroyer.Damage)
	htmlData["dd_hit"] = r.Data.Type.Destroyer.Hit
	htmlData["cv_battles"] = strconv.Itoa(r.Data.Type.AirCarrier.Battles)
	htmlData["cv_prColor"] = r.Data.Type.AirCarrier.Pr.Color
	htmlData["cv_pr"] = strconv.Itoa(r.Data.Type.AirCarrier.Pr.Value)
	htmlData["cv_winsColor"] = setWinColor(r.Data.Type.AirCarrier.Wins)
	htmlData["cv_wins"] = r.Data.Type.AirCarrier.Wins
	htmlData["cv_damageColor"] = setDamageColor("", r.Data.Type.AirCarrier.Damage)
	htmlData["cv_damage"] = strconv.Itoa(r.Data.Type.AirCarrier.Damage)
	htmlData["solo_battles"] = strconv.Itoa(r.Data.PvpSolo.Battles)
	htmlData["solo_winsColor"] = setWinColor(r.Data.PvpSolo.Wins)
	htmlData["solo_wins"] = r.Data.PvpSolo.Wins
	htmlData["solo_xp"] = r.Data.PvpSolo.Xp
	htmlData["solo_damageColor"] = setDamageColor("", r.Data.PvpSolo.Damage)
	htmlData["solo_damage"] = strconv.Itoa(r.Data.PvpSolo.Damage)
	htmlData["solo_kd"] = r.Data.PvpSolo.Kd
	htmlData["solo_hit"] = r.Data.PvpSolo.Hit
	htmlData["div2_battles"] = strconv.Itoa(r.Data.PvpTwo.Battles)
	htmlData["div2_winsColor"] = setWinColor(r.Data.PvpTwo.Wins)
	htmlData["div2_wins"] = r.Data.PvpTwo.Wins
	htmlData["div2_xp"] = r.Data.PvpTwo.Xp
	htmlData["div2_damageColor"] = setDamageColor("", r.Data.PvpTwo.Damage)
	htmlData["div2_damage"] = strconv.Itoa(r.Data.PvpTwo.Damage)
	htmlData["div2_kd"] = r.Data.PvpTwo.Kd
	htmlData["div2_hit"] = r.Data.PvpTwo.Hit
	htmlData["div3_battles"] = strconv.Itoa(r.Data.PvpThree.Battles)
	htmlData["div3_winsColor"] = setWinColor(r.Data.PvpThree.Wins)
	htmlData["div3_wins"] = r.Data.PvpThree.Wins
	htmlData["div3_xp"] = r.Data.PvpThree.Xp
	htmlData["div3_damageColor"] = setDamageColor("", r.Data.PvpThree.Damage)
	htmlData["div3_damage"] = strconv.Itoa(r.Data.PvpThree.Damage)
	htmlData["div3_kd"] = r.Data.PvpThree.Kd
	htmlData["div3_hit"] = r.Data.PvpThree.Hit
	// TODO:修复图表
	//htmlData["lv1"] = r.Data.BattleCountAll.Field1
	//htmlData["lv2"] = r.Data.BattleCountAll.Field2
	//htmlData["lv3"] = r.Data.BattleCountAll.Field3
	//htmlData["lv4"] = r.Data.BattleCountAll.Field4
	//htmlData["lv5"] = r.Data.BattleCountAll.Field5
	//htmlData["lv6"] = r.Data.BattleCountAll.Field6
	//htmlData["lv7"] = r.Data.BattleCountAll.Field7
	//htmlData["lv8"] = r.Data.BattleCountAll.Field8
	//htmlData["lv9"] = r.Data.BattleCountAll.Field9
	//htmlData["lv10"] = r.Data.BattleCountAll.Field10
	//htmlData["lv11"] = r.Data.BattleCountAll.Field11

	createPath := filepath.Join(global.CurrentPath, "/temp/", r.Data.UserName) + strconv.FormatInt(time.Now().Unix(), 10) + ".html"
	create, err := os.Create(createPath)
	if err != nil {
		return
	}

	t, err := template.ParseFiles(filepath.Join(global.CurrentPath, "/template/go-wws-info.html"))

	if err != nil {
		fmt.Println("create error")
	}

	err = t.Execute(create, htmlData)
	if err != nil {
		fmt.Println("execute error")
	}

	imagePath = ImageRender("file:///" + createPath)
	log.Printf(createPath)
	log.Printf(filepath.Join(global.CurrentPath, "/template/go-wws-info.html"))
	log.Printf(imagePath)
	return
}

func ImageRender(path string) string {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()
	log.Printf(path)

	var buf []byte

	if err := chromedp.Run(ctx, fullScreenshot(path, 90, &buf)); err != nil {
		log.Fatal(err)
	}
	path = regexp.MustCompile("^file:///").ReplaceAllString(path, "")
	if err := ioutil.WriteFile(path+".png", buf, 0o644); err != nil {
		log.Fatal(err)
	}
	return path + ".png"
}

func fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.FullScreenshot(res, quality),
	}
}
