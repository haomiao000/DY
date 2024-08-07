package discovery

import (
	"context"
	"fmt"
	"strings"

	etcd "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
)

// EtcdResolver 解析器接口实现结构体，客户端使用方法实例：
// con, err := grpc.NewClient("etcd://redis_svr", grpc.WithResolvers(discovery.GetResolver()),
// grpc.WithTransportCredentials(insecure.NewCredentials()))
type EtcdResolver struct {
	cli *etcd.Client
}

func (r *EtcdResolver) ResolveNow(rn resolver.ResolveNowOptions) {}
func (r *EtcdResolver) Close()                                   {}
func (r *EtcdResolver) Scheme() string {
	return "etcd"
}
func (r *EtcdResolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	resp, err := r.cli.Get(context.Background(), target.Endpoint(), etcd.WithPrefix(), etcd.WithSerializable())
	if err != nil {
		fmt.Printf("Failed to get key-value from etcd: %v", err)
		return nil, err
	}
	addrs := map[string]bool{}
	for _, kv := range resp.Kvs {
		addrs[string(kv.Value)] = true
	}
	cc.UpdateState(resolver.State{Addresses: getUniqueAddress(addrs)})

	go func() {
		rch := r.cli.Watch(context.Background(), target.Endpoint(), etcd.WithPrefix(), etcd.WithPrevKV())
		for n := range rch {
			for _, ev := range n.Events {
				switch ev.Type {
				case etcd.EventTypePut:
					addrs[string(ev.Kv.Value)] = true
					fmt.Println("update address:", addrs)
					cc.UpdateState(resolver.State{Addresses: getUniqueAddress(addrs)})
				case etcd.EventTypeDelete:
					fmt.Println("etcd.EventTypeDelete")
					fmt.Printf("delete %+v\n", ev.Kv)
					strs := strings.Split(string(ev.Kv.Key), "/")
					if len(strs) != 3 {
						fmt.Printf("kv invalid: %+v", ev.Kv)
						continue
					}
					delete(addrs, strs[2])
					fmt.Println("update address:", addrs)
					cc.UpdateState(resolver.State{Addresses: getUniqueAddress(addrs)})
				}
			}
		}
	}()
	return r, nil
}

func getUniqueAddress(m map[string]bool) []resolver.Address {
	res := make([]resolver.Address, len(m))
	for k := range m {
		res = append(res, resolver.Address{Addr: k})
	}
	return res
}
