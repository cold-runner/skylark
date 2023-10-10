package v1

import (
	"context"
	"github.com/cold-runner/skylark/internal/model/user"
	"github.com/cold-runner/skylark/internal/pkg/util"
	"strconv"
	"time"
)

func (s *serviceV1) SendRegisterSms(phone string, paramSet []string) error {
	err := s.smsClient.SendToSingle(phone, paramSet)
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceV1) SetExpiration(key, value string, expiration int) error {
	err := s.cacheIns.SetExpiration(key, value, time.Duration(expiration)*time.Minute)
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceV1) GetDelete(key string) (string, error) {
	value, err := s.cacheIns.GetDel(key)
	if err != nil {
		return "", err
	}
	return value, nil
}

func (s *serviceV1) Register(c context.Context, register *user.Register) error {
	cryptPassword, _ := util.Crypt(register.Password)
	// 根据系统时间戳生成唯一文件名
	fileUrl, err := s.ossIns.UploadFormFile(register.StuCardPhoto, strconv.FormatInt(util.GenUnixTimeNano(), 10))
	if err != nil {
		return err
	}
	lark := &user.Lark{
		StuNum:     register.StuNum,
		Password:   cryptPassword,
		College:    register.College,
		Major:      register.Major,
		Grade:      register.Grade,
		Name:       register.Name,
		Gender:     util.GenderTransform(register.Gender),
		StuCardUrl: fileUrl,
		Phone:      register.Phone,
		Legal:      false,
	}
	if err = s.storeIns.Register(c, lark); err != nil {
		return err
	}
	return nil
}
