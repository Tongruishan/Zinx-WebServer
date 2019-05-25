package ziface

type IDataPack interface {

	GetHeadLen()uint32

	MsgPack(message IMessage)([]byte,error)

	MsgUnPack([]byte)(IMessage,error)
}

