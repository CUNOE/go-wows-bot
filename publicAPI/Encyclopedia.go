package publicAPI

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"go-wows-bot/global"
	"go-wows-bot/model"
)

// 地区列表/public/wows/encyclopedia/nation/list
func GetNationList() (r *model.NationList, err error) {
	c := resty.New()
	r = &model.NationList{}

	c.R().SetResult(r).Get(global.API_Server + "/public/wows/encyclopedia/nation/list")

	if r.Message != "success" {

		return
	}

	return
}

func GetServerList() (r *model.ServerList, err error) {
	c := resty.New()
	r = &model.ServerList{}

	c.R().SetResult(r).Get(global.API_Server + "/public/wows/encyclopedia/server/list")

	if r.Message != "success" {
		err = errors.New("网络错误")
		return
	}

	return
}

func GetShipType() (r *model.ShipType, err error) {
	c := resty.New()
	r = &model.ShipType{}

	c.R().SetResult(r).Get(global.API_Server + "/public/wows/encyclopedia/ship/type")

	if r.Message != "success" {
		err = errors.New("网络错误")
		return
	}

	return
}
