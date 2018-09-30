package model

type Property struct {
	Port      string         `json:"port"`
	Routers   []Router       `json:"routers"`
	StaticPro StaticProperty `json:"static_property"`
}

type Router struct {
	RouterName string `json:"name"`
	Request    string `json:"request"`
	Method     string `json:"method"`
	StatusCode string `json:"status_code"`
	Response   string `json:"response"`
}

type StaticProperty struct {
	Path   string `json:"path"`
	Router string `json:"router"`
}
