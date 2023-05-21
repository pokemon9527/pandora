package tcp_box

import (
	"github.com/panjf2000/gnet/v2"
	"github.com/pokemon9527/pandora/tcp_box/define"
	"sync"
)

type sessionDict = map[gnet.Conn]*define.Session

type Handle struct {
	*gnet.BuiltinEventEngine
	SdLock       sync.RWMutex
	Sd           sessionDict
	HandlePlugin Plugin
}

func NewHandle() *Handle {
	return &Handle{}
}
