package main

import "sync"
//创建GetInstance对象
type Singleton struct{}

var singleton *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func(){
		singleton = &Singleton{}
	})
	return singleton
}

func main(){
	GetInstance()
}