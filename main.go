package main

import (
	"fmt"
	"github.com/micro/go-micro"
)

func main () {
	service := micro.NewService(micro.Name("walkie.service.user"))

	service.Init()
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}