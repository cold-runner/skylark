package v1

import (
	"github.com/cold-runner/skylark/internal/application"
	"github.com/cold-runner/skylark/internal/infrastructure/cache"
	"github.com/cold-runner/skylark/internal/infrastructure/oss"
	"github.com/cold-runner/skylark/internal/infrastructure/sms"
	"github.com/cold-runner/skylark/internal/infrastructure/store"
)

type serviceV1 struct {
	cacheIns  cache.Cache
	ossIns    oss.Oss
	smsClient sms.Sms
	storeIns  store.Store
}

func (s *serviceV1) NewInstance(cacheIns cache.Cache, ossIns oss.Oss, smsClient sms.Sms, storeIns store.Store) application.Interface {
	return &serviceV1{
		cacheIns:  cacheIns,
		ossIns:    ossIns,
		smsClient: smsClient,
		storeIns:  storeIns,
	}
}

func NewFactory() application.Factory {
	return new(serviceV1)
}
