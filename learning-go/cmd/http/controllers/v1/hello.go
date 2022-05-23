package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"learning-go/pkg/services"
)

func Hello(c *gin.Context) {
	response, err := services.NewHelloService().Hello(c)
	if err != nil {
		logrus.Error(err)
		return
	}

	c.JSON(200, gin.H{
		"message": response,
	})
}
