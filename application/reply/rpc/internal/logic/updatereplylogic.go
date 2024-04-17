package logic

import (
	"context"

	"gozero/application/reply/rpc/internal/svc"
	"gozero/application/reply/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateReplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateReplyLogic {
	return &UpdateReplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateReplyLogic) UpdateReply(in *service.UpdateReplyRequest) (*service.Response, error) {
	// todo: add your logic here and delete this line

	return &service.Response{}, nil
}
