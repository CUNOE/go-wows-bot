package wows

import (
	"go-wows-bot/model"
	"strconv"
)

//func setHistoryData(info *model.RecentInfo) string {
//	historyData := ""
//	for _, r := range info.Data.RecentList {
//		winColor := setWinColor(r.Wins)
//		damagecolor := setDamageColor(r.ShipInfo.ShipType, r.Damage)
//		historyData += "<tr>"
//		historyData += "<td class=\"blueColor\">" + r.ShipInfo.NameCn + "</td>"
//		historyData += "<td class=\"blueColor\">" + strconv.Itoa(r.ShipInfo.Level) + "</td>"
//		historyData += "<td class=\"blueColor\">" + strconv.Itoa(r.Battles) + "</td>"
//		historyData += "<td class=\"blueColor\" style=\"color: " + r.Pr.Color + "\">" + strconv.Itoa(r.Pr.Value) + r.Pr.Name + "</td>"
//		historyData += "<td class=\"blueColor\">" + strconv.Itoa(r.Xp) + "</td>"
//		historyData += "<td class=\"blueColor\" style=\"color: " + winColor + "\">" + strconv.FormatFloat(r.Wins, 'f', 2, 64) + "%</td>"
//		historyData += "<td class=\"blueColor\" style=\"color: " + damagecolor + "\">" + strconv.Itoa(r.Damage) + "</td>"
//		historyData += "<td class=\"blueColor\">" + strconv.FormatFloat(r.Hit, 'f', 2, 64) + "%" + "</td>"
//		historyData += "</tr>"
//	}
//	log.Printf(historyData)
//	return historyData
//}

func setHistoryData(info *model.RecentInfo) string {
	historyData := ""
	for _, r := range info.Data.RecentList {
		winColor := setWinColor(r.Wins)
		damagecolor := setDamageColor(r.ShipInfo.ShipType, r.Damage)
		historyData += "<tr>"
		historyData += "<td class=\"blueColor\">" + r.ShipInfo.NameCn + "</td>"
		historyData += "<td class=\"blueColor\">" + strconv.Itoa(r.ShipInfo.Level) + "</td>"
		historyData += "<td class=\"blueColor\">" + strconv.Itoa(r.Battles) + "</td>"
		historyData += "<td class=\"blueColor\" style=\"color: " + r.Pr.Color + "\">" + strconv.Itoa(r.Pr.Value) + r.Pr.Name + "</td>"
		historyData += "<td class=\"blueColor\">" + strconv.Itoa(r.Xp) + "</td>"
		historyData += "<td class=\"blueColor\" style=\"color: " + winColor + "\">" + strconv.FormatFloat(r.Wins, 'f', 2, 64) + "%</td>"
		historyData += "<td class=\"blueColor\" style=\"color: " + damagecolor + "\">" + strconv.Itoa(r.Damage) + "</td>"
		historyData += "<td class=\"blueColor\">" + strconv.FormatFloat(r.Hit, 'f', 2, 64) + "%" + "</td>"
		historyData += "</tr>"
	}
	return historyData
}

func setDamageColor(shipType string, value int) (color string) {
	var color_data map[string]string
	color_data = make(map[string]string)

	color_data["Bad"] = "#FE0E00"
	color_data["Good"] = "#44B300"
	color_data["Great"] = "#02C9B3"
	color_data["Unicum"] = "#D042F3"
	color_data["Super Unicum"] = "#A00DC5"

	switch {
	case shipType == "Destroyer":
		switch {
		case value < 33000:
			return color_data["Bad"]
		case value < 40000:
			return color_data["Good"]
		case value < 55000:
			return color_data["Great"]
		case value < 64000:
			return color_data["Unicum"]
		default:
			return color_data["Super Unicum"]
		}
	case shipType == "Cruiser":
		switch {
		case value < 47000:
			return color_data["Bad"]
		case value < 55000:
			return color_data["Good"]
		case value < 83000:
			return color_data["Great"]
		case value < 95000:
			return color_data["Unicum"]
		default:
			return color_data["Super Unicum"]
		}
	case shipType == "AirCarrier":
		switch {
		case value < 60000:
			return color_data["Bad"]
		case value < 71000:
			return color_data["Good"]
		case value < 84000:
			return color_data["Great"]
		case value < 113000:
			return color_data["Unicum"]
		default:
			return color_data["Super Unicum"]
		}
	default:
		switch {
		case value < 64000:
			return color_data["Bad"]
		case value < 72000:
			return color_data["Good"]
		case value < 97000:
			return color_data["Great"]
		case value < 108000:
			return color_data["Unicum"]
		default:
			return color_data["Super Unicum"]
		}
	}
	return
}

func setWinColor(v float64) (color string) {
	value := int(v)

	var color_data map[string]string
	color_data = make(map[string]string)

	color_data["Bad"] = "#FE0E00"
	color_data["Below Average"] = "#FE7903"
	color_data["Average"] = "#FFC71F"
	color_data["Good"] = "#44B300"
	color_data["Very Good"] = "#318000"
	color_data["Great"] = "#02C9B3"
	color_data["Unicum"] = "#D042F3"
	color_data["Super Unicum"] = "#A00DC5"

	switch {
	case value < 45:
		return color_data["Bad"]
	case value < 50:
		return color_data["Below Average"]
	case value < 55:
		return color_data["Average"]
	case value < 60:
		return color_data["Good"]
	case value < 65:
		return color_data["Great"]
	case value < 70:
		return color_data["Unicum"]
	default:
		return color_data["Super Unicum"]

	}
	return
}

func setUpinfoColor(value int) (color string) {
	var color_data map[string]string
	color_data = make(map[string]string)

	color_data["Bad"] = "#FE0E00"
	color_data["Good"] = "#44B300"

	switch {
	case value < 0:
		return color_data["Bad"]
	case value > 0:
		return color_data["Good"]
	default:
		return

	}
}
