// +build wireinject

package di

import (
	"backend/model"

	"github.com/google/wire"
)

func InitializeEvent() model.Event {
	wire.Build(model.NewEvent, model.NewGreeter, model.NewMessage)
	return model.Event{}
}

func InitializeGreeter() model.Greeter {
	wire.Build(model.NewGreeter, model.NewMessage)
	return model.Greeter{}
}
