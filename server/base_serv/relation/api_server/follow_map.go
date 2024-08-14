package api_server

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	redis "github.com/haomiao000/DY/comm/redis"
)
func (s *RelationServiceImpl) GetFollowMap(ctx context.Context , userID int64) (*map[int64]bool, error) {
	var mp map[int64]bool
	var userIdList []int64
	userIdListStr , err := redis.SMembers(ctx , strconv.FormatInt(userID, 10))
	if err != nil {
		userIdList, err = s.MysqlManager.GetFollowUserIdList(userID)
		if err != nil {
			fmt.Println("get follow user id list error")
			return nil, err
		}
	}else {
		userIdList = make([]int64, len(userIdListStr))
		for i , str := range userIdListStr {
			userId, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				return nil , errors.New("failed to parse user ID")
			}
			userIdList[i] = userId
		}
	}
	mp = make(map[int64]bool)
	for _, o := range userIdList {
		mp[o] = true
	}
	return &mp, nil
}
