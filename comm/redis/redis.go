package redis

import (
	"context"
	"encoding/json"
	"fmt"

	pb "github.com/haomiao000/DY/server/redis_svr/pb/redis_svr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

var redisCli pb.RedisSvrClient

func init() {
	con, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	redisCli = pb.NewRedisSvrClient(con)
}

func Set(ctx context.Context, key, value string) error {
	_, err := redisCli.Set(ctx, &pb.SetReq{Key: key, Val: value})
	return err
}

func BatchSet(ctx context.Context, keys map[string]string) error {
	_, err := redisCli.BatchSet(ctx, &pb.BatchSetReq{Kv: keys})
	return err
}

func SetProto(ctx context.Context, key string, value proto.Message) error {
	b, err := proto.Marshal(value)
	if err != nil {
		return err
	}
	return Set(ctx, key, string(b))
}

func SetJson(ctx context.Context, key string, value any) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return Set(ctx, key, string(b))
}

func SetIfNotExist(ctx context.Context, key, value string) (bool, error) {
	rsp, err := redisCli.SetIfNotExist(ctx, &pb.SetIfNotExistReq{Key: key, Val: value})
	if err != nil {
		return false, err
	}
	return rsp.GetOk(), nil
}

func SetWithExpire(ctx context.Context, key, value string, expire int32) error {
	_, err := redisCli.SetWithExpire(ctx, &pb.SetWithExpireReq{Key: key, Val: value, Expire: expire})
	return err
}

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

func SetIfNotExistProto(ctx context.Context, key string, value proto.Message) (bool, error) {
	b, err := proto.Marshal(value)
	if err != nil {
		return false, err
	}
	return SetIfNotExist(ctx, key, string(b))
}

func SetIfNotExistJson(ctx context.Context, key string, value any) (bool, error) {
	b, err := json.Marshal(value)
	if err != nil {
		return false, err
	}
	return SetIfNotExist(ctx, key, string(b))
}

func SetWithExpireProto(ctx context.Context, key string, value proto.Message, expire int32) error {
	b, err := proto.Marshal(value)
	if err != nil {
		return err
	}
	return SetWithExpire(ctx, key, string(b), expire)
}

func SetWithExpireJson(ctx context.Context, key string, value any, expire int32) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return SetWithExpire(ctx, key, string(b), expire)
}

func Get(ctx context.Context, key string) (string, error) {
	rsp, err := redisCli.Get(ctx, &pb.GetReq{Key: key})
	if err != nil {
		return "", err
	}
	return rsp.GetVal(), nil
}

func BatchGet(ctx context.Context, keys []string) (map[string]string, error) {
	rsp, err := redisCli.BatchGet(ctx, &pb.BatchGetReq{Keys: keys})
	if err != nil {
		return nil, err
	}
	return rsp.GetVals(), nil
}

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
		msgs[k] = msg
	}
	if errStr != "" {
		return msgs, fmt.Errorf(errStr)
	}
	return msgs, nil
}

func BatchGetJson(ctx context.Context, keys []string, msg any) (map[string]any, error) {
	m, err := BatchGet(ctx, keys)
	if err != nil {
		return nil, err
	}
	msgs := make(map[string]any, len(keys))
	errStr := ""
	for k, v := range m {
		e := json.Unmarshal([]byte(v), msg)
		if e != nil {
			errStr += fmt.Sprintf("key %s json unmarshal error: %v |", k, err)
			continue
		}
		msgs[k] = msg
	}
	if errStr != "" {
		return msgs, fmt.Errorf(errStr)
	}
	return msgs, nil
}
