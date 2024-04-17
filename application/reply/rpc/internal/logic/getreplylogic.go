package logic

import (
	"context"

	"gozero/application/reply/rpc/internal/svc"
	"gozero/application/reply/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetReplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetReplyLogic {
	return &GetReplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetReplyLogic) GetReply(in *service.GetReplyRequest) (*service.GetReplyResponse, error) {
	// todo: add your logic here and delete this line

	return &service.GetReplyResponse{}, nil
}
