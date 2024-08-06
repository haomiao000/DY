package model

//@Register
type UserRegisterResponse struct {
	BaseResp *Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}
//@Login
type UserLoginResponse struct {
	BaseResp *Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}
//@User
type UserResponse struct {
	BaseResp *Response
	User *User `json:"user"`
}
//@VideoList
type VideoListResponse struct {
	BaseResp  *Response
	VideoList []*Video `json:"video_list"`
}
//FavoriteAction
type FavoriteActionResponse struct {
	// Status code, 0-success, other values-failure
	StatusCode int32 `json:"status_code"`
	// Return status description
	StatusMsg string `json:"status_msg"`
}
//@FavoriteList
type FavoriteListResponse struct {
	StatusCode int32 `json:"status_code"`
	StatusMsg string `json:"status_msg"`
	VideoList []*Video `json:"video_list"`
}
//@CommentAction
type CommentActionResponse struct {
	BaseResp *Response `json:"base_resp"`
	Comment *Comment `json:"comment,omitempty"`
}
//@CommentList
type CommentListResponse struct {
	BaseResp *Response 
	CommentList []*Comment `json:"comment_list,omitempty"`
}
//@RelationAction
type RelationActionResponse struct {
	BaseResp *Response 
}
//@UserList(@FollowingList@FollowerList@FriendList)
type UserListResponse struct {
	BaseResp *Response 
	UserList []*User `json:"user_list"`
}