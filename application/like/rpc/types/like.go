package types

type ThumbupMsg struct {
	BizId    string ` json:"bizId,omitempty"`    // 业务id
	ObjId    string ` json:"objId,omitempty"`    // 点赞对象id
	UserId   string ` json:"userId,omitempty"`   // 用户id
	LikeType int32  ` json:"likeType,omitempty"` // 类型
}
