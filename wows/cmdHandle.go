package wows

import (
	"regexp"
)

func Handle(message string, group_id int64, sender_id string) {
	switch {
	case message == "wws help":
		wwsHelp(group_id)
	case message == "wws server":
		wwsServer(group_id)
	case message == "wws nation":
		wwsNation(group_id)
	case message == "wws ship.type":
		wwsShipType(group_id)
	case isSearchBind(message):
		searchBindId(message, group_id, sender_id)
	case isBind(message):
		bindUser(message, group_id, sender_id)
	case isChangeBind(message):
		changeBind(message, group_id, sender_id)
	case isWws(message):
		switch {
		case isRecent(message):
			isRecentInfo(message, group_id, sender_id)
		case isShip(message):
			isShipInfo(message, group_id, sender_id)
		default:
			isInfo(message, group_id, sender_id)
		}

	default:
		return
	}
}

func isChangeBind(message string) bool {
	r := regexp.MustCompile("^wws 切换绑定").FindAllStringSubmatch(message, -1)
	if r == nil {
		return false
	}
	return true
}

func isRecent(message string) bool {
	r := regexp.MustCompile(".* recent.*").FindAllStringSubmatch(message, -1)
	if r == nil {
		return false
	}
	return true
}

func isShip(message string) bool {
	r := regexp.MustCompile(".* ship .*").FindAllStringSubmatch(message, -1)
	if r == nil {
		return false
	}
	return true
}

func isWws(message string) bool {
	r := regexp.MustCompile("^wws").FindAllStringSubmatch(message, -1)
	if r == nil {
		return false
	}
	return true
}

func isSearchBind(message string) bool {
	reg := regexp.MustCompile("^wws 查询绑定|wws 查绑定|wws 绑定列表")
	result := reg.FindAllStringSubmatch(message, -1)
	if result == nil {
		return false
	}
	return true
}

func isBind(message string) bool {
	reg1 := regexp.MustCompile("^wws 绑定|wws set|wws bind")
	result1 := reg1.FindAllStringSubmatch(message, -1)
	if result1 == nil {
		return false
	}
	return true
}
