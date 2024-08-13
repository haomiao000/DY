package discovery

import (
	"context"
	"fmt"
	"time"

	etcd "go.etcd.io/etcd/client/v3"
)

var (
	endPoints = []string{"localhost:2379"}
)

var (
	dialTimeout  int64
	readTimeout  int64
	writeTimeout int64

	interval int64
)

var r *EtcdResolver

func Init() error {
	dialTimeout = 5
	readTimeout = 5
	writeTimeout = 5

	interval = 5
	etcdClient, err := etcd.New(etcd.Config{
		Endpoints:   endPoints,
		DialTimeout: time.Duration(dialTimeout) * time.Second,
	})
	if err != nil {
		return err
	}
	r = &EtcdResolver{cli: etcdClient}
	return nil
}

func GetResolver() *EtcdResolver {
	return r
}

// GetWithPrefix 根据前缀查询
func (r *EtcdResolver) GetWithPrefix(prefix string) ([][]byte, error) {
	result := make([][]byte, 0)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(readTimeout)*time.Second)
	defer cancel()
	rsp, err := r.cli.Get(ctx, prefix, etcd.WithPrefix(), etcd.WithPrevKV())

	if err != nil {
		return nil, err
	}
	if rsp == nil {
		return nil, fmt.Errorf("etcd rsp empty, prefix: %s", prefix)
	}
	for _, v := range rsp.Kvs {
		result = append(result, v.Value)
	}
	return result, nil
}

// GetLease 获取一个租约
func (r *EtcdResolver) GetLease(ttl int64) (etcd.LeaseID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(readTimeout)*time.Second)
	defer cancel()
	if r.cli.Lease == nil {
		return 0, fmt.Errorf("get lease error")
	}
	rsp, err := r.cli.Grant(ctx, ttl)
	if err != nil {
		return 0, err
	}
	return etcd.LeaseID(rsp.ID), nil
}

// PutKeyValue 设置键值对
func (r *EtcdResolver) PutKeyValue(key, value string, opts ...etcd.OpOption) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(writeTimeout)*time.Second)
	defer cancel()
	_, err := r.cli.Put(ctx, key, value, opts...)
	return err
}
