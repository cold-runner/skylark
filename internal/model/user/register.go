package user

import (
	"mime/multipart"
)

// Register 注册
type Register struct {
	StuNum       string                `vd:"stuNum($)" form:"stuNum,required"`
	Name         string                `vd:"name($)" form:"name,required"`
	College      string                `vd:"college($)" form:"college,required"`
	Major        string                `vd:"major($)" form:"major,required"`
	Grade        string                `vd:"grade($)" form:"grade,required"`
	Gender       string                `vd:"gender($)" form:"gender,required"`
	StuCardPhoto *multipart.FileHeader `vd:"isImage($)" form:"stuCardPhoto,required"`
	Phone        string                `vd:"phone($)" form:"phone,required"`
	SmsCode      string                `vd:"smsCode($)" form:"smsCode,required"`
	Password     string                `vd:"password($)" form:"password,required"`
}
