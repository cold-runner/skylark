package controller

import (
	"github.com/cold-runner/skylark/internal/pkg/util"
	"mime/multipart"
	"regexp"

	"github.com/cloudwego/hertz/pkg/app/server/binding"
	"github.com/cold-runner/skylark/internal/model/consts"
	"github.com/pkg/errors"
)

func Validator() {
	// TODO 学号校验

	// 微信号校验
	binding.MustRegValidateFunc("wechat", func(args ...interface{}) error {
		c, ok := args[0].(string)
		if !ok {
			return errors.New("微信号校验失败，类型不为string")
		}
		match, _ := regexp.MatchString("^[a-zA-Z][-_a-zA-Z0-9]{5,19}$", c)
		if !match {
			return errors.New("微信号校验失败")
		}
		return nil
	})
	// QQ号校验
	binding.MustRegValidateFunc("qq", func(args ...interface{}) error {
		c, ok := args[0].(string)
		if !ok {
			return errors.New("qq号校验失败，类型不为string")
		}
		match, _ := regexp.MatchString("^[1-9][0-9]{4,10}$", c)
		if !match {
			return errors.New("qq号校验失败")
		}
		return nil
	})
	// 中文姓名校验
	binding.MustRegValidateFunc("name", func(args ...interface{}) error {
		c, ok := args[0].(string)
		if !ok {
			return errors.New("中文姓名校验失败，类型不为string")
		}
		match, _ := regexp.MatchString("^[\u4e00-\u9fa5·]{2,16}$", c)
		if !match {
			return errors.New("中文姓名校验失败，应为2-16位纯汉字")
		}
		return nil
	})
	// 验证码校验
	binding.MustRegValidateFunc("smsCode", func(args ...interface{}) error {
		c, ok := args[0].(string)
		if !ok {
			return errors.New("验证码校验失败，类型不为string")
		}
		match, _ := regexp.MatchString("^\\d{1,10}$", c)
		if !match {
			return errors.New("验证码校验失败")
		}
		return nil
	})
	// 密码强度校验
	binding.MustRegValidateFunc("password", func(args ...interface{}) error {
		c, ok := args[0].(string)
		if !ok {
			return errors.New("密码校验失败，类型不为string")
		}
		match, _ := regexp.MatchString("^(?=.*\\d)(?=.*[a-zA-Z])(?=.*[^\\da-zA-Z\\s]).{6,20}$", c)
		if !match {
			return errors.New("密码校验失败，至少包含字母、数字、特殊字符，6-20位")
		}
		return nil
	})
	// 登陆类型校验
	binding.MustRegValidateFunc("loginType", func(args ...interface{}) error {
		c, ok := args[0].(string)
		if !ok {
			return errors.New("登陆类型校验失败，类型不为string")
		}
		match, _ := regexp.MatchString("^(phone|password)$", c)
		if !match {
			return errors.New("登陆方式校验失败，不支持的登陆类型")
		}
		return nil
	})
	// 手机号校验
	binding.MustRegValidateFunc("phone", func(args ...interface{}) error {
		c, ok := args[0].(string)
		if !ok {
			return errors.New("手机号校验失败，类型不为string")
		}
		match, _ := regexp.MatchString("^(?:(?:\\+|00)86)?1\\d{10}$", c)
		if !match {
			return errors.New("手机号校验失败")
		}
		return nil
	})
	// 图片校验
	binding.MustRegValidateFunc("isImage", func(args ...interface{}) error {
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
	binding.MustRegValidateFunc(consts.CheckCollegeFuncName(), func(args ...interface{}) error {
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
	binding.MustRegValidateFunc(consts.CheckGradeFuncName(), func(args ...interface{}) error {
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
	binding.MustRegValidateFunc(consts.CheckMajorFuncName(), func(args ...interface{}) error {
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
	binding.MustRegValidateFunc(consts.CheckGenderFuncName(), func(args ...interface{}) error {
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
	binding.MustRegValidateFunc(consts.CheckProvinceFuncName(), func(args ...interface{}) error {
		c, ok := args[0].(string)
		if !ok {
			return errors.New("省份校验失败，类型不为string")
		}
		if _, err := consts.ParseProvince(c); err != nil {
			return errors.New("省份校验失败，请参照表单内容")
		}
		return nil
	})
}
