package comment

type UserInfo struct {
	Account  string `json:"account"`
	Nickname string `json:"nickname"`
	Sex      int8   `json:"sex"`
	Avatar   string `json:"avatar"`
}

type Comment struct {
	CommentId  uint32 `json:"comment_id"`
	PostId     uint32 `json:"post_id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
	
	FromUserAccount *UserInfo `json:"from_user_account"`
	ToUserAccount   *UserInfo `json:"to_user_account"`
	
	IsTop        bool       `json:"is_top"`
	TopCommentId uint32     `json:"top_comment_id"`
	ChildComment []*Comment `json:"child_comment"`
}
