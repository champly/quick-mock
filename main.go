package main

import (
	"fmt"
	"runtime"

	"github.com/ChamPly/quick-mock/libs"
	"github.com/ChamPly/quick-mock/service"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	pro, err := libs.ReadProperty()
	if err != nil {
		fmt.Println(err)
	}
	mockService := service.NewMockService(pro)
	mockService.Start()
}
