package user

import (
	"gorm.io/gorm"
	"time"
)

// Lark 用户 对应数据库表
type Lark struct {
	*gorm.Model
	StuNum     string    `gorm:"type:char(8);unique"`                   // 学号
	Name       string    `gorm:"type:varchar(191)"`                     // 姓名
	Password   string    `gorm:"type:varchar(191)"`                     // 登陆密码
	Gender     int       `gorm:"type:tinyint(1);default:null"`          // 性别
	College    string    `gorm:"type:char(16)"`                         // 学院
	Major      string    `gorm:"type:char(30)"`                         // 专业
	Grade      string    `gorm:"type:varchar(3)"`                       // 年级
	StuCardUrl string    `gorm:"type:varchar(191)"`                     // 学生证照片url
	Phone      string    `gorm:"type:char(11);unique"`                  // 手机号码
	Province   string    `gorm:"type:varchar(3);default:null"`          // 省份
	Age        uint8     `gorm:"type:tinyint unsigned;default:null"`    // 年龄
	Birth      time.Time `gorm:"type:date;default:null"`                // 出生日期
	PhotoUrl   string    `gorm:"type:varchar(191);default:null"`        // 照片
	Email      string    `gorm:"type:varchar(191);default:null;unique"` // 电子邮箱
	Introduce  string    `gorm:"type:text;default:null"`                // 个人介绍
	Avatar     string    `gorm:"type:varchar(191);default:null"`        // 头像
	// 社会化登陆
	QqUnionId     string `gorm:"type:varchar(191);default:null;unique"`
	WechatUnionId string `gorm:"type:varchar(191);default:null;unique"`

	// 是否是合法用户，0不合法，即未通过后台认证，1合法
	Legal bool `gorm:"type:tinyint;default:0"`
	// 状态，0禁用，1启用
	State bool `gorm:"type:tinyint(1);default:1"`
}
