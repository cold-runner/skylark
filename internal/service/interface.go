package service

import (
	"context"
	"github.com/cold-runner/skylark/internal/model/user"
)

// Service defines functions used to return resource interface.
type Service interface {
	ArticleSrv
	CommentSrv
	LarkSrv
	PublicSrv
}

type LarkSrv interface {
}

type CommentSrv interface {
}

type ArticleSrv interface {
}

type PublicSrv interface {
	Register(c context.Context, register *user.Register) error
	SendRegisterSms(phone string, paramSet []string) error

	/* 服务内部调用 */

	SetExpiration(key, value string, expiration int) error
	GetDelete(key string) (value string, err error)
}
