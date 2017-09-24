package model

type Property struct {
	Port    string   `json:"port"`
	Routers []Router `json:"routers"`
}

type Router struct {
	RouterName string `json:"name"`
	Request    string `json:"request"`
	Method     string `json:"method"`
	Response   string `json:"response"`
}
