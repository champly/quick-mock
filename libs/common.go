package libs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/champly/quick-mock/conf"
	"github.com/champly/quick-mock/model"
)

func ReadProperty() (property model.Property, err error) {
	path := fmt.Sprintf("%s/%s", GetLocalPath(), conf.PropertityFileName)
	fileInput, err := os.Open(path)
	if err != nil {
		err = fmt.Errorf("打开文件错误,确认文件:%s存在,错误:%+v", path, err)
	}
	contentByte, err := ioutil.ReadAll(fileInput)
	if err != nil {
		err = fmt.Errorf("读取文件:%s错误:%+v", path, err)
	}

	property = model.Property{}
	err = json.Unmarshal(contentByte, &property)
	if err != nil {
		err = fmt.Errorf("配置文件格式错误，具体配置详见：readme")
	}
	return
}

func GetLocalPath() (path string) {
	file, _ := exec.LookPath(os.Args[0])
	path, _ = filepath.Abs(file)
	path = filepath.Dir(path)
	return
}
