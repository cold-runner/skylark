package post

import "gorm.io/gorm"

type Post struct {
	*gorm.Model
	UserId  uint     // 作者
	Content string   // 内容
	Like    int      // 点赞数量
	Watch   int      // 观看量
	Star    uint     // 收藏数量
	Pass    bool     // 审核是否通过
	Theme   string   // 所属话题
	Tags    []string // 所带有的标签
}
