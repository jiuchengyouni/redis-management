package service

import (
	"changeme/internal/define"
	"changeme/internal/utils"
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

func KeyList(req *define.KeyListRequest) ([]string, error) {
	conn, err := utils.GetConnection(req.ConnectionId)
	if err != nil {
		return nil, err
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:            conn.Address + ":" + conn.Port,
		Username:        conn.Username,
		Password:        conn.Password,
		DB:              req.Db,
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
	var count int64 = 100
	if req.Keyword != "" {
		count = 2000
	}
	result, _, err := rdb.Scan(context.Background(), 0, "*"+req.Keyword+"*", count).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}
func GetKeyValue(req *define.KeyValueRequest) (*define.KeyValueReply, error) {
	conn, err := utils.GetConnection(req.ConnectionId)
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
		DB:              req.Db,
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
	//获取类型
	_type, err := rdb.Type(context.Background(), req.Key).Result()
	if err != nil {
		return nil, err
	}
	//目前只针对string类型
	v, err := rdb.Get(context.Background(), req.Key).Result()
	if err != nil {
		return nil, err
	}
	ttl, err := rdb.TTL(context.Background(), req.Key).Result()
	if err != nil {
		return nil, err
	}
	return &define.KeyValueReply{
		Value: v,
		TTL:   ttl,
		Type:  _type,
	}, nil
}

func CreateKeyValue(req *define.CreateKeyValueRequest) error {
	conn, err := utils.GetConnection(req.ConnectionId)
	if err != nil {
		return err
	}
	rdb := redis.NewClient(&redis.Options{
		Network:         "",
		Addr:            conn.Address + ":" + conn.Port,
		Username:        conn.Username,
		Password:        conn.Password,
		DB:              req.Db,
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
	if req.Type == "string" {
		err = rdb.Set(context.Background(), req.Key, "value", 0).Err()
	} else if req.Type == "hash" {
		err = rdb.HSet(context.Background(), req.Key, map[string]string{"key": "value"}).Err()
	} else if req.Type == "list" {
		err = rdb.RPush(context.Background(), req.Key, "value").Err()
	} else if req.Type == "set" {
		err = rdb.SAdd(context.Background(), req.Key, "value").Err()
	} else if req.Type == "zset" {
		err = rdb.ZAdd(context.Background(), req.Key, &redis.Z{
			Score:  0,
			Member: "value",
		}).Err()
	}
	return err
}

func DeleteKeyValue(req *define.KeyValueRequest) error {
	conn, err := utils.GetConnection(req.ConnectionId)
	if err != nil {
		return err
	}
	rdb := redis.NewClient(&redis.Options{
		Network:         "",
		Addr:            conn.Address + ":" + conn.Port,
		Username:        conn.Username,
		Password:        conn.Password,
		DB:              req.Db,
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
	_, err = rdb.Del(context.Background(), req.Key).Result()
	return err
}

// UpdateKeyValue 更新键值对数据
func UpdateKeyValue(req *define.UpdateKeyValueRequest) error {
	conn, err := utils.GetConnection(req.ConnectionId)
	if err != nil {
		return err
	}
	rdb := redis.NewClient(&redis.Options{
		Network:  "",
		Addr:     conn.Address + ":" + conn.Port,
		Username: conn.Username,
		Password: conn.Password,
		DB:       req.Db,
	})
	// type ==> string
	temp, _ := strconv.Atoi(req.TTL)
	tmep1 := time.Duration(temp)
	err = rdb.Set(context.Background(), req.Key, req.Value, tmep1).Err()
	return err
}
