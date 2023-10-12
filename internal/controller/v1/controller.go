package v1

import (
	"github.com/cold-runner/skylark/internal/controller"
	"github.com/cold-runner/skylark/internal/service"
)

type controllerV1 struct {
	serviceIns service.Interface
}

func (c *controllerV1) NewInstance(serviceIns service.Interface) controller.Interface {
	return &controllerV1{
		serviceIns: serviceIns,
	}
}

func NewFactory() controller.Factory {
	return new(controllerV1)
}
