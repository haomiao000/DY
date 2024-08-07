package api

import (
	"context"

	"github.com/haomiao000/DY/server/redis_svr/internal"
	pb "github.com/haomiao000/DY/server/redis_svr/pb/redis_svr"
)

func (r *RedisSvrImpl) Get(ctx context.Context, req *pb.GetReq) (*pb.GetRsp, error) {
	val, exist, err := internal.Get(req.GetKey())
	if err != nil {
		return nil, err
	}
	return &pb.GetRsp{Val: val, Exist: exist}, nil
}

func (r *RedisSvrImpl) BatchGet(ctx context.Context, req *pb.BatchGetReq) (*pb.BatchGetRsp, error) {
	m, err := internal.BatchGet(req.GetKeys())
	if err != nil {
		return nil, err
	}
	return &pb.BatchGetRsp{Vals: m}, nil
}

func (r *RedisSvrImpl) LRange(ctx context.Context , req *pb.LRangeRequest) (*pb.LRangeResponse , error) {
	vals , err := internal.LRange(req.Key , int(req.Start) , int(req.Stop))
	if err != nil {
		return nil, err
	}
	return &pb.LRangeResponse{Values: vals} , nil
}

func (r *RedisSvrImpl) SISMember(ctx context.Context , req *pb.SISMemberRequest) (*pb.SISMemberResponse , error) {
	exist , err := internal.SIsMember(req.Key , req.Value)
	if err != nil {
		return nil , err
	}
	return &pb.SISMemberResponse{Exists: exist} , nil
}

func (r *RedisSvrImpl) SMembers(ctx context.Context , req *pb.SMembersRequest) (*pb.SMembersResponse , error) {
	vals , err := internal.SMembers(req.Key)
	if err != nil {
		return nil , err
	}
	return &pb.SMembersResponse{Values: vals} , nil
}

func (r *RedisSvrImpl) SCard(ctx context.Context , req *pb.SCardRequest) (*pb.SCardResponse , error) {
	count , err := internal.SCard(req.Key)
	if err != nil {
		return nil , err
	}
	return &pb.SCardResponse{Count: int32(count)} , nil
}

