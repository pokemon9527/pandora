package tcp_box

import (
	"github.com/panjf2000/gnet/v2"
	"sync"
)

type Entity struct {
	handler  gnet.EventHandler
	lock     sync.RWMutex
	shutdown bool
	host     string
	port     string
	options  gnet.Options
}

func NewEntity(handler gnet.EventHandler, host string, port string, options gnet.Options) *Entity {
	return &Entity{handler: handler, host: host, port: port, options: options}
}

func (e *Entity) Start() {
	addr := "tcp://" + e.host + ":" + e.port
	go func() {
		err := gnet.Run(e.handler, addr, gnet.WithOptions(e.options))
		if err != nil {
			panic(any(err))
		}
	}()
}

func (e *Entity) Stop() {

}
