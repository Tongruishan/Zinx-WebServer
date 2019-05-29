package main

import (
	"ZinxHouse/Zinx-WebServer/zinxDemo/protoDemo/pb"
	"github.com/golang/protobuf/proto"
	"fmt"
)

func main(){

	Person:=&pb.Person{
		Name:"天王盖地虎",
		Age:32,
		Emails:[]string{"41626709@qq.com","456789@163.com"},
		Phones:[]*pb.PhoneNumber{
			&pb.PhoneNumber{
				Number:"148487878787878",
			},
			&pb.PhoneNumber{
				Number:"5998989",
			},

		},
		//Data:&pb.Person_School{
		//	School:"密云二中",
		//},
		Data:&pb.Person_Score{
			Score:"512",
		},

	}

	data,err:=proto.Marshal(Person)
	if err!=nil{
		fmt.Println("Marshal err",err)
		return
	}

	fmt.Println("编译前",Person.GetName(),Person.GetAge(),Person.GetEmails())

	//编译后
	NewPerson:=&pb.Person{}
	err=proto.Unmarshal(data,NewPerson)
	if err!=nil{
		fmt.Println("Unmarshal err",err)
		return
	}

	fmt.Println("编译后",NewPerson.GetName(),NewPerson.GetAge(),NewPerson.Emails)
	fmt.Println("电话",NewPerson.GetPhones())
	fmt.Println("Data",NewPerson.GetSchool())
	fmt.Println("Data",NewPerson.GetScore())

}
