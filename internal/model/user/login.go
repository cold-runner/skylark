package user

const (
	QQ       = "qq"
	PHONE    = "phone"
	PASSWORD = "password"
)

// LoginUser 尝试登陆的用户
type LoginUser struct {
	Type string `vd:"loginType($)" json:"loginType"`

	StuNum   string `vd:"regexp('^\\d{8}$')" json:"stuNum"`
	Password string `vd:"password($)" json:"password"`

	Phone   string `vd:"phone($)" json:"phone"`
	SmsCode string `vd:"smsCode($)" json:"smsCode"`

	QqUnionId string `vd:"qqUnionId($)" json:"qqUnionId"`
	Wechat    string `vd:"wechat($)" json:"wechat"`
}

// LoggedUser 返回给已登陆的用户信息
type LoggedUser struct {
	// TODO 角色信息

	Avatar  string // 头像
	StuNum  string // 学号
	Name    string // 姓名
	College string // 学院
	Major   string // 专业
	Grade   string // 年级

	// TODO 其他业务信息

	//Gender  int    `gorm:"type:tinyint(1);default:null"`   // 性别
	//StuCardUrl string    `gorm:"type:varchar(191)"`                  // 学生证照片url
	//Phone      string    `gorm:"type:char(11)"`                      // 手机号码
	//Province   string    `gorm:"type:varchar(3);default:null"`       // 省份
	//Age        uint8     `gorm:"type:tinyint unsigned;default:null"` // 年龄
	//Birth      time.Time `gorm:"type:date;default:null"`             // 出生日期
	//PhotoUrl   string    `gorm:"type:varchar(191);default:null"`     // 照片
	//Email      string    `gorm:"type:varchar(191);default:null"`     // 电子邮箱
	//Introduce  string    `gorm:"type:text;default:null"`             // 个人介绍
}

type PasswordLogin struct {
	StuNum   string `form:"stuNum" json:"stuNum"  vd:"regexp('^\\d{8}$'); msg:'Illegal format'"`
	Password string `vd:"password($)" form:"password" json:"password"`
}
