package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cold-runner/skylark/biz/infrastructure/cache"
	"github.com/cold-runner/skylark/biz/infrastructure/errCode"
	"github.com/cold-runner/skylark/biz/infrastructure/store"
	"github.com/cold-runner/skylark/biz/infrastructure/store/mysql"
	"github.com/cold-runner/skylark/biz/model/user"
	"github.com/cold-runner/skylark/biz/util"
	"github.com/google/uuid"
	stdErr "github.com/pkg/errors"
	"gorm.io/gorm"
)

type LoginUser struct {
	Uuid uuid.UUID `json:"uuid"`
}

func (l *LoginUser) PasswordLogin(c context.Context, ctx *app.RequestContext, storeIns store.Store, req *user.PasswordLoginReq) *errors.Error {
	lark, err := storeIns.GetLark(c, mysql.LarkByStuNum(storeIns.(*mysql.MysqlIns), req.StuNum))

	switch {
	case err == nil:
		if util.VerifyPassword(req.Password, lark.Password) != nil {
			return errCode.WrapBizErr(ctx, stdErr.New("password is incorrect!"), errCode.ErrPasswordIncorrect)
		}
		ctx.Set("uuid", lark.ID.String())
		return nil
	case stdErr.Is(err, gorm.ErrRecordNotFound):
		errMsg := "user not exist!" + "recv stuNum: " + req.GetStuNum()
		return errCode.WrapBizErr(ctx, stdErr.New(errMsg), errCode.ErrUserNotFound)
	default:
		return errCode.WrapBizErr(ctx, err, errCode.ErrUnknown)
	}

}

func (l *LoginUser) PhoneLogin(c context.Context, storeIns store.Store, cacheIns cache.Cache, req *user.PhoneLoginReq) *errors.Error {
	return nil
}
