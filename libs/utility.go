package libs

import "github.com/ChamPly/quick-mock/model"
import "strings"
import "fmt"

func GetRouterByURL(url string, method string, routers []model.Router) (router model.Router, err error) {
	for _, item := range routers {
		if (url == item.RouterName) && strings.ToUpper(item.Method) == method {
			router = item
			return
		}
	}
	fmt.Println("not found router:", url)

	err = fmt.Errorf("error router")
	return
}
