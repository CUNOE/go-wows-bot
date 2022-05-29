package serverApi

import (
	"github.com/go-resty/resty/v2"
	"go-wows-bot/global"
	"go-wows-bot/model"
	"log"
)

func GetBindInfo(qq_id string) (r *model.BindInfo, err error) {
	c := resty.New()
	r = &model.BindInfo{}

	log.Printf("获取绑定信息：%v", global.API_Server+"/public/wows/bind/account/platform/bind/list?platformType=QQ&platformId="+qq_id)
	c.R().SetResult(r).Get(global.API_Server + "/public/wows/bind/account/platform/bind/list?platformType=QQ&platformId=" + qq_id)

	if r.Message != "success" {
		log.Printf(r.Message)
		return
	}

	return
}

func BindUserApi(server string, username string, qq_id string) bool {
	c := resty.New()
	r := &model.BindBack{}

	hearders := map[string]string{
		"accept":        "application/json",
		"Authorization": global.Token,
	}

	log.Printf("绑定用户请求：%v", global.API_Server+"/api/wows/bind/account/platform/bind/put?platformType=QQ&platformId="+qq_id+"&accountId="+GetAccountId(server, username))
	c.R().SetResult(r).SetHeaders(hearders).Get(global.API_Server + "/api/wows/bind/account/platform/bind/put?platformType=QQ&platformId=" + qq_id + "&accountId=" + GetAccountId(server, username))

	if r.Message != "success" {
		log.Printf(r.Message)
		return false
	}

	return true
}
