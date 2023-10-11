package v1

import (
	"context"
	"github.com/cold-runner/skylark/internal/model/user"
	"github.com/cold-runner/skylark/internal/pkg/code"
	"github.com/cold-runner/skylark/internal/pkg/util"
	"github.com/marmotedu/errors"
	"github.com/marmotedu/iam/pkg/log"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func (s *serviceV1) SendRegisterSms(c context.Context, phone string, paramSet []string) error {
	err := s.smsClient.SendToSingle(c, phone, paramSet)
	switch {
	case err == nil:
		log.L(c).V(log.InfoLevel).Infof("向%s发送验证码：%s 过期时间：%s 分钟", phone, paramSet[0], paramSet[1])
		return nil

	// 服务内部错误
	default:
		e := errors.WithCode(code.ErrUnknown, "公共服务层SendRegisterSms方法失败！err: %v", err)
		log.L(c).Errorf("获取缓存失败！err: %v", err)
		return e
	}
}

func (s *serviceV1) SetExpiration(c context.Context, key, value string, expiration time.Duration) error {
	err := s.cacheIns.SetExpiration(c, key, value, expiration)
	switch {
	case err == nil:
		return nil

	// 服务内部错误
	default:
		e := errors.WithCode(code.ErrUnknown, "公共服务层SetExpiration方法失败！err: %v", err)
		log.L(c).Errorf("设置缓存失败！err: %v", e)
		return e
	}
}

func (s *serviceV1) GetSmsCode(c context.Context, key string) (string, error) {
	value, err := s.cacheIns.Get(c, key)
	switch {
	case err == nil:
		return value, nil

	// 处理业务错误
	case errors.Is(err, redis.Nil):
		return "", errors.WithCode(code.ErrSmsCodeExpired, "验证码已过期, err: %v", err)

	// 服务内部错误
	default:
		e := errors.WithCode(code.ErrUnknown, "公共服务层GetSmsCode方法失败！err: %v", err)
		log.L(c).Errorf("获取缓存失败！err: %v", e)
		return "", e
	}
}

func (s *serviceV1) ValidateSmsCode(c context.Context, phone, inCache, userPass string) error {
	if inCache != userPass {
		return errors.WithCode(code.ErrSmsCode, "验证码错误！缓存：%s, 用户：%s", inCache, userPass)
	}
	if err := s.cacheIns.Del(c, phone); err != nil {
		e := errors.WithCode(code.ErrUnknown, "公共服务层ValidateSmsCode方法失败！err: %v", err)
		log.L(c).Errorf("删除验证码失败！err: %v", e.Error())
		// 用户侧不能承担内部错误后果
		return nil
	}
	return nil
}

func (s *serviceV1) Register(c context.Context, register *user.Register) error {
	// 密码加盐加密
	cryptPassword, err := util.Crypt(register.Password)
	if err != nil {
		e := errors.WithCode(code.ErrUnknown, "util包Crypt方法失败！err: %v", err)
		log.L(c).Errorf("加密密码失败！err: %v", e)
		return e
	}

	// 学生证照片文件上传，文件名根据系统时间戳生成，保证唯一
	fileUrl, err := s.ossIns.UploadFormFile(c, register.StuCardPhoto, strconv.FormatInt(util.GenUnixTimeNano(), 10))
	if err != nil {
		e := errors.WithCode(code.ErrUnknown, "对象存储UploadFormFile失败！err: %v", err)
		log.L(c).Errorf("上传文件失败！err: %v", e)
		return e
	}

	err = s.storeIns.Register(c, &user.Lark{
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
	})
	switch {
	case err == nil:
		return nil
	// 用户已注册
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return errors.WithCode(code.ErrUserAlreadyExist, "已创建用户")

	// 服务内部错误
	default:
		e := errors.WithCode(code.ErrUnknown, "公共服务层Register方法失败！err: %v", err)
		log.L(c).Errorf("创建用户失败！err: %v", err)
		return e
	}
}
