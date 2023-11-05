package v1

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cold-runner/skylark/internal/model/user"
	"github.com/cold-runner/skylark/internal/pkg/errCode"
	"github.com/cold-runner/skylark/internal/service/share"
)

func (s *serviceV1) Authenticator(c context.Context, ctx *app.RequestContext, loginUser *user.LoginUser) (*user.LoggedUser, *errors.Error) {
	switch {
	case loginUser.Type == user.PASSWORD:
		hlog.Debugf("密码登陆 学号：%s 密码：%s", loginUser.StuNum, loginUser.Password)
		return share.PasswordLogin(c, ctx, loginUser, s.storeIns)
	case loginUser.Type == user.QQ:
		hlog.Debugf("qq登陆 qqUnionId：%s", loginUser.QqUnionId)
		return share.QqUnionIdLogin(c, ctx, loginUser, s.storeIns)
	case loginUser.Type == user.PHONE:
		hlog.Debugf("手机号登陆 手机号：%s 验证码：%s", loginUser.Phone, loginUser.SmsCode)
		return share.PhoneLogin(c, ctx, loginUser, s.storeIns, s.cacheIns)
	}
	return nil, errCode.ErrUnknown
}
