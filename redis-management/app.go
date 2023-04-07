package main

import (
	"changeme/internal/define"
	"changeme/internal/service"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) ConnectionList() H {
	connections, err := service.ConnectionList()
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR" + err.Error(),
		}
	}
	return M{
		"code": 200,
		"data": connections,
	}
}

func (a *App) ConnectionCreate(connection *define.Connection) interface{} {
	err := service.ConnectionCreate(connection)
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR" + err.Error(),
		}
	}
	return M{
		"code": 200,
		"msg":  "创建成功",
	}
}

func (a *App) ConnectionEdit(connection *define.Connection) interface{} {
	err := service.ConnectionEdit(connection)
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR" + err.Error(),
		}
	}
	return M{
		"code": 200,
		"msg":  "编辑成功",
	}
}

func (a *App) ConnectionDelete(id string) interface{} {
	err := service.ConnectionDelete(id)
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR" + err.Error(),
		}
	}
	return M{
		"code": 200,
		"msg":  "删除成功",
	}
}

func (a *App) DbList(id string) H {
	dbs, err := service.DbList(id)
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR" + err.Error(),
		}
	}
	return M{
		"code": 200,
		"data": dbs,
	}
}

func (a *App) KeyList(req *define.KeyListRequest) H {
	if req.ConnectionId == "" {
		return M{
			"code": -1,
			"msg":  "唯一标识不能为空",
		}
	}
	data, err := service.KeyList(req)
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR" + err.Error(),
		}
	}
	return M{
		"code": 200,
		"data": data,
	}
}

func (a *App) GetKeyValue(req *define.KeyValueRequest) H {
	if req.ConnectionId == "" || req.Key == "" {
		return M{
			"code": -1,
			"msg":  "必填参数不能为空",
		}
	}
	data, err := service.GetKeyValue(req)
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR" + err.Error(),
		}
	}
	return M{
		"code": 200,
		"data": data,
	}
}

func (a *App) CreateValue(req *define.CreateKeyValueRequest) H {
	if req.ConnectionId == "" || req.Key == "" || req.Type == "" {
		return M{
			"code": -1,
			"msg":  "必填参数不能为空",
		}
	}
	err := service.CreateKeyValue(req)
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR" + err.Error(),
		}
	}
	return M{
		"code": 200,
		"data": "新增成功",
	}
}

// DeleteKeyValue 键值对删除
func (a *App) DeleteKeyValue(req *define.KeyValueRequest) H {
	if req.Key == "" || req.ConnectionId == "" {
		return M{
			"code": -1,
			"msg":  "必填参不能为空",
		}
	}
	err := service.DeleteKeyValue(req)
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	return M{
		"code": 200,
		"msg":  "删除成功",
	}
}

func (a *App) UpdateKeyValue(req *define.UpdateKeyValueRequest) H {
	if req.Key == "" || req.ConnectionId == "" || req.Value == "" {
		return M{
			"code": -1,
			"msg":  "必填参不能为空",
		}
	}
	err := service.UpdateKeyValue(req)
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	return M{
		"code": 200,
		"msg":  "更新成功",
	}
}
