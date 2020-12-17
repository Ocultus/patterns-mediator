package models

import (
	"fmt"
)

type Daemon struct {
	Name   string
	Actors []Actor
}

func (c *Daemon) Register(u *Actor) {
	fmt.Println("Register :", u)
	c.Actors = append(c.Actors, *u)

	fmt.Println(c.Actors)
}

func (c *Daemon) SendMessage(message string, fromPID int, to int) {
	if to != -1 {
		for _, item := range c.Actors {
			if item.PID == to {
				item.Receive(message, fromPID)
			}
		}
		return
	}
	for _, item := range c.Actors {
		item.Receive(message, fromPID)
	}
}
