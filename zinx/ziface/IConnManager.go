package ziface

type IConnManager interface {
	Add(conn IConnection)
	Remove(connId uint32)
	Get(connId uint32)(IConnection,error)
	Len()int
	ClearConn()
}

