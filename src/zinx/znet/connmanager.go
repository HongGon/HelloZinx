package znet

import (
	"errors"
	"fmt"
	"sync"
	"zinx/ziface"
)

/*
	connect manage module
*/

type ConnManager struct {
	// manage the conn info
	connections map[uint32]ziface.IConnection
	// read-write lock
	connLock sync.RWMutex
}

/*
	Create a link manage
*/

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}

// add a conn
func (connMgr *ConnManager) Add(conn ziface.IConnection) {
	// protect shared resources map, lock
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()
	//  add conn to ConnManager
	connMgr.connections[conn.GetConnID()] = conn
	fmt.Println("connection add to ConnManager successfully: conn num = ", connMgr.Len())
}

// del the conn
func (connMgr *ConnManager) Remove(conn ziface.IConnection) {
	// protect the shared resource Map
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()
	// del the conn info
	delete(connMgr.connections,conn.GetConnID())
	fmt.Println("connection Remove ConnID=",conn.GetConnID(),"successfully: conn num = ",connMgr.Len())
}

//  obtain the ID by ConnID
func (connMgr *ConnManager) Get(connID uint32) (ziface.IConnection, error) {
	connMgr.connLock.RLock()
	defer connMgr.connLock.RUnlock()
	if conn, ok := connMgr.connections[connID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not found")
	}
}

// obtain the current conn
func (connMgr *ConnManager) Len() int {
	return len(connMgr.connections)
}

// clear and stop all conns
func (connMgr *ConnManager) ClearConn() {
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()
	// stop and del all conn info
	for connID, conn := range connMgr.connections {
		// stop
		conn.Stop()
		// del
		delete(connMgr.connections,connID)
	}

	fmt.Println("Clear All Connections successfully: conn num = ", connMgr.Len())

}




















