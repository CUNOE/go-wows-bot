package serverApi

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"go-wows-bot/global"
	"go-wows-bot/model"
	"log"
	"net/url"
)

// 地区列表/public/wows/encyclopedia/nation/list
func GetNationList() (r *model.NationList, err error) {
	c := resty.New()
	r = &model.NationList{}

	log.Printf("获取地区列表：%v", global.API_Server+"/public/wows/encyclopedia/nation/list")
	c.R().SetResult(r).Get(global.API_Server + "/public/wows/encyclopedia/nation/list")

	if r.Message != "success" {
		log.Printf(r.Message)
		err = errors.New("网络错误")
		return
	}

	return
}

func GetServerList() (r *model.ServerList, err error) {
	c := resty.New()
	r = &model.ServerList{}
	log.Printf("获取服务器列表：%v", global.API_Server+"/public/wows/encyclopedia/server/list")
	c.R().SetResult(r).Get(global.API_Server + "/public/wows/encyclopedia/server/list")

	if r.Message != "success" {
		log.Printf(r.Message)
		err = errors.New("网络错误")
		return
	}

	return
}

func GetShipType() (r *model.ShipType, err error) {
	c := resty.New()
	r = &model.ShipType{}
	log.Printf("获取舰船类型：%v", global.API_Server+"/public/wows/encyclopedia/ship/type")
	c.R().SetResult(r).Get(global.API_Server + "/public/wows/encyclopedia/ship/type")

	if r.Message != "success" {
		log.Printf(r.Message)
		err = errors.New("网络错误")
		return
	}

	return
}

func ShipSearch(shipName string) (r *model.ShipSearch, err error) {
	c := resty.New()
	r = &model.ShipSearch{}
	log.Printf("获取舰船Id：%v", global.API_Server+"/public/wows/encyclopedia/ship/search?shipName="+url.QueryEscape(shipName))
	c.R().SetResult(r).Get(global.API_Server + "/public/wows/encyclopedia/ship/search?shipName=" + url.QueryEscape(shipName))

	if r.Message != "success" {
		log.Printf(r.Message)
		err = errors.New("网络错误")
		return
	}

	return
}
