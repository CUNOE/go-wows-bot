package model

type NationList struct {
	Code int `json:"code"`
	Data []struct {
		Cn     string `json:"cn"`
		Nation string `json:"nation"`
	} `json:"data"`
	Message string `json:"message"`
}

type ServerList struct {
	Code int `json:"code"`
	Data []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"data"`
	Message string `json:"message"`
}

type ShipType struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		CRUISER struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"CRUISER"`
		BATTLESHIP struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"BATTLESHIP"`
		DESTROYER struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"DESTROYER"`
		SUBMARINE struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"SUBMARINE"`
		AUXILIARY struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"AUXILIARY"`
		AIRCARRIER struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"AIRCARRIER"`
	} `json:"data"`
}
