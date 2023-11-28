package v1

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cold-runner/skylark/internal/application/share"
	"github.com/cold-runner/skylark/internal/domain/entity"
	"github.com/cold-runner/skylark/internal/domain/model_gen"
	"github.com/cold-runner/skylark/internal/infrastructure/errCode"
	"github.com/cold-runner/skylark/internal/infrastructure/store/mysql"
	"github.com/cold-runner/skylark/internal/infrastructure/store/mysql/query_gen"
	pkgErr "github.com/pkg/errors"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

func (s *serviceV1) GetLarkInfo(c context.Context, ctx *app.RequestContext, stuNum string) (*entity.LoggedInUser, error) {
	// 查找是否存在用户
	larkOV, err := s.storeIns.GetLark(c, mysql.LarkByStuNum(stuNum))
	switch {
	// 未找到用户
	case pkgErr.Is(err, gorm.ErrRecordNotFound):
		return nil, share.DebugfAndResp(ctx, err, errCode.ErrUserNotExist, "未找到目标用户！\n学号：%v", stuNum)

	// 内部错误
	case err != nil:
		return nil, share.ErrorfAndResp(ctx, err, errCode.ErrUnknown, "查询用户信息错误！\n学号：%v", stuNum)
	}

	lark := &entity.LoggedInUser{}
	return lark.Formated(larkOV), nil
}

func (s *serviceV1) BindQq(c context.Context, ctx *app.RequestContext, stuNum, qqUnionId string) error {
	// 查找是否存在用户
	exist, err := s.storeIns.Exist(c, mysql.LarkByStuNum(stuNum))
	switch {
	// 未找到用户
	case !exist:
		return share.DebugfAndResp(ctx, err, errCode.ErrUserNotExist, "未找到目标用户！\n学号：%v", stuNum)

	// 内部错误
	case err != nil:
		return share.ErrorfAndResp(ctx, err, errCode.ErrUnknown, "绑定qqUnionId错误！\n学号：%v\nqqUnionId: %v", stuNum, qqUnionId)
	}

	err = s.storeIns.UpdateLarkSelect(c, []field.Expr{query_gen.Lark.QqUnionID}, []gen.Condition{mysql.LarkQqUnionIdIsNull()}, &model_gen.Lark{QqUnionID: qqUnionId})
	switch {
	// 该qqUnionId已被绑定
	case pkgErr.Is(err, gorm.ErrDuplicatedKey):
		return share.WarnfAndResp(ctx, err, errCode.ErrRepeatBindQq, "已绑定过qqUnionId！\n学号：%v", stuNum)

	// 内部错误
	case err != nil:
		return share.ErrorfAndResp(ctx, err, errCode.ErrUnknown, "绑定qqUnionId错误！\n学号：%v\nqqUnionId: %v", stuNum, qqUnionId)
	}

	return nil
}
