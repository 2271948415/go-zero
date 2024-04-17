package logic

import (
	"context"

	"gozero/application/reply/rpc/internal/svc"
	"gozero/application/reply/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteReplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteReplyLogic {
	return &DeleteReplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteReplyLogic) DeleteReply(in *service.DeleteReplyRequest) (*service.Response, error) {
	// todo: add your logic here and delete this line

	return &service.Response{}, nil
}
