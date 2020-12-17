package main

import (
	"local/models"
	"math/rand"
)

func main () { 

	beam := &models.Daemon{
		Name: "BEAM",
	}

	chrome := &models.Actor{
		PID: rand.Intn(1000),
		VirtualMachine: beam,
		IsActive: true,
	}

	mozilla := &models.Actor{
		PID: rand.Intn(1000),
		VirtualMachine: beam,
		IsActive: true,
	}

	codeOSS := &models.Actor{
		PID: rand.Intn(1000),
		VirtualMachine: beam,
		IsActive: true,
	}
	beam.Register(chrome)
	beam.Register(mozilla)
	beam.Register(codeOSS)
	chrome.SendMessage("STOP",mozilla.PID)
	chrome.SendMessage("I need 1 GB of RAM, especially",codeOSS.PID)
	chrome.SendMessage("I need more RAM!",-1)
}