package serverApi

import (
	"github.com/go-resty/resty/v2"
	"go-wows-bot/global"
	"go-wows-bot/model"
	"log"
)

func GetRecentInfo(server string, accountId string, seconds string) (r *model.RecentInfo, err error) {
	c := resty.New()
	r = &model.RecentInfo{}

	hearders := map[string]string{
		"accept":        "application/json",
		"Authorization": global.Token,
	}
	log.Printf("获取用户recent信息：%v", global.API_Server+"/api/wows/recent/recent/info?server="+server+"&accountId="+accountId+"&seconds="+seconds)
	c.R().SetResult(r).SetHeaders(hearders).Get(global.API_Server + "/api/wows/recent/recent/info?server=" + server + "&accountId=" + accountId + "&seconds=" + seconds)

	if r.Message != "success" {
		log.Printf(r.Message)
		return
	}

	return
}
