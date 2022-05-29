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

type ShipSearch struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		Id                int64       `json:"id"`
		ShipNameNumbers   string      `json:"shipNameNumbers"`
		ShipNameCn        string      `json:"shipNameCn"`
		AnotherName       interface{} `json:"anotherName"`
		ShipNameUpperCase interface{} `json:"shipNameUpperCase"`
		Name              interface{} `json:"name"`
		Country           string      `json:"country"`
		ShipType          string      `json:"shipType"`
		Tier              int         `json:"tier"`
		ShipIdValue       string      `json:"shipIdValue"`
		Special           interface{} `json:"special"`
		ImgSmall          string      `json:"imgSmall"`
		ImgLarge          string      `json:"imgLarge"`
		ImgMedium         string      `json:"imgMedium"`
	} `json:"data"`
	QueryTime int `json:"queryTime"`
}
