package api

import (
	"context"
	"errors"
	"fmt"

	"github.com/haomiao000/DY/server/redis_svr/internal"
	pb "github.com/haomiao000/DY/server/redis_svr/pb/redis_svr"
)

func (r *RedisSvrImpl) Set(ctx context.Context, req *pb.SetReq) (*pb.SetRsp, error) {
	fmt.Println("-------")
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

func (r *RedisSvrImpl) Delete(ctx context.Context , req *pb.DeleteReq) (*pb.DeleteRsp , error) {
	exist , err := internal.Delete(req.GetKey())
	if err != nil {
		return nil , err
	}
	return &pb.DeleteRsp{Exist: exist} , nil
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

func (r *RedisSvrImpl) LPush(ctx context.Context , req *pb.LPushRequest) (*pb.ListResponse , error) {
	err := internal.LPush(req.Key , int(req.ExpireTime) , req.Values...)
	if err != nil {
		return nil , err
	}
	return &pb.ListResponse{} , nil
}

func (r *RedisSvrImpl) RPush(ctx context.Context , req *pb.RPushRequest) (*pb.ListResponse , error) {
	err := internal.RPush(req.Key , int(req.ExpireTime) , req.Values...)
	if err != nil {
		return nil , err
	}
	return &pb.ListResponse{} , nil
}

func (r *RedisSvrImpl) LPop(ctx context.Context , req *pb.LPopRequest) (*pb.PopResponse , error) {
	err := internal.LPop(req.Key)
	if err != nil {
		return nil , err
	}
	return &pb.PopResponse{} , nil
}

func (r *RedisSvrImpl) RPop(ctx context.Context , req *pb.RPopRequest) (*pb.PopResponse , error) {
	err := internal.RPop(req.Key)
	if err != nil {
		return nil , err
	}
	return &pb.PopResponse{} , nil
}

func (r *RedisSvrImpl) SAdd(ctx context.Context , req *pb.SAddRequest) (*pb.SAddResponse , error) {
	err := internal.SAdd(req.Key , int(req.Expire) , req.Value)
	if err != nil {
		return nil , err
	}
	return &pb.SAddResponse{} , nil
}

func (r *RedisSvrImpl) SRem(ctx context.Context, req *pb.SRemRequest) (*pb.SRemResponse, error) {
	count, err := internal.SRem(req.Key, req.Value)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return &pb.SRemResponse{}, errors.New("item not exist in redis")
	}
	return &pb.SRemResponse{}, nil
}