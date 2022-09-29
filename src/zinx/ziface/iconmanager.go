package ziface

/*
	Connect to manage layer
*/

type IConnManager interface {
	// add a conn
	Add(conn IConnection)
	// del a conn
	Remove(conn IConnection)
	// obtain the conn by ConnID
	Get(connID uint32) (IConnection, error)
	// obtain the current conn
	Len()	int
	// del and stop all conns
	ClearConn()
}




