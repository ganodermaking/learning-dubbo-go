package dubbo

import (
	"context"
)

type HelloRequest struct {
	Name string
}

func (h HelloRequest) JavaClassName() string {
	return "com.ganodermaking.learningapi.dto.HelloRequest"
}

type HelloResponse struct {
	Name string
}

func (h HelloResponse) JavaClassName() string {
	return "com.ganodermaking.learningapi.dto.HelloResponse"
}

type IGreeter struct {
	SayHello func(ctx context.Context, request HelloRequest) (HelloResponse, error) `dubbo:"sayHello"`
}

func NewGreeter() *IGreeter {
	iGreeter := new(IGreeter)
	return iGreeter
}

func (p *IGreeter) Reference() string {
	return "IGreeter"
}
