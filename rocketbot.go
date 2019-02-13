package rocketbot

import (
	"sync"

	"github.com/blushft/go-rocketbot/client"
)

type bot struct {
	exit   chan bool
	client *client.Client

	sync.RWMutex
	inputs   map[string]interface{}
	commands map[string]Command
}

type Command interface {
	Exec(...string) ([]byte, error)
	Help() string
	Name() string
	Description() string
}
