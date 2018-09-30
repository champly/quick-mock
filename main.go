package main

import (
	"fmt"
	"runtime"

	"github.com/champly/quick-mock/libs"
	"github.com/champly/quick-mock/service"
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
