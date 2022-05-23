package main

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/config"
	hessian "github.com/apache/dubbo-go-hessian2"
	"github.com/sirupsen/logrus"
	"learning-go/pkg/dubbo"
)

func main()  {
	hessian.RegisterPOJO(&dubbo.HelloRequest{})
	hessian.RegisterPOJO(&dubbo.HelloResponse{})

	consumerConfig := *config.GetConsumerConfig()
	consumerConfig.References = make(map[string]*config.ReferenceConfig)
	referenceConfig := config.ReferenceConfig{
		URL: "dubbo://127.0.0.1:22880",
		Cluster: "failver",
		InterfaceName: "com.ganodermaking.learningapi.api.IGreeter",
	}
	consumerConfig.References["IGreeter"] = &referenceConfig
	config.SetConsumerConfig(consumerConfig)

	config.Load()

	request := dubbo.HelloRequest{}
	response, err := dubbo.NewGreeter().SayHello(context.Background(), request)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info(response)
}
