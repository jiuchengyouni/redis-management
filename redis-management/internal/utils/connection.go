package utils

import (
	"changeme/internal/define"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

func GetConnection(id string) (*define.Connection, error) {
	conf := new(define.Config)
	nowPath, _ := os.Getwd()
	data, err := ioutil.ReadFile(nowPath + string(os.PathSeparator) + define.ConfigName)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, conf)
	for _, v := range conf.Connections {
		if v.Id == id {
			return v, nil
		}
	}
	return nil, errors.New("连接数据不存在")
}
