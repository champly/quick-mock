package service

import (
	"net/http"
	"strings"

	"github.com/ChamPly/quick-mock/libs"
	"github.com/ChamPly/quick-mock/model"
)

type MockService struct {
	Property model.Property
}

func NewMockService(pro model.Property) *MockService {
	return &MockService{
		Property: pro,
	}
}

func (m *MockService) Start() {
	http.HandleFunc("/", m.handler)
	http.ListenAndServe(":"+m.Property.Port, nil)
}

func (m *MockService) handler(w http.ResponseWriter, r *http.Request) {
	router, err := libs.GetRouterByURL(r.RequestURI, r.Method, m.Property.Routers)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Write([]byte(strings.Replace(router.Response, "'", "\"", -1)))
}
