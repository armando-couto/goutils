package goutils

import (
	"fmt"
	"github.com/kardianos/service"
)

type program struct{}

func (p program) Start(s service.Service) error {
	fmt.Println(s.String() + " Rodando!")
	return nil
}

func (p program) Stop(s service.Service) error {
	fmt.Println(s.String() + " Parado!")
	return nil
}

func ExecuteService(serviceConfig *service.Config) {
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
