package znet

import (
	"errors"
	"fmt"
	"github.com/lorenzoyu2000/zinx/ziface"
	"sync"
)

type ConnManager struct {
	connections map[uint32]ziface.IConnection
	connLock    sync.RWMutex
}

func (c *ConnManager) AddConn(conn ziface.IConnection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	c.connections[conn.GetConnID()] = conn
	fmt.Println("Add Connection succeed, ConnID = ", conn.GetConnID())
}

func (c *ConnManager) RemoveConn(conn ziface.IConnection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	delete(c.connections, conn.GetConnID())
	fmt.Println("Remove Connection succeed, ConnID = ", conn.GetConnID())
}

func (c *ConnManager) GetConn(connID uint32) (ziface.IConnection, error) {
	c.connLock.RLock()
	defer c.connLock.RUnlock()

	conn, ok := c.connections[connID]
	if ok {
		return conn, nil
	}
	return nil, errors.New("connection not found")
}

func (c *ConnManager) Size() int {
	return len(c.connections)
}

func (c *ConnManager) ClearAllConn() {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	for connID, conn := range c.connections {
		conn.Stop()
		delete(c.connections, connID)
	}
	fmt.Println("Clear All Conn Succeed")
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}
