package serverApi

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"go-wows-bot/global"
	"go-wows-bot/model"
	"log"
	"strconv"
)

func GetAccountId(server string, userName string) (account_id string) {
	c := resty.New()
	r := &model.SearchUserInfo{}
	log.Printf("查询用户accountId请求：%v", global.API_Server+"/public/wows/account/search/user?server="+server+"&userName="+userName)
	c.R().SetResult(r).Get(global.API_Server + "/public/wows/account/search/user?server=" + server + "&userName=" + userName)

	if r.Message != "success" {
		return "Error"
	}

	return strconv.Itoa(r.Data.AccountId)
}

func GetUserInfo(server string, account_id string) (r *model.UserInfo, err error) {
	c := resty.New()
	r = &model.UserInfo{}

	//account_id := GetAccountId(server, userName)
	log.Printf("获取用户信息请求：%v", global.API_Server+"/public/wows/account/user/info?server="+server+"&accountId="+account_id)
	c.R().SetResult(r).Get(global.API_Server + "/public/wows/account/user/info?server=" + server + "&accountId=" + account_id)

	if r.Message != "success" {
		log.Printf(r.Message)
		return r, errors.New("GerUserInfo Error")
	}
	return
}

func GetShipInfo(server string, account_id string, shipId string) (r *model.ShipInfo, err error) {
	c := resty.New()
	r = &model.ShipInfo{}

	//account_id := GetAccountId(server, userName)
	log.Printf("获取舰船信息请求: %v", global.API_Server+"/public/wows/account/ship/info?server="+server+"&accountId="+account_id+"&shipId="+shipId)
	c.R().SetResult(r).Get(global.API_Server + "/public/wows/account/ship/info?server=" + server + "&accountId=" + account_id + "&shipId=" + shipId)

	if r.Message != "success" {
		log.Printf(r.Message)
		return
	}

	return
}
