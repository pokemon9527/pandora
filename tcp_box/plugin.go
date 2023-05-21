package tcp_box

import (
	"github.com/panjf2000/gnet/v2"
	"github.com/pokemon9527/pandora/tcp_box/define"
)

type Plugin interface {
	DecodeMsg(conn gnet.Conn, sess define.Session) define.PackageStatus
}
