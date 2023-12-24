package user

import (
	"context"
	"encoding/base64"
	"github.com/cold-runner/skylark/biz/infrastructure/oss"
	"strconv"
	"time"

	"github.com/cold-runner/skylark/biz/infrastructure/errCode"
	"github.com/cold-runner/skylark/biz/infrastructure/store"
	"github.com/cold-runner/skylark/biz/infrastructure/store/orm_gen"
	"github.com/cold-runner/skylark/biz/model/user"
	"github.com/cold-runner/skylark/biz/util"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/google/uuid"
	stdErr "github.com/pkg/errors"
	"gorm.io/gorm"
)

type RegisterDto struct {
}

func (r *RegisterDto) IsRegistered(c context.Context, ctx *app.RequestContext, req *user.RegisterReq) *errors.Error {
	storeIns := store.GetIns()
	_, err := storeIns.GetLark(c, storeIns.LarkByStuNum(req.StuNum))
	if err == nil {
		return errCode.WrapBizErr(ctx, stdErr.New("学号为"+req.StuNum+"的用户已注册!"), errCode.ErrUserAlreadyExist)
	}
	if stdErr.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return errCode.WrapBizErr(ctx, err, errCode.ErrUnknown)
}

func (r *RegisterDto) Convert(c context.Context, ctx *app.RequestContext, req *user.RegisterReq) (*orm_gen.Lark, *errors.Error) {
	ossIns := oss.GetIns()
	encrypted, err := util.Crypt(req.Password)
	if err != nil {
		errMsg := "encrypt password failed! err: " + err.Error()
		return nil, errCode.WrapBizErr(ctx, stdErr.New(errMsg), errCode.ErrUnknown)
	}

	fileType, err := util.FileTypeFromBs64(req.StuCardPhoto)
	if err != nil || (fileType != "jpg" && fileType != "png") {
		return nil, errCode.WrapBizErr(ctx, err, errCode.ErrValidation)
	}

	// 保存图片到文件
	data, _ := base64.StdEncoding.DecodeString(req.StuCardPhoto)

	url, err := ossIns.UploadFileFromBytes(c, data, strconv.FormatInt(time.Now().UnixNano(), 10), fileType)
	if err != nil {
		return nil, errCode.WrapBizErr(ctx, err, errCode.ErrUnknown)
	}
	return &orm_gen.Lark{
		ID:         uuid.New(),
		StuNum:     req.StuNum,
		Name:       req.StuName,
		Password:   encrypted,
		Gender:     req.Gender,
		College:    req.College,
		Major:      req.Major,
		Grade:      req.Grade,
		StuCardURL: url,
		Phone:      req.Phone,
		State:      0,
	}, nil
}

func (r *RegisterDto) Store(c context.Context, ctx *app.RequestContext, lark *orm_gen.Lark) *errors.Error {
	storeIns := store.GetIns()
	err := storeIns.CreateLark(c, lark)
	if err == nil {
		return nil
	}

	switch {
	case stdErr.Is(err, gorm.ErrDuplicatedKey):
		return errCode.WrapBizErr(ctx, err, errCode.ErrUserAlreadyExist)
	default:
		return errCode.WrapBizErr(ctx, err, errCode.ErrUnknown)
	}

}
