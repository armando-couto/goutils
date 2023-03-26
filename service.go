package goutils

import (
	"fmt"
	"github.com/kardianos/service"
)

type program struct{}

func (p program) Start(s service.Service) error {
	fmt.Println(s.String() + " started")
	return nil
}

func (p program) Stop(s service.Service) error {
	fmt.Println(s.String() + " stopped")
	return nil
}

func Execute(serviceConfig *service.Config) {
	prg := &program{}
	s, err := service.New(prg, serviceConfig)
	if err != nil {
		fmt.Println("Não pode criar o serviço: " + err.Error())
	}
	err = s.Run()
	if err != nil {
		fmt.Println("Não é possível iniciar o serviço: " + err.Error())
	}
}
