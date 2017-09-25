package libs

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ChamPly/quick-mock/model"
)

func GetRouterByURL(url string, method string, routers []model.Router) (router model.Router, err error) {
	for _, item := range routers {
		if (url == item.RouterName) && strings.Contains(strings.ToUpper(item.Method), method) {
			router = item
			return
		}
	}
	fmt.Println("not found router:", url)

	err = fmt.Errorf("error router")
	return
}

func AnalysisResponse(router model.Router) (responseStr string, statusCode int) {
	responseStr = router.Response

	if router.StatusCode == "" {
		statusCode = 200
		return
	}
	statusCode, err := strconv.Atoi(router.StatusCode)
	if err != nil {
		responseStr = "status_code参数必须为数字"
		statusCode = 500
	}

	return
}
