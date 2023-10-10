package v1

import (
	"github.com/cold-runner/skylark/internal/pkg/cache"
	"github.com/cold-runner/skylark/internal/pkg/oss"
	"github.com/cold-runner/skylark/internal/pkg/sms"
	"github.com/cold-runner/skylark/internal/service"
	"github.com/cold-runner/skylark/internal/store"
)

type serviceV1 struct {
	cacheIns  cache.Cache
	ossIns    oss.Oss
	smsClient sms.Sms
	storeIns  store.Store
}

func NewServiceV1(cacheIns cache.Cache, ossIns oss.Oss, smsClient sms.Sms, storeIns store.Store) service.Service {
	return &serviceV1{
		cacheIns:  cacheIns,
		ossIns:    ossIns,
		smsClient: smsClient,
		storeIns:  storeIns,
	}
}
