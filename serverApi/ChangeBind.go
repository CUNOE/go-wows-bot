package serverApi

import (
	"github.com/go-resty/resty/v2"
	"go-wows-bot/global"
	"go-wows-bot/model"
	"log"
)

func ChangeBind(server string, serverId string, accountId string) bool {
	c := resty.New()
	r := &model.ChangeBind{}

	hearders := map[string]string{
		"accept":        "application/json",
		"Authorization": global.Token,
	}

	log.Printf("切换绑定信息请求：%v", global.API_Server+"/api/wows/bind/account/platform/bind/put?platformType="+server+"&platformId="+serverId+"&accountId="+accountId)
	c.R().SetResult(r).SetHeaders(hearders).Get(global.API_Server + "/api/wows/bind/account/platform/bind/put?platformType=" + server + "&platformId=" + serverId + "&accountId=" + accountId)

	if r.Message != "success" {
		log.Printf(r.Message)
		return false
	}

	return true
}
