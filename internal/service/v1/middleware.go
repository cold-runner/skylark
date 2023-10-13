package v1

import (
	"context"
	"github.com/cold-runner/skylark/internal/model/user"
	"github.com/cold-runner/skylark/internal/pkg/code"
	"github.com/cold-runner/skylark/internal/service/share"
	"github.com/marmotedu/errors"
)

func (s *serviceV1) Authenticator(c context.Context, loginUser *user.LoginUser) (*user.LoggedUser, error) {
	switch {
	case loginUser.Type == user.PASSWORD:
		return share.PasswordLogin(c, loginUser, s.storeIns)
	case loginUser.Type == user.QQ:
		return share.QqUnionIdLogin(c, loginUser, s.storeIns)
	case loginUser.Type == user.PHONE:
		return share.PhoneLogin(c, loginUser, s.storeIns, s.cacheIns)
	}
	return nil, errors.WithCode(code.ErrUnknown, "", nil)
}
