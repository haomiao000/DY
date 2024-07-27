package api

import (
	"context"

	"github.com/haomiao000/DY/server/redis_svr/internal"
	pb "github.com/haomiao000/DY/server/redis_svr/pb/redis_svr"
)

func (r *RedisSvrImpl) Set(ctx context.Context, req *pb.SetReq) (*pb.SetRsp, error) {
	err := internal.Set(req.GetKey(), req.GetVal())
	if err != nil {
		return nil, err
	}
	return &pb.SetRsp{}, nil
}

func (r *RedisSvrImpl) BatchSet(ctx context.Context, req *pb.BatchSetReq) (*pb.BatchSetRsp, error) {
	err := internal.BatchSet(req.GetKv())
	if err != nil {
		return nil, err
	}
	return &pb.BatchSetRsp{}, nil
}

// SetWithExpire 带过期时间的set
func (r *RedisSvrImpl) SetWithExpire(ctx context.Context, req *pb.SetWithExpireReq) (*pb.SetWithExpireRsp, error) {
	err := internal.SetWithExpire(req.GetKey(), req.GetVal(), int(req.GetExpire()))
	if err != nil {
		return nil, err
	}
	return &pb.SetWithExpireRsp{}, nil
}

// SetIfNotExist 若key不存在则设置，并返回true
func (r *RedisSvrImpl) SetIfNotExist(ctx context.Context, req *pb.SetIfNotExistReq) (*pb.SetIfNotExistRsp, error) {
	ok, err := internal.SetIfNotExist(req.GetKey(), req.GetVal())
	if err != nil {
		return nil, err
	}
	return &pb.SetIfNotExistRsp{Ok: ok}, nil
}
