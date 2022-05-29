package model

type RecentInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ServerName string `json:"serverName"`
		ClanInfo   struct {
			ClanId           int    `json:"clanId"`
			Tag              string `json:"tag"`
			Name             string `json:"name"`
			Description      string `json:"description"`
			MembersCount     int    `json:"membersCount"`
			CreatorName      string `json:"creatorName"`
			CreatorAccountId int    `json:"creatorAccountId"`
			LeaderName       string `json:"leaderName"`
			LeaderAccountId  int    `json:"leaderAccountId"`
			ColorRgb         string `json:"colorRgb"`
			CreatedAt        int    `json:"createdAt"`
			UpdatedAt        int    `json:"updatedAt"`
		} `json:"clanInfo"`
		UserName string `json:"userName"`
		Data     struct {
			ShipId interface{} `json:"shipId"`
			Pr     struct {
				Value       int    `json:"value"`
				NextValue   int    `json:"nextValue"`
				Name        string `json:"name"`
				EnglishName string `json:"englishName"`
				Color       string `json:"color"`
			} `json:"pr"`
			ShipInfo struct {
				NameCn      interface{} `json:"nameCn"`
				NameEnglish interface{} `json:"nameEnglish"`
				Level       interface{} `json:"level"`
				ShipType    interface{} `json:"shipType"`
				Country     interface{} `json:"country"`
				ImgSmall    interface{} `json:"imgSmall"`
				ImgLarge    interface{} `json:"imgLarge"`
				ImgMedium   interface{} `json:"imgMedium"`
			} `json:"shipInfo"`
			Battles           int         `json:"battles"`
			Wins              float64     `json:"wins"`
			Damage            int         `json:"damage"`
			Xp                int         `json:"xp"`
			Kd                float64     `json:"kd"`
			Hit               float64     `json:"hit"`
			Frags             float64     `json:"frags"`
			LastBattlesTime   interface{} `json:"lastBattlesTime"`
			ExtensionDataInfo struct {
				MaxDamageShipId         interface{} `json:"maxDamageShipId"`
				MaxDamage               int         `json:"maxDamage"`
				MaxDamageScoutingShipId interface{} `json:"maxDamageScoutingShipId"`
				MaxDamageScouting       int         `json:"maxDamageScouting"`
				MaxTotalAgroShipId      interface{} `json:"maxTotalAgroShipId"`
				MaxTotalAgro            int         `json:"maxTotalAgro"`
				MaxXpShipId             interface{} `json:"maxXpShipId"`
				MaxXp                   int         `json:"maxXp"`
				MaxFragsShipId          interface{} `json:"maxFragsShipId"`
				MaxFrags                int         `json:"maxFrags"`
				MaxPlanesKilledShipId   interface{} `json:"maxPlanesKilledShipId"`
				MaxPlanesKilled         int         `json:"maxPlanesKilled"`
			} `json:"extensionDataInfo"`
		} `json:"data"`
		RecentList []struct {
			ShipId int64 `json:"shipId"`
			Pr     struct {
				Value       int    `json:"value"`
				NextValue   int    `json:"nextValue"`
				Name        string `json:"name"`
				EnglishName string `json:"englishName"`
				Color       string `json:"color"`
			} `json:"pr"`
			ShipInfo struct {
				NameCn      string `json:"nameCn"`
				NameEnglish string `json:"nameEnglish"`
				Level       int    `json:"level"`
				ShipType    string `json:"shipType"`
				Country     string `json:"country"`
				ImgSmall    string `json:"imgSmall"`
				ImgLarge    string `json:"imgLarge"`
				ImgMedium   string `json:"imgMedium"`
			} `json:"shipInfo"`
			Battles           int     `json:"battles"`
			Wins              float64 `json:"wins"`
			Damage            int     `json:"damage"`
			Xp                int     `json:"xp"`
			Kd                float64 `json:"kd"`
			Hit               float64 `json:"hit"`
			Frags             float64 `json:"frags"`
			LastBattlesTime   int     `json:"lastBattlesTime"`
			ExtensionDataInfo struct {
				MaxDamageShipId         interface{} `json:"maxDamageShipId"`
				MaxDamage               int         `json:"maxDamage"`
				MaxDamageScoutingShipId interface{} `json:"maxDamageScoutingShipId"`
				MaxDamageScouting       int         `json:"maxDamageScouting"`
				MaxTotalAgroShipId      interface{} `json:"maxTotalAgroShipId"`
				MaxTotalAgro            int         `json:"maxTotalAgro"`
				MaxXpShipId             interface{} `json:"maxXpShipId"`
				MaxXp                   int         `json:"maxXp"`
				MaxFragsShipId          interface{} `json:"maxFragsShipId"`
				MaxFrags                int         `json:"maxFrags"`
				MaxPlanesKilledShipId   interface{} `json:"maxPlanesKilledShipId"`
				MaxPlanesKilled         int         `json:"maxPlanesKilled"`
			} `json:"extensionDataInfo"`
		} `json:"recentList"`
		LastBattlesTime int `json:"lastBattlesTime"`
		RecordTime      int `json:"recordTime"`
	} `json:"data"`
}
