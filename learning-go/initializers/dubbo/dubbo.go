package dubbo

import (
	"learning-go/pkg/dubbo"

	"dubbo.apache.org/dubbo-go/v3/config"
	hessian "github.com/apache/dubbo-go-hessian2"

	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

var Greeter *dubbo.IGreeter

func Init() {
	Greeter = new(dubbo.IGreeter)
	hessian.RegisterPOJO(&dubbo.HelloRequest{})
	hessian.RegisterPOJO(&dubbo.HelloResponse{})
	config.SetConsumerService(Greeter)
	config.Load()
}
