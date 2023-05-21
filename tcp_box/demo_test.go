package tcp_box

import (
	"github.com/panjf2000/gnet/v2"
	"testing"
)

func TestNew(t *testing.T) {
	NewTcpManager("云平台").
		AddEntity("文件服务1",
			NewEntity(
				NewHandle(),
				"127.0.0.1",
				1600,
				gnet.Options{}))
}
