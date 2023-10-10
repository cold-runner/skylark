package v1

import (
	"context"
	"github.com/cold-runner/skylark/internal/controller"
	"github.com/cold-runner/skylark/internal/service"
)

type controllerV1 struct {
	context    context.Context
	serviceIns service.Service
}

func NewControllerV1(context context.Context, serviceIns service.Service) controller.Controller {
	return &controllerV1{
		context:    context,
		serviceIns: serviceIns,
	}
}
