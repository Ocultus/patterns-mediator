package models

type Mediator interface {
	Register(actor Actor)	
	SendMessage(message string, fromPID string, to Actor) 
}