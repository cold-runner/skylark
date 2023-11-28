package v1

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cold-runner/skylark/internal/application/share"
	"github.com/cold-runner/skylark/internal/domain/entity"
	"github.com/cold-runner/skylark/internal/infrastructure/errCode"
)

func (s *serviceV1) Authenticator(c context.Context, ctx *app.RequestContext, loginUser *entity.LoginUser) (loggedUser *entity.LoggedInUser, err error) {
	switch {
	case loginUser.Type == entity.PASSWORD:
		hlog.Debugf("密码登陆 学号：%s 密码：%s", loginUser.StuNum, loginUser.Password)
		return share.PasswordLogin(c, ctx, loginUser, s.storeIns)
	case loginUser.Type == entity.QQ:
		hlog.Debugf("qq登陆 qqUnionId：%s", loginUser.QqUnionId)
		return share.QqUnionIdLogin(c, ctx, loginUser, s.storeIns)
	case loginUser.Type == entity.PHONE:
		hlog.Debugf("手机号登陆 手机号：%s 验证码：%s", loginUser.Phone, loginUser.SmsCode)
		return share.PhoneLogin(c, ctx, loginUser, s.storeIns, s.cacheIns)
	}
	return nil, errCode.ErrUnknown
}
