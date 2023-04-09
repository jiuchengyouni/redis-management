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

type KeyValue struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type KeyValueRequest struct {
	ConnectionId string `json:"connection_id"`
	Db           int    `json:"db"`
	Key          string `json:"key"`
}

type KeyValueReply struct {
	Value interface{} `json:"value"`
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

type Email struct {
	Address string `json:"address"`
}

type HashFieldDeleteRequest struct {
	KeyValueRequest
	Field []string `json:"field"`
}

type HashAddOrUpdateFieldRequest struct {
	KeyValueRequest
	Field string `json:"field"`
	Value string `json:"value"`
}

type ListValueRequest struct {
	KeyValueRequest
	Value string `json:"value"`
}

type SetValueRequest struct {
	KeyValueRequest
	Value string `json:"value"`
}

type ZSetValueRequest struct {
	KeyValueRequest
	Score  float64     `json:"score"`
	Member interface{} `json:"member"`
}

var (
	// DefaultKeyLen 键列表的默认查询长度
	DefaultKeyLen int64 = 100
	// MaxKeyLen 键列表的最大查询长度
	MaxKeyLen int64 = 2000
	// MaxHashLen hash列表的最大查询长度
	MaxHashLen int64 = 200
	// MaxListLen list列表的最大查询长度
	MaxListLen int64 = 200
	// MaxSetLen set的最大查询长度
	MaxSetLen int64 = 200
	// MaxZSetLen zset的最大查询长度
	MaxZSetLen int64 = 200
)
