package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type LikeCount struct {
	Id         string    `gorm:"id" json:"id"`
	BizId      string    `gorm:"biz_id" json:"biz_id"`
	ObjId      string    `gorm:"obj_id" json:"obj_id"`
	LikeNum    int       `gorm:"like_num" json:"like_num"`
	DisLikeNum int       `gorm:"dislike_num" json:"dislike_num"`
	CreateTime time.Time `gorm:"create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"update_time" json:"update_time"`
}

func (m *LikeCount) TableName() string {
	return "like_count"
}

type LikeCountModel struct {
	db *gorm.DB
}

func NewLikeCountModel(db *gorm.DB) *LikeCountModel {
	return &LikeCountModel{
		db: db,
	}
}

func (m *LikeCountModel) FindOne(ctx context.Context, id string) (*LikeCount, error) {
	var result LikeCount
	err := m.db.WithContext(ctx).Where("id = ?", id).First(&result).Error
	return &result, err
}

func (m *LikeCountModel) Insert(ctx context.Context, data *LikeCount) error {
	return m.db.WithContext(ctx).Create(data).Error
}

func (m *LikeCountModel) Update(ctx context.Context, data *LikeCount) error {
	return m.db.WithContext(ctx).Save(data).Error
}

func (m *LikeCountModel) UpdateFields(ctx context.Context, id string, values map[string]interface{}) error {
	return m.db.WithContext(ctx).Model(&LikeCount{}).Where("id = ?", id).Updates(values).Error
}
