package model

type SearchUserInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		PlatformType interface{} `json:"platformType"`
		PlatformId   interface{} `json:"platformId"`
		AccountId    int         `json:"accountId"`
		UserName     string      `json:"userName"`
		ServerType   interface{} `json:"serverType"`
		DefaultId    interface{} `json:"defaultId"`
	} `json:"data"`
}

type UserInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		DwpDataVO struct {
			Damage int     `json:"damage"`
			Wins   float64 `json:"wins"`
			Pr     int     `json:"pr"`
		} `json:"dwpDataVO"`
		Karma    int `json:"karma"`
		ClanInfo struct {
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
		UserName   string `json:"userName"`
		ServerName string `json:"serverName"`
		Pr         struct {
			Value       int    `json:"value"`
			NextValue   int    `json:"nextValue"`
			Name        string `json:"name"`
			EnglishName string `json:"englishName"`
			Color       string `json:"color"`
		} `json:"pr"`
		LastDateTime int `json:"lastDateTime"`
		Type         struct {
			Cruiser struct {
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
			} `json:"Cruiser"`
			Auxiliary struct {
				ShipId interface{} `json:"shipId"`
				Pr     struct {
					Value       interface{} `json:"value"`
					NextValue   interface{} `json:"nextValue"`
					Name        interface{} `json:"name"`
					EnglishName interface{} `json:"englishName"`
					Color       interface{} `json:"color"`
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
				ExtensionDataInfo interface{} `json:"extensionDataInfo"`
			} `json:"Auxiliary"`
			Destroyer struct {
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
			} `json:"Destroyer"`
			Submarine struct {
				ShipId interface{} `json:"shipId"`
				Pr     struct {
					Value       interface{} `json:"value"`
					NextValue   interface{} `json:"nextValue"`
					Name        interface{} `json:"name"`
					EnglishName interface{} `json:"englishName"`
					Color       interface{} `json:"color"`
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
				ExtensionDataInfo interface{} `json:"extensionDataInfo"`
			} `json:"Submarine"`
			Battleship struct {
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
			} `json:"Battleship"`
			AirCarrier struct {
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
			} `json:"AirCarrier"`
		} `json:"type"`
		Pvp struct {
			ShipId int `json:"shipId"`
			Pr     struct {
				Value       interface{} `json:"value"`
				NextValue   interface{} `json:"nextValue"`
				Name        interface{} `json:"name"`
				EnglishName interface{} `json:"englishName"`
				Color       interface{} `json:"color"`
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
			Battles           int     `json:"battles"`
			Wins              float64 `json:"wins"`
			Damage            int     `json:"damage"`
			Xp                int     `json:"xp"`
			Kd                float64 `json:"kd"`
			Hit               float64 `json:"hit"`
			Frags             float64 `json:"frags"`
			LastBattlesTime   int     `json:"lastBattlesTime"`
			ExtensionDataInfo struct {
				MaxDamageShipId         int64 `json:"maxDamageShipId"`
				MaxDamage               int   `json:"maxDamage"`
				MaxDamageScoutingShipId int64 `json:"maxDamageScoutingShipId"`
				MaxDamageScouting       int   `json:"maxDamageScouting"`
				MaxTotalAgroShipId      int64 `json:"maxTotalAgroShipId"`
				MaxTotalAgro            int   `json:"maxTotalAgro"`
				MaxXpShipId             int64 `json:"maxXpShipId"`
				MaxXp                   int   `json:"maxXp"`
				MaxFragsShipId          int64 `json:"maxFragsShipId"`
				MaxFrags                int   `json:"maxFrags"`
				MaxPlanesKilledShipId   int64 `json:"maxPlanesKilledShipId"`
				MaxPlanesKilled         int   `json:"maxPlanesKilled"`
			} `json:"extensionDataInfo"`
		} `json:"pvp"`
		PvpSolo struct {
			ShipId int `json:"shipId"`
			Pr     struct {
				Value       interface{} `json:"value"`
				NextValue   interface{} `json:"nextValue"`
				Name        interface{} `json:"name"`
				EnglishName interface{} `json:"englishName"`
				Color       interface{} `json:"color"`
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
			Battles           int     `json:"battles"`
			Wins              float64 `json:"wins"`
			Damage            int     `json:"damage"`
			Xp                int     `json:"xp"`
			Kd                float64 `json:"kd"`
			Hit               float64 `json:"hit"`
			Frags             float64 `json:"frags"`
			LastBattlesTime   int     `json:"lastBattlesTime"`
			ExtensionDataInfo struct {
				MaxDamageShipId         int64 `json:"maxDamageShipId"`
				MaxDamage               int   `json:"maxDamage"`
				MaxDamageScoutingShipId int64 `json:"maxDamageScoutingShipId"`
				MaxDamageScouting       int   `json:"maxDamageScouting"`
				MaxTotalAgroShipId      int64 `json:"maxTotalAgroShipId"`
				MaxTotalAgro            int   `json:"maxTotalAgro"`
				MaxXpShipId             int64 `json:"maxXpShipId"`
				MaxXp                   int   `json:"maxXp"`
				MaxFragsShipId          int64 `json:"maxFragsShipId"`
				MaxFrags                int   `json:"maxFrags"`
				MaxPlanesKilledShipId   int64 `json:"maxPlanesKilledShipId"`
				MaxPlanesKilled         int   `json:"maxPlanesKilled"`
			} `json:"extensionDataInfo"`
		} `json:"pvpSolo"`
		PvpTwo struct {
			ShipId int `json:"shipId"`
			Pr     struct {
				Value       interface{} `json:"value"`
				NextValue   interface{} `json:"nextValue"`
				Name        interface{} `json:"name"`
				EnglishName interface{} `json:"englishName"`
				Color       interface{} `json:"color"`
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
			Battles           int     `json:"battles"`
			Wins              float64 `json:"wins"`
			Damage            int     `json:"damage"`
			Xp                int     `json:"xp"`
			Kd                float64 `json:"kd"`
			Hit               float64 `json:"hit"`
			Frags             float64 `json:"frags"`
			LastBattlesTime   int     `json:"lastBattlesTime"`
			ExtensionDataInfo struct {
				MaxDamageShipId         int64 `json:"maxDamageShipId"`
				MaxDamage               int   `json:"maxDamage"`
				MaxDamageScoutingShipId int64 `json:"maxDamageScoutingShipId"`
				MaxDamageScouting       int   `json:"maxDamageScouting"`
				MaxTotalAgroShipId      int64 `json:"maxTotalAgroShipId"`
				MaxTotalAgro            int   `json:"maxTotalAgro"`
				MaxXpShipId             int64 `json:"maxXpShipId"`
				MaxXp                   int   `json:"maxXp"`
				MaxFragsShipId          int64 `json:"maxFragsShipId"`
				MaxFrags                int   `json:"maxFrags"`
				MaxPlanesKilledShipId   int64 `json:"maxPlanesKilledShipId"`
				MaxPlanesKilled         int   `json:"maxPlanesKilled"`
			} `json:"extensionDataInfo"`
		} `json:"pvpTwo"`
		PvpThree struct {
			ShipId int `json:"shipId"`
			Pr     struct {
				Value       interface{} `json:"value"`
				NextValue   interface{} `json:"nextValue"`
				Name        interface{} `json:"name"`
				EnglishName interface{} `json:"englishName"`
				Color       interface{} `json:"color"`
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
			Battles           int     `json:"battles"`
			Wins              float64 `json:"wins"`
			Damage            int     `json:"damage"`
			Xp                int     `json:"xp"`
			Kd                float64 `json:"kd"`
			Hit               float64 `json:"hit"`
			Frags             float64 `json:"frags"`
			LastBattlesTime   int     `json:"lastBattlesTime"`
			ExtensionDataInfo struct {
				MaxDamageShipId         int64 `json:"maxDamageShipId"`
				MaxDamage               int   `json:"maxDamage"`
				MaxDamageScoutingShipId int64 `json:"maxDamageScoutingShipId"`
				MaxDamageScouting       int   `json:"maxDamageScouting"`
				MaxTotalAgroShipId      int64 `json:"maxTotalAgroShipId"`
				MaxTotalAgro            int   `json:"maxTotalAgro"`
				MaxXpShipId             int64 `json:"maxXpShipId"`
				MaxXp                   int   `json:"maxXp"`
				MaxFragsShipId          int64 `json:"maxFragsShipId"`
				MaxFrags                int   `json:"maxFrags"`
				MaxPlanesKilledShipId   int64 `json:"maxPlanesKilledShipId"`
				MaxPlanesKilled         int   `json:"maxPlanesKilled"`
			} `json:"extensionDataInfo"`
		} `json:"pvpThree"`
		BattleCount struct {
			Cruiser struct {
				Field1  int `json:"1"`
				Field2  int `json:"2"`
				Field3  int `json:"3"`
				Field4  int `json:"4"`
				Field5  int `json:"5"`
				Field6  int `json:"6"`
				Field7  int `json:"7"`
				Field8  int `json:"8"`
				Field9  int `json:"9"`
				Field10 int `json:"10"`
				Field11 int `json:"11"`
			} `json:"Cruiser"`
			Auxiliary struct {
				Field1  int `json:"1"`
				Field2  int `json:"2"`
				Field3  int `json:"3"`
				Field4  int `json:"4"`
				Field5  int `json:"5"`
				Field6  int `json:"6"`
				Field7  int `json:"7"`
				Field8  int `json:"8"`
				Field9  int `json:"9"`
				Field10 int `json:"10"`
				Field11 int `json:"11"`
			} `json:"Auxiliary"`
			Destroyer struct {
				Field1  int `json:"1"`
				Field2  int `json:"2"`
				Field3  int `json:"3"`
				Field4  int `json:"4"`
				Field5  int `json:"5"`
				Field6  int `json:"6"`
				Field7  int `json:"7"`
				Field8  int `json:"8"`
				Field9  int `json:"9"`
				Field10 int `json:"10"`
				Field11 int `json:"11"`
			} `json:"Destroyer"`
			Submarine struct {
				Field1  int `json:"1"`
				Field2  int `json:"2"`
				Field3  int `json:"3"`
				Field4  int `json:"4"`
				Field5  int `json:"5"`
				Field6  int `json:"6"`
				Field7  int `json:"7"`
				Field8  int `json:"8"`
				Field9  int `json:"9"`
				Field10 int `json:"10"`
				Field11 int `json:"11"`
			} `json:"Submarine"`
			Battleship struct {
				Field1  int `json:"1"`
				Field2  int `json:"2"`
				Field3  int `json:"3"`
				Field4  int `json:"4"`
				Field5  int `json:"5"`
				Field6  int `json:"6"`
				Field7  int `json:"7"`
				Field8  int `json:"8"`
				Field9  int `json:"9"`
				Field10 int `json:"10"`
				Field11 int `json:"11"`
			} `json:"Battleship"`
			AirCarrier struct {
				Field1  int `json:"1"`
				Field2  int `json:"2"`
				Field3  int `json:"3"`
				Field4  int `json:"4"`
				Field5  int `json:"5"`
				Field6  int `json:"6"`
				Field7  int `json:"7"`
				Field8  int `json:"8"`
				Field9  int `json:"9"`
				Field10 int `json:"10"`
				Field11 int `json:"11"`
			} `json:"AirCarrier"`
		} `json:"battleCount"`
		BattleCountAll struct {
			Field1  int `json:"1"`
			Field2  int `json:"2"`
			Field3  int `json:"3"`
			Field4  int `json:"4"`
			Field5  int `json:"5"`
			Field6  int `json:"6"`
			Field7  int `json:"7"`
			Field8  int `json:"8"`
			Field9  int `json:"9"`
			Field10 int `json:"10"`
			Field11 int `json:"11"`
		} `json:"battleCountAll"`
	} `json:"data"`
}

type ShipInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		DwpDataVO struct {
			Damage int     `json:"damage"`
			Wins   float64 `json:"wins"`
			Pr     int     `json:"pr"`
		} `json:"dwpDataVO"`
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
		ShipInfo struct {
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
		} `json:"shipInfo"`
	} `json:"data"`
	QueryTime int `json:"queryTime"`
}
