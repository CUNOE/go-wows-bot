package model

type BindInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		PlatformType string `json:"platformType"`
		PlatformId   string `json:"platformId"`
		AccountId    int    `json:"accountId"`
		UserName     string `json:"userName"`
		ServerType   string `json:"serverType"`
		DefaultId    bool   `json:"defaultId"`
	} `json:"data"`
}

type BindBack struct {
	Code    int    `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}
