package injector

import (
	"backend/model"
)

func InjectMessage() model.Message {
	message := model.NewMessage()
	return message
}

func InjectGreeter() model.Greeter {
	message := InjectMessage()
	return model.NewGreeter(message)
}

func InjectEvent() model.Event {
	greeter := InjectGreeter()
	return model.NewEvent(greeter)
}
