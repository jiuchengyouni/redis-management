package service

import (
	"changeme/internal/define"
	"changeme/internal/utils"
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"strconv"
	"strings"
)

func DbList(id string) (DbItems []*define.DbItem, err error) {
	if id == "" {
		return nil, errors.New("连接唯一标识不能为空")
	}
	conn, err := utils.GetConnection(id)
	if err != nil {
		return nil, err
	}
	rdb := redis.NewClient(&redis.Options{
		Network:         "",
		Addr:            conn.Address + ":" + conn.Port,
		Dialer:          nil,
		OnConnect:       nil,
		Username:        conn.Username,
		Password:        conn.Password,
		DB:              0,
		MaxRetries:      0,
		MinRetryBackoff: 0,
		MaxRetryBackoff: 0,
		DialTimeout:     0,
		ReadTimeout:     0,
		WriteTimeout:    0,
		PoolFIFO:        false,
		PoolSize:        0,
		PoolTimeout:     0,
		MinIdleConns:    0,
		TLSConfig:       nil,
		Limiter:         nil,
	})
	keySpace, err := rdb.Info(context.Background(), "keyspace").Result()
	if err != nil {
		return nil, err
	}
	// keyspace 数据格式
	// # Keyspace
	// db0:keys=2,avg_ttl...
	//
	m := make(map[string]int)
	v := strings.Split(keySpace, "\n")
	for i := 1; i < len(v)-1; i++ {
		databases := strings.Split(v[i], ":")
		if len(databases) < 2 {
			continue
		}
		vv := strings.Split(databases[1], ",")
		if len(vv) < 1 {
			continue
		}
		keyNumber := strings.Split(vv[0], "=")
		if len(keyNumber) < 2 {
			continue
		}
		num, err := strconv.Atoi(keyNumber[1])
		if err != nil {
			continue
		}
		m[databases[0]] = num
	}
	// config get 获取数据库的个数
	databasesRes, err := rdb.ConfigGet(context.Background(), "databases").Result()
	if err != nil {
		return nil, err
	}
	if len(databasesRes) < 2 {
		return nil, errors.New("连接数据异常")
	}
	dbNum, err := strconv.Atoi(databasesRes[1].(string))
	if err != nil {
		return nil, err
	}
	data := make([]*define.DbItem, 0)
	for i := 0; i < dbNum; i++ {
		item := &define.DbItem{
			Key: "db" + strconv.Itoa(i),
		}
		if n, ok := m["db"+strconv.Itoa(i)]; ok {
			item.Number = n
		}
		data = append(data, item)
	}
	return data, nil
}
