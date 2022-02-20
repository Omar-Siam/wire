//go:build wireinject
// +build wireinject

package main

import (
	"Wire/greetings"
	"github.com/google/wire"
)

func InitializeEvent() greetings.Event {
	wire.Build(greetings.NewEvent, greetings.NewGreeter, greetings.NewMessage)
	return greetings.Event{}
}
