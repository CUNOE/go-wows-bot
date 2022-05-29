package wows

import (
	"go-wows-bot/serverApi"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func bindUser(message string, group_id int64, qq_id string) {
	log.Printf("QQ：%v在群：%v绑定信息", qq_id, group_id)
	server, username := reBind(message)
	if serverApi.BindUserApi(server, username, qq_id) {
		WriteMessage("绑定成功", group_id)
		return
	}
	WriteMessage("绑定失败", group_id)
	return
}

func reBind(message string) (server string, username string) {
	reg1 := regexp.MustCompile("^wws 绑定|wws set|wws bind")
	result1 := reg1.FindAllStringSubmatch(message, -1)
	if result1 == nil {
		return "", ""
	}
	server = ""
	message = reg1.ReplaceAllString(message, "")
	switch {
	case regexp.MustCompile("^.asia|亚服|asian|亚洲").FindAllStringSubmatch(message, -1) != nil:
		server = "asia"
		message = regexp.MustCompile("^.asia|亚服|asian|亚洲").ReplaceAllString(message, "")
	case regexp.MustCompile("^.eu|欧服|europe|欧洲").FindAllStringSubmatch(message, -1) != nil:
		server = "eu"
		message = regexp.MustCompile("^.eu|欧服|europe|欧洲").ReplaceAllString(message, "")
	case regexp.MustCompile("^.na|美服|America|NorthAmerican|美洲|南美|美国").FindAllStringSubmatch(message, -1) != nil:
		server = "na"
		message = regexp.MustCompile("^.na|美服|America|NorthAmerican|美洲|南美|美国").ReplaceAllString(message, "")
	case regexp.MustCompile("^.ru|俄服|russia|Russia|毛子|俄罗斯|毛服").FindAllStringSubmatch(message, -1) != nil:
		server = "ru"
		message = regexp.MustCompile("^.ru|俄服|russia|Russia|毛子|俄罗斯|毛服").ReplaceAllString(message, "")
	case regexp.MustCompile("^.cn|国服|china|China").FindAllStringSubmatch(message, -1) != nil:
		server = "cn"
		message = regexp.MustCompile("^.cn|国服|china|China").ReplaceAllString(message, "")
	default:
		return "", ""
	}
	return server, strings.Replace(message, " ", "", -1)
}

func changeBind(message string, group_id int64, qq_id string) {
	log.Printf("QQ：%v在群：%v切换绑定信息", qq_id, group_id)
	message = regexp.MustCompile("^wws 切换绑定").ReplaceAllString(message, "")

	message = strings.Replace(message, " ", "", -1)
	log.Printf(message)
	i, err := strconv.Atoi(message)
	if err != nil {
		log.Printf("信息错误")
		WriteMessage("绑定错误", group_id)
		return
	}
	r, err := serverApi.GetBindInfo(qq_id)
	if err != nil {
		WriteMessage("绑定错误", group_id)
	}
	log.Printf("本次请求信息如下：绑定账号：%v 平台：%v 平台账户：%v， 绑定战舰世界账户id：%v", r.Data[i].UserName, r.Data[i].PlatformType, r.Data[i].PlatformId, strconv.Itoa(r.Data[i].AccountId))
	if serverApi.ChangeBind(r.Data[i].PlatformType, r.Data[i].PlatformId, strconv.Itoa(r.Data[i].AccountId)) {
		WriteMessage("切换绑定成功", group_id)
		return
	}
	WriteMessage("切换绑定失败", group_id)
	return

}
