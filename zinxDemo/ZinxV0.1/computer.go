package main

import "fmt"

//

type Card interface {
	display()
}

type Memory interface {
	storage()
}

type CPU interface {
	calculate()
}

type computer struct {
	card Card
	mem Memory
	cpu CPU
}

func NewComputer(card Card,mem Memory,cpu CPU)*computer{

	return &computer{
		card:card,
		mem:mem,
		cpu:cpu,
	}

}

func (this *computer)Work() {
	this.card.display()
	this.mem.storage()
	this.cpu.calculate()
}


//
type IntelMemory struct {
	Memory
}
func(this *IntelMemory)storage(){
	fmt.Println("this is IntelMemory")
}


type IntelCard struct {
	Card
}
func(this *IntelCard)display(){
	fmt.Println("this is IntelCard")
}

type IntelCPU struct {
	CPU
}
func(this *IntelCPU)calculate(){
	fmt.Println("this is IntelCPU")
}

//
func main(){
	computer:=NewComputer(&IntelCard{},&IntelMemory{},&IntelCPU{})
	computer.Work()
}

