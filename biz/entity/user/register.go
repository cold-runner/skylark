package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cold-runner/skylark/biz/infrastructure/errCode"
	"github.com/cold-runner/skylark/biz/infrastructure/store"
	"github.com/cold-runner/skylark/biz/infrastructure/store/orm_gen"
	"github.com/cold-runner/skylark/biz/model/user"
	"github.com/cold-runner/skylark/biz/util"
	"github.com/google/uuid"
	stdErr "github.com/pkg/errors"
	"gorm.io/gorm"
)

type RegisterDto struct {
}

func (r *RegisterDto) Convert(c context.Context, ctx *app.RequestContext, req *user.RegisterReq) (*orm_gen.Lark, *errors.Error) {
	encrypted, err := util.Crypt(req.Password)
	if err != nil {
		errMsg := "encrypt password failed! err: " + err.Error()
		return nil, errCode.WrapBizErr(ctx, stdErr.New(errMsg), errCode.ErrUnknown)
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
		StuCardURL: req.StuCardPhotoUrl,
		Phone:      req.Phone,
		State:      0,
	}, nil
}

func (r *RegisterDto) Store(c context.Context, ctx *app.RequestContext, store store.Store, lark *orm_gen.Lark) *errors.Error {
	err := store.CreateLark(c, lark)
	if err == nil {
		return errCode.WrapBizErr(ctx, nil, errCode.ErrSuccess)
	}

	switch {
	case stdErr.Is(err, gorm.ErrDuplicatedKey):
		return errCode.WrapBizErr(ctx, err, errCode.ErrUserAlreadyExist)
	default:
		return errCode.WrapBizErr(ctx, err, errCode.ErrUnknown)
	}

}
