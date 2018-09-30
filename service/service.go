package service

import (
	"net/http"
	"strings"

	"github.com/champly/quick-mock/libs"
	"github.com/champly/quick-mock/model"
)

type MockService struct {
	Property      model.Property
	StaticHandler http.Handler
}

func NewMockService(pro model.Property) *MockService {
	var staticHandler http.Handler

	return &MockService{
		Property:      pro,
		StaticHandler: staticHandler,
	}
}

func (m *MockService) Start() {
	if m.Property.StaticPro.Path != "" {
		println("设置静态文件")
		// http.HandleFunc(m.Property.StaticPro.Router, m.fileHandler)
		m.setStaticFile()
	}
	http.HandleFunc("/", m.handler)
	http.ListenAndServe(":"+m.Property.Port, nil)
}

func (m *MockService) handler(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, m.Property.StaticPro.Router) {
		m.StaticHandler.ServeHTTP(w, r)
	}

	router, err := libs.GetRouterByURL(r.URL.Path, r.Method, m.Property.Routers)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	responseStr, statusCode := libs.AnalysisResponse(router)

	if statusCode == 302 || statusCode == 301 {
		http.Redirect(w, r, responseStr, statusCode)
		return
	}

	w.WriteHeader(statusCode)
	w.Write([]byte(strings.Replace(responseStr, "'", "\"", -1)))
}

func (m *MockService) setStaticFile() {
	// dir, err := os.Getwd()
	// if err != nil {
	// 	println(err)
	// 	w.WriteHeader(http.StatusServiceUnavailable)
	// }
	// http.StripPrefix(m.Property.StaticPro.Router, http.FileServer(http.Dir(m.Property.StaticPro.Path))).ServeHTTP(w, r)
	m.StaticHandler = http.FileServer(http.Dir(m.Property.StaticPro.Path))
}
