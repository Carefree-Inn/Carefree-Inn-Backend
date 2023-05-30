package like

type Like struct {
	LikeId     uint32 `json:"like_id"`
	Account    string `json:"account"`
	PostId     uint32 `json:"post_id"`
	CreateTime string `json:"create_time"`
}
