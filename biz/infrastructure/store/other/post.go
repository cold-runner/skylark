package other

import "time"

type PostWithAllDetail struct {
	ID           string             `gorm:"column:id" json:"id"`                       // 自然主键
	Title        string             `gorm:"column:title" json:"title"`                 // 博文标题
	CreatedAt    time.Time          `gorm:"column:created_at" json:"created_at"`       // 创建时间
	CoverImage   string             `gorm:"column:cover_image" json:"cover_image"`     // 博文标题配图
	Summary      string             `gorm:"column:summary" json:"summary"`             // 博文概览
	Content      string             `gorm:"column:content" json:"content"`             // 博文内容
	Temperature  int64              `gorm:"column:temperature" json:"temperature"`     // 博文热度（排序文章时用）
	Tags         []string           `gorm:"-" json:"tags"`                             // 标签
	LikeCount    int64              `gorm:"column:like_count" json:"like_count"`       // 博文点赞量
	ViewCount    int64              `gorm:"column:view_count" json:"view_count"`       // 观看量
	StarCount    int64              `gorm:"column:star_count" json:"star_count"`       // 收藏数量
	CommentCount int64              `gorm:"column:comment_count" json:"comment_count"` // 评论数量
	ShareCount   int64              `gorm:"column:share_count" json:"share_count"`     // 分享数量
	UserInfo     `json:"user_info"` // 用户信息
}

type UserInfo struct {
	UserId string `gorm:"column:user_id" json:"user_id"` // 用户唯一标识
	Name   string `gorm:"column:name" json:"name"`       // 姓名
}

type PostWithTag struct {
	ID      string `gorm:"column:id"` // 自然主键
	PostId  string `gorm:"column:post_id" json:"post_id"`
	TagName string `gorm:"column:name" json:"tag_name"`
}
