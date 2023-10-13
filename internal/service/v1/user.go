package v1

import (
	"context"
	"github.com/cold-runner/skylark/internal/pkg/code"
	"github.com/cold-runner/skylark/internal/store/mysql"
	"github.com/marmotedu/errors"
)

func (s *serviceV1) BindQq(c context.Context, stuNum, qqUnionId string) error {
	err := s.storeIns.UpdateLark(c, qqUnionId, mysql.LarkByStuNum(stuNum), mysql.LarkQqUnionIdIsNull)
	if err != nil {
		return errors.WithCode(code.ErrRepeatBindQq, "", nil)
	}
	return nil
}
