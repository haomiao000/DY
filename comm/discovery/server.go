package discovery

import (
	"fmt"
	"time"

	etcd "go.etcd.io/etcd/client/v3"
)

// Register 服务端调用 service为服务名，如redis_svr，addr为ip地址，如43.138.235.2:50050
func Register(service, addr string) error {
	key := fmt.Sprintf("/%s/%s", service, addr)
	go func() {
		ticker := time.NewTicker(time.Duration(interval) * time.Second)
		for {
			leaseID, err := r.GetLease(interval)
			if err != nil {
				fmt.Printf("get lease error: %v", err)
				continue
			}
			err = r.PutKeyValue(key, addr, etcd.WithLease(leaseID))
			if err != nil {
				fmt.Printf("put key value error: %v", err)
			}
			<-ticker.C
		}
	}()
	return nil
}
