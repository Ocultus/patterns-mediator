package models

import (
	"fmt"
	"math/rand"
)

type IActor interface{
	SendMessage(message string,to int)
	Receive(message string,fromPID int) 
	Spawn() *IActor
}

type Actor struct {
	PID	int
	VirtualMachine *Daemon
	IsActive	bool
}

func (u *Actor) SendMessage(message string,to int) {
	u.VirtualMachine.SendMessage(message,u.PID,to)
}

func (u *Actor) Receive(message string,fromPID int) {
	if message == "STOP" {
		u.IsActive = false
		fmt.Println("PID :", u.PID," Stop")
		return
	}
	fmt.Println(fromPID,"=>",u.PID,"Message:",message)

}

func (u *Actor) Spawn() *Actor{
	return &Actor{
		PID: rand.Intn(1000),
		VirtualMachine: u.VirtualMachine,
		IsActive: true,
	}
}

