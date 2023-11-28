package entity

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cold-runner/skylark/internal/domain/model_gen"
	"github.com/cold-runner/skylark/internal/infrastructure/oss"
	"github.com/cold-runner/skylark/internal/infrastructure/store"
	"github.com/cold-runner/skylark/internal/infrastructure/util"
	"github.com/google/uuid"
	"mime/multipart"
	"strconv"
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
	Password     string                `vd:"password($)" form:"password,required"`

	SmsCode string `vd:"smsCode($)" form:"smsCode,required"`
}

func (r *Register) PreCreat(c context.Context, ctx *app.RequestContext, ossIns oss.Oss) (*model_gen.Lark, error) {
	// 密码加盐加密
	cryptPassword, err := util.Crypt(r.Password)
	if err != nil {
		return nil, ctx.Error(err)
	}

	// 学生证照片文件上传，文件名根据系统时间戳生成，保证唯一
	fileUrl, err := ossIns.UploadFormFile(c, r.StuCardPhoto, strconv.FormatInt(util.GenUnixTimeNano(), 10))
	if err != nil {
		return nil, ctx.Error(err)
	}

	return &model_gen.Lark{
		ID:         uuid.New(),
		StuNum:     r.StuNum,
		Name:       r.Name,
		Password:   cryptPassword,
		Gender:     r.Gender,
		College:    r.College,
		Major:      r.Major,
		Grade:      r.Grade,
		StuCardURL: fileUrl,
		Phone:      r.Phone,
		State:      0,
	}, nil
}

func (r *Register) Store(c context.Context, db store.Store, ov *model_gen.Lark) error {
	return db.CreateLark(c, ov)
}
