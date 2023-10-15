package post

import (
	"github.com/cloudwego/hertz/pkg/common/json"
	"gorm.io/gorm"
)

type Draft struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (d *Draft) MarshalBinary() ([]byte, error) {
	return json.Marshal(d)
}

func (d *Draft) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, d)
}

type LearningPost struct {
	Title      string   `vd:"" json:"title"`                        // 标题
	PictureUrl []byte   `vd:"" json:"pictureUrl"`                   // 文章标题图
	UserId     uint     `vd:"" json:"userId"`                       // 作者 TODO 约束
	Summary    string   `vd:"" json:"summary"`                      // 摘要
	Content    string   `vd:"" json:"content" gorm:"type:longtext"` // 内容
	IsDraft    bool     `vd:"" json:"isDraft"`                      // 是否是草稿
	Theme      string   `vd:"" json:"theme"`                        // 所属主题
	Tags       []string `vd:"" json:"tags"`                         // 所带有的标签

}

type StoredLearningPost struct {
	LearningPost

	// 后台运营相关
	gorm.Model
	Temperature uint `gorm:"default:0"` // 文章热度
	Like        int  `gorm:"default:0"` // 点赞数量
	Watch       int  `gorm:"default:0"` // 观看量
	Star        uint `gorm:"default:0"` // 收藏数量
	Pass        bool `gorm:"default:0"` // 审核是否通过
	Report      bool `gorm:"default:0"` // 文章是否被举报
}
