package v1

import (
	"context"
	"github.com/cold-runner/skylark/internal/model/user"
	"github.com/cold-runner/skylark/internal/pkg/code"
	"github.com/marmotedu/errors"
)

var (
	qqLoginStrategy       loginStrategy
	passwordLoginStrategy loginStrategy
	phoneLoginStrategy    loginStrategy
)

type loginStrategy interface {
	login(context.Context, *user.LoginUser, *serviceV1) (*user.LoggedUser, error)
}

type qqUnionIdLogin struct{}
type passwordLogin struct{}
type phoneLogin struct{}

func newLoginStrategy(strategy string) loginStrategy {
	switch {
	case strategy == user.PHONE:
		return &phoneLogin{}
	case strategy == user.QQ:
		return &qqUnionIdLogin{}
	case strategy == user.PASSWORD:
		return &passwordLogin{}
	}
	panic("不支持的登陆模式")
}

func (s *serviceV1) ProcessLogin(c context.Context, loginUser *user.LoginUser) (loggedUser *user.LoggedUser, err error) {
	// 策略模式选择登陆策略
	switch {
	case loginUser.Type == user.QQ:
		return qqLoginStrategy.login(c, loginUser, s)
	case loginUser.Type == user.PASSWORD:
		return passwordLoginStrategy.login(c, loginUser, s)
	case loginUser.Type == user.PHONE:
		return phoneLoginStrategy.login(c, loginUser, s)
	}
	return nil, errors.WithCode(code.ErrLoginType, "", nil)
}
