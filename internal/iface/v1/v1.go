package v1

import (
	"github.com/cold-runner/skylark/internal/application"
	"github.com/cold-runner/skylark/internal/iface"
)

type userInterface struct {
	application application.Interface
}

// NewInstance 工厂方法模式
func (c *userInterface) NewInstance(application application.Interface) iface.UserInterface {
	return &userInterface{application: application}
}

// NewFactory 工厂方法模式
func NewFactory() iface.Factory {
	return new(userInterface)
}
