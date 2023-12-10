package util

import (
	"mime/multipart"

	"github.com/cold-runner/skylark/biz/entity/consts"
	"github.com/dlclark/regexp2"
	"github.com/go-playground/validator/v10"
	"github.com/hertz-contrib/binding/go_playground"
)

// 各个校验函数的名称
const (
	STU_NUM     = "stuNum"
	QQ_UNION_ID = "qqUnionId"
	WECHAT      = "wechat"
	QQ          = "qq"
	NAME        = "stuName"
	SMS_CODE    = "smsCode"
	PASSWORD    = "password"
	LOGIN_TYPE  = "loginType"
	IS_IMAGE    = "isImage"
	COLLEGE     = "college"
	GRADE       = "grade"
	MAJOR       = "major"
	GENDER      = "gender"
	PROVINCE    = "province"
	PHONE       = "phone"
)

var validatorMap = map[string]string{
	STU_NUM:     `^\d{8}$`,
	QQ_UNION_ID: `^\d{1,200}$`,
	WECHAT:      `^[a-zA-Z][-_a-zA-Z0-9]{5,19}$`,
	QQ:          `^[1-9][0-9]{4,10}$`,
	NAME:        `^[\u4e00-\u9fa5·]{2,16}$`,
	SMS_CODE:    `^\d{1,10}$`,
	PASSWORD:    `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*()\-_=+{};:,<.>]).{8,16}$`, // 8-16位，包括至少1个大写字母，1个小写字母，1个数字，1个特殊字符
	LOGIN_TYPE:  "^(phone|password)$",
	PHONE:       `^1[3456789]\d{9}$`,
}

func CustomValidator() *go_playground.Validator {
	vd := go_playground.NewValidator()
	vd.SetValidateTag("vd")
	vdEngine := vd.Engine().(*validator.Validate)

	// 学号校验
	vdEngine.RegisterValidation(STU_NUM, func(fl validator.FieldLevel) bool {
		val := fl.Field().Interface().(string)
		reg, _ := regexp2.Compile(validatorMap[STU_NUM], 0)
		match, _ := reg.MatchString(val)
		return match
	})
	// qqUnionId校验
	vdEngine.RegisterValidation(QQ_UNION_ID, func(fl validator.FieldLevel) bool {
		val := fl.Field().Interface().(string)
		reg, _ := regexp2.Compile(validatorMap[QQ_UNION_ID], 0)
		match, _ := reg.MatchString(val)
		return match
	})
	// 微信号校验
	vdEngine.RegisterValidation(WECHAT, func(fl validator.FieldLevel) bool {
		val := fl.Field().Interface().(string)
		reg, _ := regexp2.Compile(validatorMap[WECHAT], 0)
		match, _ := reg.MatchString(val)
		return match
	})
	// qq号校验
	vdEngine.RegisterValidation(QQ, func(fl validator.FieldLevel) bool {
		val := fl.Field().Interface().(string)
		reg, _ := regexp2.Compile(validatorMap[QQ], 0)
		match, _ := reg.MatchString(val)
		return match
	})
	// 中文姓名校验
	vdEngine.RegisterValidation(NAME, func(fl validator.FieldLevel) bool {
		val := fl.Field().Interface().(string)
		reg, _ := regexp2.Compile(validatorMap[NAME], 0)
		match, _ := reg.MatchString(val)
		return match
	})
	// 手机验证码校验
	vdEngine.RegisterValidation(SMS_CODE, func(fl validator.FieldLevel) bool {
		val := fl.Field().Interface().(string)
		reg, _ := regexp2.Compile(validatorMap[SMS_CODE], 0)
		match, _ := reg.MatchString(val)
		return match
	})
	// 密码强度校验
	vdEngine.RegisterValidation(PASSWORD, func(fl validator.FieldLevel) bool {
		val := fl.Field().Interface().(string)
		reg, _ := regexp2.Compile(validatorMap[PASSWORD], 0)
		match, _ := reg.MatchString(val)
		return match
	})
	// 登陆类型校验
	vdEngine.RegisterValidation(LOGIN_TYPE, func(fl validator.FieldLevel) bool {
		val := fl.Field().Interface().(string)
		reg, _ := regexp2.Compile(validatorMap[LOGIN_TYPE], 0)
		match, _ := reg.MatchString(val)
		return match
	})
	// 手机号校验
	vdEngine.RegisterValidation(PHONE, func(fl validator.FieldLevel) bool {
		val := fl.Field().Interface().(string)
		reg, _ := regexp2.Compile(validatorMap[PHONE], 0)
		match, _ := reg.MatchString(val)
		return match
	})

	// 图片校验
	vdEngine.RegisterValidation(IS_IMAGE, func(fl validator.FieldLevel) bool {
		val := fl.Field().Interface().(*multipart.FileHeader)
		return IsFormFileImage(val)
	})
	// 学院校验
	vdEngine.RegisterValidation(COLLEGE, func(fl validator.FieldLevel) bool {
		val := fl.Field().Interface().(string)
		if _, err := consts.ParseCollege(val); err != nil {
			return false
		}
		return true
	})
	// 年级校验
	vdEngine.RegisterValidation(GRADE, func(fl validator.FieldLevel) bool {
		val := fl.Field().Interface().(string)
		if _, err := consts.ParseGrade(val); err != nil {
			return false
		}
		return true
	})
	// 专业校验
	vdEngine.RegisterValidation(MAJOR, func(fl validator.FieldLevel) bool {
		val := fl.Field().Interface().(string)
		if _, err := consts.ParseMajor(val); err != nil {
			return false
		}
		return true
	})
	// 性别校验
	vdEngine.RegisterValidation(GENDER, func(fl validator.FieldLevel) bool {
		val := fl.Field().Interface().(string)
		if _, err := consts.ParseGender(val); err != nil {
			return false
		}
		return true
	})
	// 省份校验
	vdEngine.RegisterValidation(PROVINCE, func(fl validator.FieldLevel) bool {
		val := fl.Field().Interface().(string)
		if _, err := consts.ParseProvince(val); err != nil {
			return false
		}
		return true
	})

	return vd
}
