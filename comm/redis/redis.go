package main

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	pb "github.com/haomiao000/DY/server/redis_svr/pb/redis_svr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

var redisCli pb.RedisSvrClient

// 初始化 gRPC 客户端连接
func init() {
	con, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	redisCli = pb.NewRedisSvrClient(con)
}

// 设置 Redis 中的键值对
func Set(ctx context.Context, key, value string) error {
	_, err := redisCli.Set(ctx, &pb.SetReq{Key: key, Val: value})
	return err
}

// 批量设置 Redis 中的键值对
func BatchSet(ctx context.Context, keys map[string]string) error {
	_, err := redisCli.BatchSet(ctx, &pb.BatchSetReq{Kv: keys})
	return err
}

// 设置 Protobuf 序列化后的值
func SetProto(ctx context.Context, key string, value proto.Message) error {
	b, err := proto.Marshal(value)
	if err != nil {
		return err
	}
	return Set(ctx, key, string(b))
}

// 设置 JSON 序列化后的值
func SetJson(ctx context.Context, key string, value any) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return Set(ctx, key, string(b))
}

// 如果键不存在则设置键值对
func SetIfNotExist(ctx context.Context, key, value string) (bool, error) {
	rsp, err := redisCli.SetIfNotExist(ctx, &pb.SetIfNotExistReq{Key: key, Val: value})
	if err != nil {
		return false, err
	}
	return rsp.GetOk(), nil
}

// 设置带过期时间的键值对
func SetWithExpire(ctx context.Context, key, value string, expire int32) error {
	_, err := redisCli.SetWithExpire(ctx, &pb.SetWithExpireReq{Key: key, Val: value, Expire: expire})
	return err
}

// 批量设置 Protobuf 序列化后的值
func BatchSetProto(ctx context.Context, kv map[string]proto.Message) error {
	m := make(map[string]string, len(kv))
	for k, v := range kv {
		b, err := proto.Marshal(v)
		if err != nil {
			return fmt.Errorf("key: %s proto marshal error: %v", k, err)
		}
		m[k] = string(b)
	}
	return BatchSet(ctx, m)
}

// 批量设置 JSON 序列化后的值
func BatchSetJson(ctx context.Context, kv map[string]any) error {
	m := make(map[string]string, len(kv))
	for k, v := range kv {
		b, err := json.Marshal(v)
		if err != nil {
			return fmt.Errorf("key: %s json marshal error: %v", k, err)
		}
		m[k] = string(b)
	}
	return BatchSet(ctx, m)
}

// 如果键不存在则设置 Protobuf 序列化后的值
func SetIfNotExistProto(ctx context.Context, key string, value proto.Message) (bool, error) {
	b, err := proto.Marshal(value)
	if err != nil {
		return false, err
	}
	return SetIfNotExist(ctx, key, string(b))
}

// 如果键不存在则设置 JSON 序列化后的值
func SetIfNotExistJson(ctx context.Context, key string, value any) (bool, error) {
	b, err := json.Marshal(value)
	if err != nil {
		return false, err
	}
	return SetIfNotExist(ctx, key, string(b))
}

// 设置带过期时间的 Protobuf 序列化后的值
func SetWithExpireProto(ctx context.Context, key string, value proto.Message, expire int32) error {
	b, err := proto.Marshal(value)
	if err != nil {
		return err
	}
	return SetWithExpire(ctx, key, string(b), expire)
}

// 设置带过期时间的 JSON 序列化后的值
func SetWithExpireJson(ctx context.Context, key string, value any, expire int32) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return SetWithExpire(ctx, key, string(b), expire)
}

// 获取 Redis 中的键值对
func Get(ctx context.Context, key string) (string, bool, error) {
	rsp, err := redisCli.Get(ctx, &pb.GetReq{Key: key})
	if err != nil {
		return "", false, err
	}
	return rsp.GetVal(), rsp.GetExist(), nil
}

// 获取并反序列化 Protobuf 值
func GetProto(ctx context.Context, key string, msg proto.Message) (bool, error) {
	val, exist, err := Get(ctx, key)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, nil
	}
	err = proto.Unmarshal([]byte(val), msg)
	if err != nil {
		return false, err
	}
	return exist, nil
}

// 获取并反序列化 JSON 值
func GetJson(ctx context.Context, key string, msg any) (bool, error) {
	val, exist, err := Get(ctx, key)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, nil
	}
	err = json.Unmarshal([]byte(val), &msg)
	if err != nil {
		return false, err
	}
	return exist, nil
}

// 批量获取 Redis 中的键值对
func BatchGet(ctx context.Context, keys []string) (map[string]string, error) {
	rsp, err := redisCli.BatchGet(ctx, &pb.BatchGetReq{Keys: keys})
	if err != nil {
		return nil, err
	}
	return rsp.GetVals(), nil
}

// 批量获取并反序列化 Protobuf 值
func BatchGetProto(ctx context.Context, keys []string, msg proto.Message) (map[string]proto.Message, error) {
	m, err := BatchGet(ctx, keys)
	if err != nil {
		return nil, err
	}
	msgs := make(map[string]proto.Message, len(keys))
	errStr := ""
	for k, v := range m {
		e := proto.Unmarshal([]byte(v), msg)
		if e != nil {
			errStr += fmt.Sprintf("key %s proto unmarshal error: %v |", k, err)
			continue
		}
		msgs[k] = proto.Clone(msg) // 深拷贝
	}
	if errStr != "" {
		return msgs, fmt.Errorf(errStr)
	}
	return msgs, nil
}

// 批量获取并反序列化 JSON 值
func BatchGetJson(ctx context.Context, keys []string, msg any) (map[string]any, error) {
	m, err := BatchGet(ctx, keys)
	if err != nil {
		return nil, err
	}
	msgs := make(map[string]any, len(keys))
	errStr := ""
	for k, v := range m {
		// 深拷贝
		newMsg := reflect.New(reflect.TypeOf(msg).Elem()).Interface()
		e := json.Unmarshal([]byte(v), &newMsg)
		if e != nil {
			errStr += fmt.Sprintf("key %s json unmarshal error: %v |", k, err)
			continue
		}
		msgs[k] = newMsg
	}
	if errStr != "" {
		return msgs, fmt.Errorf(errStr)
	}
	return msgs, nil
}