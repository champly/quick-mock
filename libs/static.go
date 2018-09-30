package libs

import "net/http"

type StaticMethod struct {
}

func (s *StaticMethod) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
