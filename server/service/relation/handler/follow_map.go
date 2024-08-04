package handler

import "fmt"

func (s *RelationServiceImpl)GetFollowMap(userID int64) (*map[int64]bool , error) {
	var mp map[int64]bool
	userIdList , err := s.MysqlManager.GetFollowUserIdList(userID); 
	if err != nil {
		fmt.Println("get follow user id list error")
		return nil , err
	}
	mp = make(map[int64]bool)
	for _ , o := range userIdList {
		mp[o] = true
	}
	return &mp , nil
}