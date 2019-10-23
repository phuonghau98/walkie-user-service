package main

import (
	"fmt"
	"github.com/micro/go-micro"
	userPB "github.com/phuonghau98/walkie-user-service/proto"
)

func main () {
	service := micro.NewService(micro.Name("walkie.service.user"))
	service.Init()
	userPB.RegisterUserServiceHandler(service.Server(), &Handler{})
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}