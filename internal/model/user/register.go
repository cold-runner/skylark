package user

import (
	"github.com/cloudwego/hertz/pkg/app/server/binding"
	"github.com/cold-runner/skylark/internal/model/consts"
	"github.com/cold-runner/skylark/internal/pkg/util"
	"github.com/marmotedu/errors"
	"mime/multipart"
)

// Register 注册
type Register struct {
	StuNum       string                `vd:"regexp(^\d{8}$)" form:"stuNum,required"`
	Password     string                `vd:"regexp(/^\S*(?=\S{6,})(?=\S*\d)(?=\S*[A-Z])(?=\S*[a-z])(?=\S*[!@#$%^&*? ])\S*$/)" form:"password,required"`
	College      string                `vd:"college($)" form:"college,required"`
	Major        string                `vd:"major($)" form:"major,required"`
	Grade        string                `vd:"grade($)" form:"grade,required"`
	Name         string                `vd:"" form:"name,required"`
	Gender       string                `vd:"gender($)" form:"gender,required"`
	StuCardPhoto *multipart.FileHeader `vd:"isImage($)" form:"stuCardPhoto,required"`
	Phone        string                `vd:"regexp(^(?:(?:\+|00)86)?1\d{10}$)" form:"phone,required"`
	SmsCode      string                `vd:"" form:"smsCode,required"`
}

func CheckRegister() {
	var college consts.College
	binding.MustRegValidateFunc(college.CheckFuncName(), func(args ...interface{}) error {
		c, ok := args[0].(string)
		if !ok {
			return errors.New("学院校验失败，类型不为string")
		}
		if _, ok := consts.CollegeList[c]; !ok {
			return errors.New("不存在的学院")
		}
		return nil
	})

	var grade consts.Grade
	binding.MustRegValidateFunc(grade.CheckFuncName(), func(args ...interface{}) error {
		c, ok := args[0].(string)
		if !ok {
			return errors.New("学院校验失败，类型不为string")
		}
		if _, ok := consts.GradeList[c]; !ok {
			return errors.New("不存在的年级")
		}
		return nil
	})

	var major consts.Major
	binding.MustRegValidateFunc(major.CheckFuncName(), func(args ...interface{}) error {
		c, ok := args[0].(string)
		if !ok {
			return errors.New("专业校验失败，类型不为string")
		}
		if _, ok := consts.MajorList[c]; !ok {
			return errors.New("不存在的专业")
		}
		return nil
	})

	var gender consts.Gender
	binding.MustRegValidateFunc(gender.CheckFuncName(), func(args ...interface{}) error {
		c, ok := args[0].(string)
		if !ok {
			return errors.New("性别校验失败，类型不为string")
		}
		if _, ok := consts.GenderList[c]; !ok {
			return errors.New("不存在的性别")
		}
		return nil
	})

	var province consts.Province
	binding.MustRegValidateFunc(province.CheckFuncName(), func(args ...interface{}) error {
		c, ok := args[0].(string)
		if !ok {
			return errors.New("省份校验失败，类型不为string")
		}
		if _, ok := consts.ProvinceList[c]; !ok {
			return errors.New("请输入完整的中国省份名！如：辽宁")
		}
		return nil
	})

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
}
