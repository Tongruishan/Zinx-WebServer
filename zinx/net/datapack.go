package net

import (
	"ZinxHouse/Zinx-WebServer/zinx/ziface"
	"bytes"
	"encoding/binary"
	"fmt"
)

type DataPack struct {

}

func NewDataPack()*DataPack{

	return &DataPack{
	}
}

func(dp *DataPack) GetHeadLen()uint32{
	return 8
}

//包装
func(dp *DataPack) MsgPack(msg ziface.IMessage)([]byte,error){

	//创建新的缓冲区
	databuffer:=bytes.NewBuffer([]byte{})

	//将message的内容存到buffer中
	err:=binary.Write(databuffer,binary.LittleEndian,msg.GetMsgLen())
	if err!=nil{
		fmt.Println("binary Write Len err:",err)
		return nil,err
	}

	err=binary.Write(databuffer,binary.LittleEndian,msg.GetMsgId())
	if err!=nil{
		fmt.Println("binary Write Id err:",err)
		return nil,err
	}

	err=binary.Write(databuffer,binary.LittleEndian,msg.GetMsgData())
	if err!=nil{
		fmt.Println("binary Write data err:",err)
		return nil,err
	}

	//
	return databuffer.Bytes(),nil



}

//解装
func(dp *DataPack) MsgUnPack(data []byte)(ziface.IMessage,error){

	msgHead:=&Message{}

	dataBuf:=bytes.NewBuffer(data)

	err:=binary.Read(dataBuf,binary.LittleEndian,&msgHead.MsgLen)
	if err!=nil{
		fmt.Println("binary Read MsgLen err:",err)
		return nil,err
	}

	err=binary.Read(dataBuf,binary.LittleEndian,&msgHead.MsgId)
	if err!=nil{
		fmt.Println("binary Read MsgLen err:",err)
		return nil,err
	}


	return msgHead,nil

}
