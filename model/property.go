package model

type Property struct {
	Port    string   `json:"port"`
	Routers []Router `json:"routers"`
}

type Router struct {
	RouterName string `json:"name"`
	Request    string `json:"request"`
	Method     string `json:"method"`
	StatusCode string `json:"status_code"`
	Response   string `json:"response"`
}
