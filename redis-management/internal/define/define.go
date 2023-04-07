package define

import "time"

var ConfigName = "redis-management.conf"

type Connection struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Config struct {
	Connections []*Connection `json:"connections"`
}

type DbItem struct {
	Key    string `json:"key"`
	Number int    `json:"number"`
}

type KeyListRequest struct {
	ConnectionId string `json:"connection_id"`
	Db           int    `json:"db"`
	Keyword      string `json:"keyword"`
}

type KeyValueRequest struct {
	ConnectionId string `json:"connection_id"`
	Db           int    `json:"db"`
	Key          string `json:"key"`
}

type KeyValueReply struct {
	Value string `json:"value"`
	//超时时间
	TTL  time.Duration `json:"ttl"`
	Type string        `json:"type"`
}
type CreateKeyValueRequest struct {
	KeyValueRequest
	Type string `json:"type"`
}

type UpdateKeyValueRequest struct {
	KeyValueRequest
	TTL   string `json:"ttl"`
	Value string `json:"value"`
}
