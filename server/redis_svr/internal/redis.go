package internal

// 具体逻辑

import (
	"errors"
	"fmt"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/haomiao000/DY/server/redis_svr/entity"
)

type RedisCon struct {
	redisCon redigo.Conn
}

var (
	ErrKeyCrash = errors.New("key crash")
	ErrNotExist = errors.New("not exist")
)

var redis *RedisCon

func Init() error {
	c, err := redigo.Dial("tcp", "localhost:6379")
	if err != nil {
		return err
	}
	redis = &RedisCon{c}
	return nil
}
func getConn() (*RedisCon, error) {
	return redis, nil
}

func Do(cmd string, args ...interface{}) (interface{}, error) {
	c, err := getConn()
	if err != nil {
		return nil, err
	}
	return c.do(cmd, args...)
}

func (r *RedisCon) do(cmd string, args ...interface{}) (interface{}, error) {
	return r.redisCon.Do(cmd, args...)
}

func Get(key string) (string, bool, error) {
	res, err := redigo.String(Do("get", key))
	if err != nil && err != redigo.ErrNil {
		return "", false, fmt.Errorf("get value error: %v, key: %s", err, key)
	}
	if err == redigo.ErrNil {
		return "", false, nil
	}

	return res, true, nil
}

func Set(key, val string) error {
	ret, err := redigo.String(Do("set", key, val))
	if err != nil {
		return fmt.Errorf("set key: %s error: %v", key, err)
	}
	if ret != "OK" {
		return errors.New(ret)
	}
	return nil
}

func SetWithExpire(key, val string, expire int) error {
	ret, err := redigo.String(Do("setex", key, expire, val))
	if err != nil {
		return fmt.Errorf("setex key: %s error: %v", key, err)
	}
	if ret != "OK" {
		return errors.New(ret)
	}
	return nil
}

func BatchGet(keys []string) (map[string]string, error) {
	vals, err := redigo.Strings(Do("mget", redigo.Args{}.AddFlat(keys)...))
	if err != nil {
		return nil, err
	}
	res := make(map[string]string, len(vals))
	for i, val := range vals {
		key := keys[i]
		// 不存在
		if val == "" {
			continue
		}
		res[key] = val
	}
	return res, nil
}

func BatchSet(kv map[string]string) error {
	args := redigo.Args{}
	for k, v := range kv {
		args = args.Add(k, v)
	}
	ret, err := redigo.String(Do("mset", args...))
	if err != nil {
		return fmt.Errorf("set key: %s error: %v", kv, err)
	}
	if ret != "OK" {
		return errors.New(ret)
	}
	return nil
}

func SetIfNotExist(key, val string) (bool, error) {
	ok, err := redigo.Bool(Do("EVAL", entity.GetAndSetLua, 1, key, val))
	if err != nil && err != redigo.ErrNil {
		return false, err
	}
	if !ok {
		return false, nil
	}
	return true, nil
}
