package logic

import (
	"context"

	"gozero/application/reply/rpc/internal/svc"
	"gozero/application/reply/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostReplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostReplyLogic {
	return &PostReplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostReplyLogic) PostReply(in *service.PostReplyRequest) (*service.Response, error) {
	// todo: add your logic here and delete this line

	return &service.Response{}, nil
}
