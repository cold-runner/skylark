package other

import "time"

type PostWithAllDetail struct {
	ID          string    `gorm:"column:id"`          // 自然主键
	CreatedAt   time.Time `gorm:"column:created_at"`  // 创建时间
	Title       string    `gorm:"column:title"`       // 博文标题
	PictureURL  string    `gorm:"column:picture_url"` // 博文标题配图
	Summary     string    `gorm:"column:summary"`     // 博文概览
	Content     string    `gorm:"column:content"`     // 博文内容
	Temperature int64     `gorm:"column:temperature"` // 博文热度（排序文章时用）
	Like        int64     `gorm:"column:like"`        // 博文点赞量
	Watch       int64     `gorm:"column:watch"`       // 观看量
	Star        int64     `gorm:"column:star"`        // 收藏数量
	Tag         []string  `gorm:"-"`                  // 标签

	UserId    string `gorm:"column:user_id"`   // 用户唯一标识
	StuNum    string `gorm:"column:stu_num"`   // 学号
	Name      string `gorm:"column:name"`      // 姓名
	Gender    string `gorm:"column:gender"`    // 用户性别：女，男，其他
	College   string `gorm:"column:college"`   // 用户所在学院
	Major     string `gorm:"column:major"`     // 用户专业
	Grade     string `gorm:"column:grade"`     // 用户年级：大一，大二，大三，大四，研究生,毕业生
	Province  string `gorm:"column:province"`  // 用户家乡省份
	Age       int64  `gorm:"column:age"`       // 用户年龄
	PhotoURL  string `gorm:"column:photo_url"` // 照片url
	Introduce string `gorm:"column:introduce"` // 用户个人介绍
	Avatar    string `gorm:"column:avatar"`    // 用户头像url
}

type PostWithTag struct {
	ID      string `gorm:"column:id"` // 自然主键
	PostId  string `gorm:"column:post_id"`
	TagName string `gorm:"column:name"`
}
