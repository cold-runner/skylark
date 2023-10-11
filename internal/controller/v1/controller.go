package v1

import (
	"context"
	"github.com/cold-runner/skylark/internal/controller"
	"github.com/cold-runner/skylark/internal/service"
)

type controllerV1 struct {
	context    context.Context
	serviceIns service.Interface
}

func (c *controllerV1) NewInstance(context context.Context, serviceIns service.Interface) controller.Interface {
	return &controllerV1{
		context:    context,
		serviceIns: serviceIns,
	}
}

func NewFactory() controller.Factory {
	return new(controllerV1)
}
