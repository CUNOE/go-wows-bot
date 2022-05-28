package wows

import (
	"go-wows-bot/publicAPI"
)

// wows server
func wwsServer(group_id int64) {
	data, _ := publicAPI.GetServerList()

	message := ""
	for _, d := range data.Data {
		message += d.Value + "  " + d.Key + "\n"
	}

	WriteMessage(message, group_id)
}

// wows nation
func wwsNation(group_id int64) {
	data, _ := publicAPI.GetNationList()

	message := ""
	for _, d := range data.Data {
		message += d.Cn + "  " + d.Nation + "\n"
	}

	WriteMessage(message, group_id)
}

// wows ship.type
func wwsShipType(group_id int64) {
	data, _ := publicAPI.GetShipType()

	message := ""
	message += data.Data.AIRCARRIER.Key + "  " + data.Data.AIRCARRIER.Value + "\n"
	message += data.Data.AUXILIARY.Key + "  " + data.Data.AUXILIARY.Value + "\n"
	message += data.Data.BATTLESHIP.Key + "  " + data.Data.BATTLESHIP.Value + "\n"
	message += data.Data.CRUISER.Key + "  " + data.Data.CRUISER.Value + "\n"
	message += data.Data.DESTROYER.Key + "  " + data.Data.DESTROYER.Value + "\n"
	message += data.Data.SUBMARINE.Key + "  " + data.Data.SUBMARINE.Value + "\n"

	WriteMessage(message, group_id)
}
