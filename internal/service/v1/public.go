package v1

import (
	"context"
	"github.com/cold-runner/skylark/internal/model/user"
	"github.com/cold-runner/skylark/internal/pkg/code"
	"github.com/cold-runner/skylark/internal/pkg/config"
	"github.com/cold-runner/skylark/internal/pkg/util"
	"github.com/marmotedu/errors"
	"github.com/marmotedu/iam/pkg/log"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func (s *serviceV1) SendSms(c context.Context, phone string) error {
	// 查看是否已经发送过验证码
	_, err := s.getSmsCode(c, phone)

	// 已发送过且未过期
	if err == nil {
		return errors.WithCode(code.ErrAlreadySendSmsCode, "", nil)
	}

	// 内部错误
	if errors.ParseCoder(err).Code() == code.ErrUnknown {
		e := errors.WrapC(err, code.ErrUnknown, "查询注册验证码失败！err: %v", err)
		log.Errorf("%#v", e)
		return e
	}

	serverConf := config.GetConfig().ServerConfig()
	expirationConfig := serverConf.SmsExpiration          // 从全局配置对象获取验证码过期时间
	randCode, _ := util.RandSmsCode(serverConf.SmsNumber) // 从全局配置对象获取验证码长度，根据验证码长度生成验证码

	// 缓存验证码
	if err := s.cacheIns.SetExpiration(c, phone, randCode, time.Duration(expirationConfig)*time.Minute); err != nil {
		e := errors.WrapC(err, code.ErrUnknown, "缓存注册验证码失败！err: %v", err)
		log.Errorf("%#v", e)
		return e
	}

	// 发送验证码
	if err := s.smsClient.SendToSingle(c, phone, []string{randCode, strconv.Itoa(expirationConfig)}); err != nil {
		e := errors.WrapC(err, code.ErrUnknown, "注册发送验证码失败！err: %v", err)
		log.Errorf("%#v", e)
		return e
	}

	log.L(c).V(log.InfoLevel).Infof("向%s发送验证码：%s 过期时间：%d 分钟", phone, randCode, expirationConfig)
	return nil
}

func (s *serviceV1) Register(c context.Context, register *user.Register) error {
	// 从缓存中获取验证码
	storeSmsCode, err := s.getSmsCode(c, register.Phone)
	switch {
	case err == nil:
		{
			// 验证码校验
			correct, err := s.validateAndDelSmsCode(c, register.Phone, storeSmsCode, register.SmsCode)
			// 用户传入的验证码不正确
			if !correct {
				return errors.WithCode(code.ErrSmsCode, "", nil)
			}

			// 验证码正确但删除缓存失败
			if err != nil {
				e := errors.WrapC(err, code.ErrUnknown, "注册时删除验证码失败！err: %v", err)
				log.Errorf("%#v", e)
			}

			// 密码加盐加密
			cryptPassword, err := util.Crypt(register.Password)
			if err != nil {
				e := errors.WrapC(err, code.ErrUnknown, "注册时加盐加密密码失败！err: %v", err)
				log.Errorf("%#v", e)
				return e
			}

			// 学生证照片文件上传，文件名根据系统时间戳生成，保证唯一
			fileUrl, err := s.ossIns.UploadFormFile(c, register.StuCardPhoto, strconv.FormatInt(util.GenUnixTimeNano(), 10))
			if err != nil {
				e := errors.WrapC(err, code.ErrUnknown, "注册时上传学生证照片文件失败！err: %v", err)
				log.Errorf("%#v", e)
				return e
			}

			// 存储进数据库
			err = s.storeIns.Create(c, &user.Lark{
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
			// 注册成功
			case err == nil:
				return nil

			// 用户已注册
			case errors.Is(err, gorm.ErrDuplicatedKey):
				return errors.WrapC(err, code.ErrUserAlreadyExist, "", nil)

			// 服务内部错误
			default:
				e := errors.WrapC(err, code.ErrUnknown, "注册时存储至数据库失败！err: %v", err)
				log.Errorf("%#v", e)
				return e
			}
		}
	// 验证码过期
	case errors.ParseCoder(err).Code() == code.ErrSmsCodeExpired:
		return err

	// 内部错误
	default:
		// err.(errors.Coder).Code() == code.ErrUnknown
		e := errors.WrapC(err, code.ErrUnknown, "从缓存中获取注册验证码失败！err: %v", err)
		log.Errorf("%#v", e)
		return e
	}
}
