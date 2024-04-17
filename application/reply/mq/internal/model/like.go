package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Like struct {
	Id         string `gorm:"id" json:"id"`
	BizId      string `gorm:"biz_id" json:"biz_id"`
	ObjId      string `gorm:"obj_id" json:"obj_id"`
	UserId     string `gorm:"user_id" json:"user_id"`
	LikeType   int32  `gorm:"like_type" json:"like_type"`
	CreateTime string `gorm:"create_time" json:"create_time"`
	UpdateTime string `gorm:"update_time" json:"update_time"`
}

func (m *Like) TableName() string {
	return "like_record"
}

type LikeModel struct {
	db *gorm.DB
}

func NewLikeModel(db *gorm.DB) *LikeModel {
	return &LikeModel{
		db: db,
	}
}

func (m *LikeModel) FindOne(ctx context.Context, bizId, objId, uid string) (*Like, error) {
	var result Like
	err := m.db.WithContext(ctx).Where("biz_id = ?", bizId).Where("obj_id = ?", objId).
		Where("user_id = ?", uid).
		First(&result).Error
	return &result, err
}

func (m *LikeModel) Insert(ctx context.Context, data *Like) error {
	return m.db.WithContext(ctx).Create(data).Error
}

func (m *LikeModel) Update(ctx context.Context, data *Like) error {
	return m.db.WithContext(ctx).Model(&Like{}).
		Where("id = ?", data.Id).
		Updates(map[string]interface{}{
			"like_type":   data.LikeType,
			"update_time": time.Now().Format("2006-01-02 15:04:05"),
		}).Error
}

func (m *LikeModel) UpdateFields(ctx context.Context, id string, values map[string]interface{}) error {
	return m.db.WithContext(ctx).Model(&Like{}).Where("id = ?", id).Updates(values).Error
}
