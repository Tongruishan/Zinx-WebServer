package config

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type Globle struct {
	Name string
	Host string
	Port int
	Version string
	MaxLen int
	WorkerPoolSize uint32
	MaxWorkerTaskLen uint32
	MaxConn uint32
}

var GlobleConf *Globle

func(g *Globle)LoadConfig(){
	data,err:=ioutil.ReadFile("conf/zinx.json")
	if err!=nil{
		fmt.Println("LoadConfig err",err)
		return
	}

	err=json.Unmarshal(data,&GlobleConf)
	if err!=nil{
		fmt.Println("LoadConfig Unmarshal err",err)
		return
	}

}

func init(){

	GlobleConfInit:=&Globle{
		Name:"zinxSever",
		Host:"0.0.0.0",
		Port:8999,
		Version:"Zinx V0.4",
		MaxLen:4096,
		WorkerPoolSize:10,
		MaxWorkerTaskLen:4096,
		MaxConn:1000,
	}

	GlobleConfInit.LoadConfig()

}

