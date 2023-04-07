package service

import (
	"changeme/internal/define"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"io/ioutil"
	"os"
)

func ConnectionList() (connections []*define.Connection, err error) {
	nowPath, _ := os.Getwd()
	data, err := ioutil.ReadFile(nowPath + string(os.PathSeparator) + define.ConfigName)
	if errors.Is(err, os.ErrExist) {
		return nil, nil
	}
	conf := new(define.Config)
	//data数据写入config
	err = json.Unmarshal(data, conf)
	if err != nil {
		return nil, err
	}
	return conf.Connections, nil
}

func ConnectionCreate(connection *define.Connection) (err error) {
	if connection.Address == "" {
		return errors.New("连接地址不能为空")
	}
	if connection.Name == "" {
		connection.Name = connection.Address
	}
	if connection.Port == "" {
		connection.Port = "6379"
	}
	connection.Id = uuid.New().String()
	conf := new(define.Config)
	nowPath, _ := os.Getwd()
	data, err := ioutil.ReadFile(nowPath + string(os.PathSeparator) + define.ConfigName)
	if errors.Is(err, os.ErrExist) {
		//初始化内容
		conf.Connections = []*define.Connection{connection}
		data, _ = json.Marshal(conf)
		//配置内容导入
		os.MkdirAll(nowPath, 0666)
		ioutil.WriteFile(nowPath+string(os.PathSeparator)+define.ConfigName, data, 0666)
	}
	json.Unmarshal(data, conf)
	conf.Connections = append(conf.Connections, connection)
	data, _ = json.Marshal(conf)
	ioutil.WriteFile(nowPath+string(os.PathSeparator)+define.ConfigName, data, 0666)
	return nil
}

func ConnectionEdit(connection *define.Connection) (err error) {
	if connection.Id == "" {
		return errors.New("标识不能为空")
	}
	if connection.Address == "" {
		return errors.New("连接地址不能为空")
	}
	if connection.Name == "" {
		connection.Name = connection.Address
	}
	if connection.Port == "" {
		connection.Port = "6379"
	}
	conf := new(define.Config)
	nowPath, _ := os.Getwd()
	data, err := ioutil.ReadFile(nowPath + string(os.PathSeparator) + define.ConfigName)
	if err != nil {
		return err
	}
	json.Unmarshal(data, conf)
	for i, v := range conf.Connections {
		if v.Id == connection.Id {
			conf.Connections[i] = connection
		}
	}
	data, _ = json.Marshal(conf)
	ioutil.WriteFile(nowPath+string(os.PathSeparator)+define.ConfigName, data, 0666)
	return nil
}

func ConnectionDelete(id string) (err error) {
	if id == "" {
		return errors.New("唯一标识不能为空")
	}
	conf := new(define.Config)
	nowPath, _ := os.Getwd()
	data, err := ioutil.ReadFile(nowPath + string(os.PathSeparator) + define.ConfigName)
	if err != nil {
		return err
	}
	json.Unmarshal(data, conf)
	for i, v := range conf.Connections {
		if v.Id == id {
			conf.Connections = append(conf.Connections[:i], conf.Connections[i+1:]...)
			break
		}
	}
	data, _ = json.Marshal(conf)
	ioutil.WriteFile(nowPath+string(os.PathSeparator)+define.ConfigName, data, 0666)
	return nil
}
