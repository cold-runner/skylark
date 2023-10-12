package user

const (
	QQ       = "qq"
	PHONE    = "phone"
	PASSWORD = "password"
)

// LoginUser 尝试登陆的用户
type LoginUser struct {
	Type string `json:"loginType" vd:"regexp('^(qq|password)$')"`

	StuNum   string `json:"stuNum" vd:"regexp('^\\d{8}$')"`
	Password string `json:"password" vd:"(len($) > 0 && len($) < 30)"`

	Phone   string `json:"phone"`
	SmsCode string `json:"smsCode"`

	Qq string `json:"qq"`

	Wechat string `json:"wechat"`
}

// LoggedUser 返回给已登陆的用户信息
type LoggedUser struct {
	// TODO 角色信息

	Avatar  string `gorm:"type:varchar(191);default:null"` // 头像
	StuNum  string `gorm:"type:char(8);unique"`            // 学号
	Name    string `gorm:"type:varchar(191)"`              // 姓名
	College string `gorm:"type:char(16)"`                  // 学院
	Major   string `gorm:"type:char(30)"`                  // 专业
	Grade   string `gorm:"type:varchar(3)"`                // 年级

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
	Password string `form:"password" json:"password" vd:"(len($) > 0 && len($) < 30); msg:'Illegal password'"`
}
