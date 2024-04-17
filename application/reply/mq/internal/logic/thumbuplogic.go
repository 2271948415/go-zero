package logic

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gozero/application/like/mq/internal/model"
	"gozero/application/like/rpc/types"
	"time"

	"gozero/application/like/mq/internal/svc"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
)

type ThumbupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewThumbupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThumbupLogic {
	return &ThumbupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ThumbupLogic) Consume(_, val string) error {
	logx.Infof("Consume msg val: %s", val)
	var msg *types.ThumbupMsg
	err := json.Unmarshal([]byte(val), &msg)
	if err != nil {
		logx.Errorf("Consume val: %s error: %v", val, err)
		return err
	}
	result, err := l.svcCtx.LikeModel.FindOne(l.ctx, msg.BizId, msg.ObjId, msg.UserId)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			logx.Errorf("Consume find record fail: %v", err)
			return err
		}
		record := &model.Like{
			Id:         uuid.NewString(),
			BizId:      msg.BizId,
			ObjId:      msg.ObjId,
			UserId:     msg.UserId,
			LikeType:   msg.LikeType,
			CreateTime: time.Now().Format("2006-01-02 15:04:05"),
			UpdateTime: time.Now().Format("2006-01-02 15:04:05"),
		}
		err = l.svcCtx.LikeModel.Insert(l.ctx, record)
		if err != nil {
			logx.Errorf("Error inserting record fail: %v", err)
			return err
		}

	}
	result.LikeType = msg.LikeType
	err = l.svcCtx.LikeModel.Update(l.ctx, result)
	if err != nil {
		logx.Errorf("Consume update record fail: %v", err)
		return err
	}
	return nil
}

func Consumers(ctx context.Context, svcCtx *svc.ServiceContext) []service.Service {
	return []service.Service{
		kq.MustNewQueue(svcCtx.Config.KqConsumerConf, NewThumbupLogic(ctx, svcCtx)),
	}
}
