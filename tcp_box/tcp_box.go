package tcp_box

import "errors"

type Manager struct {
	AppName  string
	Entities map[string]*Entity
}

func NewTcpManager(name string) *Manager {
	return &Manager{AppName: name}
}

func (ma *Manager) AddEntity(Name string, entity *Entity) error {
	if _, ok := ma.Entities[Name]; ok {
		return errors.New("服务已存在")
	}
	entity.Start()
	ma.Entities[Name] = entity
	return nil
}

func (ma *Manager) StopEntity(Name string) {
	if _, ok := ma.Entities[Name]; !ok {
		return
	}
	ma.Entities[Name].Stop()
}
