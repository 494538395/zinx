package net

import (
	"errors"
	"fmt"
	"github.com/494538395/zinx/iface"
	logger "github.com/494538395/zinx/log"
	"sync"
)

type ConnManager struct {
	// Conns 连接集合 connID => conn
	Conns map[uint32]iface.IConnection
	lock  sync.RWMutex
}

func NewConnManager() iface.IConnManager {
	return &ConnManager{
		Conns: make(map[uint32]iface.IConnection),
	}
}

func (cm *ConnManager) AddConn(conn iface.IConnection) {
	cm.lock.Lock()
	defer cm.lock.Unlock()

	cm.Conns[uint32(conn.GetConnID())] = conn
	logger.Info("[ConnManager][AddConn] add conn,connID:", conn.GetConnID())

}

func (cm *ConnManager) RemoveConn(conn iface.IConnection) {
	cm.lock.Lock()
	defer cm.lock.Unlock()

	delete(cm.Conns, uint32(conn.GetConnID()))
	logger.Info("[ConnManager][RemoveConn] remove conn,connID:", conn.GetConnID())

}

func (cm *ConnManager) GetConn(connID uint32) (iface.IConnection, error) {
	cm.lock.RLock()
	defer cm.lock.RUnlock()

	c, found := cm.Conns[connID]
	if !found {
		return nil, errors.New(fmt.Sprintf("conn not found,connID:%d", connID))
	}
	return c, nil
}

func (cm *ConnManager) GetConnCnt() int {
	return len(cm.Conns)
}

func (cm *ConnManager) RemoveAllConn() {
	cm.lock.Lock()
	defer cm.lock.Unlock()

	for _, conn := range cm.Conns {
		delete(cm.Conns, uint32(conn.GetConnID()))
		conn.Stop()
	}

	logger.Info("[ConnManager][RemoveAllConn] all conn remove and closed")
}
