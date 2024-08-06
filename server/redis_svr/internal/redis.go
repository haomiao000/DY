package internal

// 具体逻辑

import (
	"errors"
	"fmt"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/haomiao000/DY/server/redis_svr/entity"
)

// RedisCon 封装了 Redis 连接对象
type RedisCon struct {
	redisCon redigo.Conn
}

// 定义可能的错误
var (
	ErrKeyCrash = errors.New("key crash") // key 冲突错误
	ErrNotExist = errors.New("not exist") // key 不存在错误
)

// redis 是 RedisCon 的全局实例
var redis *RedisCon

// Init 初始化 Redis 连接
func Init() error {
	c, err := redigo.Dial("tcp", "localhost:6379") // 连接 Redis 服务器
	if err != nil {
		return err
	}
	redis = &RedisCon{c} // 将连接封装到 RedisCon 结构体中
	return nil
}

// getConn 获取 Redis 连接
func getConn() (*RedisCon, error) {
	return redis, nil
}

// Do 执行 Redis 命令
func Do(cmd string, args ...interface{}) (interface{}, error) {
	c, err := getConn()
	if err != nil {
		return nil, err
	}
	return c.do(cmd, args...)
}

// (r *RedisCon) do 执行具体的 Redis 命令
func (r *RedisCon) do(cmd string, args ...interface{}) (interface{}, error) {
	return r.redisCon.Do(cmd, args...)
}

// Get 获取指定 key 的值
func Get(key string) (string, bool, error) {
	res, err := redigo.String(Do("get", key)) // 执行 GET 命令
	if err != nil && err != redigo.ErrNil {
		return "", false, fmt.Errorf("get value error: %v, key: %s", err, key)
	}
	if err == redigo.ErrNil { // key 不存在
		return "", false, nil
	}

	return res, true, nil
}

// Set 设置指定 key 的值
func Set(key, val string) error {
	ret, err := redigo.String(Do("set", key, val)) // 执行 SET 命令
	if err != nil {
		return fmt.Errorf("set key: %s error: %v", key, err)
	}
	if ret != "OK" { // 如果返回值不是 "OK"，则表示设置失败
		return errors.New(ret)
	}
	return nil
}

// SetWithExpire 设置带过期时间的 key
func SetWithExpire(key, val string, expire int) error {
	ret, err := redigo.String(Do("setex", key, expire, val)) // 执行 SETEX 命令
	if err != nil {
		return fmt.Errorf("setex key: %s error: %v", key, err)
	}
	if ret != "OK" {
		return errors.New(ret)
	}
	return nil
}

// BatchGet 批量获取指定 keys 的值
func BatchGet(keys []string) (map[string]string, error) {
	vals, err := redigo.Strings(Do("mget", redigo.Args{}.AddFlat(keys)...)) // 执行 MGET 命令
	if err != nil {
		return nil, err
	}
	res := make(map[string]string, len(vals))
	for i, val := range vals {
		key := keys[i]
		if val == "" { // 如果值为空，则表示 key 不存在
			continue
		}
		res[key] = val
	}
	return res, nil
}

// BatchSet 批量设置指定 keys 的值
func BatchSet(kv map[string]string) error {
	args := redigo.Args{}
	for k, v := range kv {
		args = args.Add(k, v)
	}
	ret, err := redigo.String(Do("mset", args...)) // 执行 MSET 命令
	if err != nil {
		return fmt.Errorf("set key: %s error: %v", kv, err)
	}
	if ret != "OK" {
		return errors.New(ret)
	}
	return nil
}

// SetIfNotExist 仅在 key 不存在时设置值
func SetIfNotExist(key, val string) (bool, error) {
	ok, err := redigo.Bool(Do("EVAL", entity.GetAndSetLua, 1, key, val)) // 执行 Lua 脚本实现的 SETNX 命令
	if err != nil && err != redigo.ErrNil {
		return false, err
	}
	if !ok { // key 已存在
		return false, nil
	}
	return true, nil
}