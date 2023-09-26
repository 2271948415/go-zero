package logic

import (
	"context"
	"github.com/google/uuid"
	"gozero/application/user/code"
	"gozero/application/user/internal/model"
	"time"

	"gozero/application/user/internal/svc"
	"gozero/application/user/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *service.RegisterRequest) (*service.RegisterResponse, error) {
	if len(in.Username) == 0 {
		return nil, code.RegisterNameEmpty
	}
	userId := uuid.NewString()
	_, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Id:         userId,
		Username:   in.Username,
		Mobile:     in.Mobile,
		Avatar:     in.Avatar,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	if err != nil {
		logx.Errorf("Register req: %v error: %v", in, err)
		return nil, err
	}

	return &service.RegisterResponse{UserId: userId}, nil
}
