package wows

func Handle(message string, group_id int64) {
	switch {
	case message == "wws server":
		wwsServer(group_id)
	case message == "wws nation":
		wwsNation(group_id)
	case message == "wws ship.type":
		wwsShipType(group_id)
	}
}
