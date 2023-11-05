package v1

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cold-runner/skylark/internal/model"
	"github.com/cold-runner/skylark/internal/pkg/errCode"
	"github.com/cold-runner/skylark/internal/store/mysql"
	"github.com/cold-runner/skylark/internal/store/mysql/query"
	pkgErr "github.com/pkg/errors"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

func (s *serviceV1) BindQq(c context.Context, ctx *app.RequestContext, stuNum, qqUnionId string) *errors.Error {
	// 查找是否存在用户
	exist, err := s.storeIns.Exist(c, mysql.LarkByStuNum(stuNum))
	switch {
	// 未找到用户
	case !exist:
		return errCode.ErrUserNotExist

	// 内部错误
	case err != nil:
		hlog.Errorf("绑定qqUnionId错误！err: %#v", err)
		return errCode.ErrUnknown
	}

	err = s.storeIns.UpdateLarkSelect(c, []field.Expr{query.Lark.QqUnionID}, []gen.Condition{mysql.LarkQqUnionIdIsNull()}, &model.Lark{QqUnionID: qqUnionId})
	switch {
	// 该qqUnionId已被绑定
	case pkgErr.Is(err, gorm.ErrDuplicatedKey):
		return errCode.ErrRepeatBindQq

	// 内部错误
	case err != nil:
		hlog.Errorf("绑定qqUnionId错误！err: %#v", err)
		return errCode.ErrUnknown
	}

	return nil
}
