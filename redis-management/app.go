package main

import (
	"changeme/internal/define"
	"changeme/internal/service"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
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

func (a *App) ConnectionList() gin.H {
	connections, err := service.ConnectionList()
	if err != nil {
		return gin.H{
			"code": -1,
			"msg":  "ERROR" + err.Error(),
		}
	}
	return gin.H{
		"code": 200,
		"data": connections,
	}
}

func (a *App) ConnectionCreate(connection *define.Connection) gin.H {
	err := service.ConnectionCreate(connection)
	if err != nil {
		return gin.H{
			"code": -1,
			"msg":  "ERROR" + err.Error(),
		}
	}
	return gin.H{
		"code": 200,
		"msg":  "创建成功",
	}
}

func (a *App) ConnectionEdit(connection *define.Connection) gin.H {
	err := service.ConnectionEdit(connection)
	if err != nil {
		return gin.H{
			"code": -1,
			"msg":  "ERROR" + err.Error(),
		}
	}
	return gin.H{
		"code": 200,
		"msg":  "编辑成功",
	}
}

func (a *App) ConnectionDelete(id string) gin.H {
	err := service.ConnectionDelete(id)
	if err != nil {
		return gin.H{
			"code": -1,
			"msg":  "ERROR" + err.Error(),
		}
	}
	return gin.H{
		"code": 200,
		"msg":  "删除成功",
	}
}

func (a *App) DbList(id string) gin.H {
	dbs, err := service.DbList(id)
	if err != nil {
		return gin.H{
			"code": -1,
			"msg":  "ERROR" + err.Error(),
		}
	}
	return gin.H{
		"code": 200,
		"data": dbs,
	}
}

func (a *App) KeyList(req *define.KeyListRequest) gin.H {
	if req.ConnectionId == "" {
		return gin.H{
			"code": -1,
			"msg":  "唯一标识不能为空",
		}
	}
	data, err := service.KeyList(req)
	if err != nil {
		return gin.H{
			"code": -1,
			"msg":  "ERROR" + err.Error(),
		}
	}
	return gin.H{
		"code": 200,
		"data": data,
	}
}

func (a *App) GetKeyValue(req *define.KeyValueRequest) gin.H {
	if req.ConnectionId == "" || req.Key == "" {
		return gin.H{
			"code": -1,
			"msg":  "必填参数不能为空",
		}
	}
	data, err := service.GetKeyValue(req)
	if err != nil {
		return gin.H{
			"code": -1,
			"msg":  "ERROR" + err.Error(),
		}
	}
	return gin.H{
		"code": 200,
		"data": data,
	}
}

func (a *App) CreateValue(req *define.CreateKeyValueRequest) gin.H {
	if req.ConnectionId == "" || req.Key == "" || req.Type == "" {
		return gin.H{
			"code": -1,
			"msg":  "必填参数不能为空",
		}
	}
	err := service.CreateKeyValue(req)
	if err != nil {
		return gin.H{
			"code": -1,
			"msg":  "ERROR" + err.Error(),
		}
	}
	return gin.H{
		"code": 200,
		"data": "新增成功",
	}
}

// DeleteKeyValue 键值对删除
func (a *App) DeleteKeyValue(req *define.KeyValueRequest) gin.H {
	if req.Key == "" || req.ConnectionId == "" {
		return gin.H{
			"code": -1,
			"msg":  "必填参不能为空",
		}
	}
	err := service.DeleteKeyValue(req)
	if err != nil {
		return gin.H{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	return gin.H{
		"code": 200,
		"msg":  "删除成功",
	}
}

func (a *App) UpdateKeyValue(req *define.UpdateKeyValueRequest) gin.H {
	if req.Key == "" || req.ConnectionId == "" || req.Value == "" {
		return gin.H{
			"code": -1,
			"msg":  "必填参不能为空",
		}
	}
	err := service.UpdateKeyValue(req)
	if err != nil {
		return gin.H{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	return gin.H{
		"code": 200,
		"msg":  "更新成功",
	}
}
func (a *App) SendEmail(req *define.Email) gin.H {
	if req.Address == "" {
		return gin.H{
			"code": -1,
			"msg":  "必填参不能为空",
		}
	}
	service.SendEmail(req)
	return gin.H{
		"code": 200,
		"msg":  "发送成功",
	}
}

// HashFieldDelete hash字段删除
func (a *App) HashFieldDelete(req *define.HashFieldDeleteRequest) gin.H {
	if req.Key == "" || req.ConnectionId == "" || len(req.Field) == 0 {
		return gin.H{
			"code": -1,
			"msg":  "必填参不能为空",
		}
	}
	err := service.HashFieldDelete(req)
	if err != nil {
		return gin.H{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	return gin.H{
		"code": 200,
		"msg":  "删除成功",
	}
}

// HashAddOrUpdateField hash字段新增、更新
func (a *App) HashAddOrUpdateField(req *define.HashAddOrUpdateFieldRequest) gin.H {
	if req.Key == "" || req.ConnectionId == "" || req.Field == "" {
		return gin.H{
			"code": -1,
			"msg":  "必填参不能为空",
		}
	}
	err := service.HashAddOrUpdateField(req)
	if err != nil {
		return gin.H{
			"code": -1,
			"msg":  "ERROR" + err.Error(),
		}
	}
	return gin.H{
		"code": 200,
		"data": "修改成功",
	}
}

// ListValueDelete 列表值删除
func (a *App) ListValueDelete(req *define.ListValueRequest) gin.H {
	if req.Key == "" || req.ConnectionId == "" || req.Value == "" {
		return gin.H{
			"code": -1,
			"msg":  "必填参不能为空",
		}
	}
	err := service.ListValueDelete(req)
	if err != nil {
		return gin.H{
			"code": -1,
			"msg":  "ERROR" + err.Error(),
		}
	}
	return gin.H{
		"code": 200,
		"data": "删除成功",
	}
}

// ListValueCreate 列表值新增
func (a *App) ListValueCreate(req *define.ListValueRequest) gin.H {
	if req.Key == "" || req.ConnectionId == "" || req.Value == "" {
		return gin.H{
			"code": -1,
			"msg":  "必填参不能为空",
		}
	}
	err := service.ListValueCreate(req)
	if err != nil {
		return gin.H{
			"code": -1,
			"msg":  "ERROR" + err.Error(),
		}
	}
	return gin.H{
		"code": 200,
		"data": "新增成功",
	}
}

// SetValueDelete 集合值删除
func (a *App) SetValueDelete(req *define.SetValueRequest) gin.H {
	if req.Key == "" || req.ConnectionId == "" || req.Value == "" {
		return gin.H{
			"code": -1,
			"msg":  "必填参不能为空",
		}
	}
	err := service.SetValueDelete(req)
	if err != nil {
		return gin.H{
			"code": -1,
			"msg":  "ERROR" + err.Error(),
		}
	}
	return gin.H{
		"code": 200,
		"data": "删除成功",
	}
}

// SetValueCreate 集合新增
func (a *App) SetValueCreate(req *define.SetValueRequest) gin.H {
	if req.Key == "" || req.ConnectionId == "" || req.Value == "" {
		return gin.H{
			"code": -1,
			"msg":  "必填参不能为空",
		}
	}
	err := service.SetValueCreate(req)
	if err != nil {
		return gin.H{
			"code": -1,
			"msg":  "ERROR" + err.Error(),
		}
	}
	return gin.H{
		"code": 200,
		"data": "新增成功",
	}
}

// ZSetValueDelete 有序集合值删除
func (a *App) ZSetValueDelete(req *define.ZSetValueRequest) gin.H {
	if req.Key == "" || req.ConnectionId == "" || req.Member == nil {
		return gin.H{
			"code": -1,
			"msg":  "必填参不能为空",
		}
	}
	err := service.ZSetValueDelete(req)
	if err != nil {
		return gin.H{
			"code": -1,
			"msg":  "ERROR" + err.Error(),
		}
	}
	return gin.H{
		"code": 200,
		"data": "删除成功",
	}
}

// ZSetValueCreate 有序集合新增
func (a *App) ZSetValueCreate(req *define.ZSetValueRequest) gin.H {
	if req.Key == "" || req.ConnectionId == "" || req.Member == nil {
		return gin.H{
			"code": -1,
			"msg":  "必填参不能为空",
		}
	}
	err := service.ZSetValueCreate(req)
	if err != nil {
		return gin.H{
			"code": -1,
			"msg":  "ERROR" + err.Error(),
		}
	}
	return gin.H{
		"code": 200,
		"data": "新增成功",
	}
}
