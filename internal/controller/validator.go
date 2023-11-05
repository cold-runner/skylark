package controller

import (
	"github.com/cloudwego/hertz/pkg/app/server/binding"
	"github.com/cold-runner/skylark/internal/model/consts"
	"github.com/cold-runner/skylark/internal/pkg/util"
	"github.com/dlclark/regexp2"
	"github.com/pkg/errors"
	"mime/multipart"
)

func Validator() *binding.ValidateConfig {
	validateConfig := &binding.ValidateConfig{}
	// TODO 学号校验
	validateConfig.MustRegValidateFunc("stuNum", func(args ...interface{}) error {
		c, ok := args[0].(string)
		if !ok {
			return errors.New("学号校验失败，类型不为string")
		}
		reg, _ := regexp2.Compile(`^\d{8}$`, 0)
		match, _ := reg.MatchString(c)
		if !match {
			return errors.New("学号校验失败")
		}
		return nil
	})
	// qqUnionId校验
	validateConfig.MustRegValidateFunc("qqUnionId", func(args ...interface{}) error {
		c, ok := args[0].(string)
		if !ok {
			return errors.New("qqUnionId校验失败，类型不为string")
		}
		reg, _ := regexp2.Compile(`^\d{1,200}$`, 0)
		match, _ := reg.MatchString(c)
		if !match {
			return errors.New("qqUnionId校验失败")
		}
		return nil
	})
	// 微信号校验
	validateConfig.MustRegValidateFunc("wechat", func(args ...interface{}) error {
		c, ok := args[0].(string)
		if !ok {
			return errors.New("微信号校验失败，类型不为string")
		}
		reg, _ := regexp2.Compile(`^[a-zA-Z][-_a-zA-Z0-9]{5,19}$`, 0)
		match, _ := reg.MatchString(c)
		if !match {
			return errors.New("微信号校验失败")
		}
		return nil
	})
	// QQ号校验
	validateConfig.MustRegValidateFunc("qq", func(args ...interface{}) error {
		c, ok := args[0].(string)
		if !ok {
			return errors.New("qq号校验失败，类型不为string")
		}
		reg, _ := regexp2.Compile(`^[1-9][0-9]{4,10}$`, 0)
		match, _ := reg.MatchString(c)
		if !match {
			return errors.New("qq号校验失败")
		}
		return nil
	})
	// 中文姓名校验
	validateConfig.MustRegValidateFunc("name", func(args ...interface{}) error {
		c, ok := args[0].(string)
		if !ok {
			return errors.New("中文姓名校验失败，类型不为string")
		}
		reg, _ := regexp2.Compile(`^[\u4e00-\u9fa5·]{2,16}$`, 0)
		match, _ := reg.MatchString(c)
		if !match {
			return errors.New("中文姓名校验失败，应为2-16位纯汉字")
		}
		return nil
	})
	// 验证码校验
	validateConfig.MustRegValidateFunc("smsCode", func(args ...interface{}) error {
		c, ok := args[0].(string)
		if !ok {
			return errors.New("验证码校验失败，类型不为string")
		}
		reg, _ := regexp2.Compile(`^\d{1,10}$`, 0)
		match, _ := reg.MatchString(c)
		if !match {
			return errors.New("验证码校验失败")
		}
		return nil
	})
	// 密码强度校验
	validateConfig.MustRegValidateFunc("password", func(args ...interface{}) error {
		c, ok := args[0].(string)
		if !ok {
			return errors.New("密码校验失败，类型不为string")
		}
		reg, _ := regexp2.Compile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*()\-_=+{};:,<.>]).{8,16}$`, 0)
		match, _ := reg.MatchString(c)
		if !match {
			return errors.New("8-16位，包括至少1个大写字母，1个小写字母，1个数字，1个特殊字符")
		}
		return nil
	})
	// 登陆类型校验
	validateConfig.MustRegValidateFunc("loginType", func(args ...interface{}) error {
		c, ok := args[0].(string)
		if !ok {
			return errors.New("登陆类型校验失败，类型不为string")
		}
		reg, _ := regexp2.Compile("^(phone|password)$", 0)
		match, _ := reg.MatchString(c)
		if !match {
			return errors.New("登陆方式校验失败，不支持的登陆类型")
		}
		return nil
	})
	// 图片校验
	validateConfig.MustRegValidateFunc("isImage", func(args ...interface{}) error {
		fileHeader, ok := args[0].(*multipart.FileHeader)
		if !ok {
			return errors.New("文件校验不通过")
		}
		if ok, err := util.IsFormFileImage(fileHeader); err != nil || !ok {
			return err
		}
		return nil
	})
	// 学院校验
	validateConfig.MustRegValidateFunc(consts.CheckCollegeFuncName(), func(args ...interface{}) error {
		c, ok := args[0].(string)
		if !ok {
			return errors.New("学院校验失败，类型不为string")
		}
		if _, err := consts.ParseCollege(c); err != nil {
			return errors.New("未知的学院")
		}
		return nil
	})
	// 年级校验
	validateConfig.MustRegValidateFunc(consts.CheckGradeFuncName(), func(args ...interface{}) error {
		c, ok := args[0].(string)
		if !ok {
			return errors.New("年级校验失败，类型不为string")
		}
		if _, err := consts.ParseGrade(c); err != nil {
			return errors.New("不支持的年级")
		}
		return nil
	})
	// 专业校验
	validateConfig.MustRegValidateFunc(consts.CheckMajorFuncName(), func(args ...interface{}) error {
		c, ok := args[0].(string)
		if !ok {
			return errors.New("专业校验失败，类型不为string")
		}
		if _, err := consts.ParseMajor(c); err != nil {
			return errors.New("未知专业")
		}
		return nil
	})
	// 性别校验
	validateConfig.MustRegValidateFunc(consts.CheckGenderFuncName(), func(args ...interface{}) error {
		c, ok := args[0].(string)
		if !ok {
			return errors.New("性别校验失败，类型不为string")
		}
		if _, err := consts.ParseGender(c); err != nil {
			return errors.New("不支持的性别")
		}
		return nil
	})
	// 省份校验
	validateConfig.MustRegValidateFunc(consts.CheckProvinceFuncName(), func(args ...interface{}) error {
		c, ok := args[0].(string)
		if !ok {
			return errors.New("省份校验失败，类型不为string")
		}
		if _, err := consts.ParseProvince(c); err != nil {
			return errors.New("省份校验失败，请参照表单内容")
		}
		return nil
	})

	return validateConfig
}
