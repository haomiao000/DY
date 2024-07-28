package api

import (
	"context"

	"github.com/haomiao000/DY/server/redis_svr/internal"
	pb "github.com/haomiao000/DY/server/redis_svr/pb/redis_svr"
)

func (r *RedisSvrImpl) Get(ctx context.Context, req *pb.GetReq) (*pb.GetRsp, error) {
	val, err := internal.Get(req.GetKey())
	if err != nil {
		return nil, err
	}
	return &pb.GetRsp{Val: val}, nil
}

func (r *RedisSvrImpl) BatchGet(ctx context.Context, req *pb.BatchGetReq) (*pb.BatchGetRsp, error) {
	m, err := internal.BatchGet(req.GetKeys())
	if err != nil {
		return nil, err
	}
	return &pb.BatchGetRsp{Vals: m}, nil
}
