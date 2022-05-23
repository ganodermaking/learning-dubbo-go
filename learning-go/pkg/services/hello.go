package services

import (
	"context"
	dubboProvider "learning-go/initializers/dubbo"
	"learning-go/pkg/dubbo"
)

type HelloService struct {
}

func NewHelloService() *HelloService {
	return &HelloService{}
}

func (service *HelloService) Hello(ctx context.Context) (*dubbo.HelloResponse, error) {
	request := dubbo.HelloRequest{}
	request.Name = "Dubbo"
	response, err := dubboProvider.Greeter.SayHello(ctx, request)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
