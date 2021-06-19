package model

type Message string

func NewMessage() Message {
	return Message("Hi there!")
}
