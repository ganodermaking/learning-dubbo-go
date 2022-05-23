package routers

import (
	"github.com/gin-contrib/pprof"
	_ "github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmgin"
	v1Ctrl "learning-go/cmd/http/controllers/v1"
)

// Init ...
func Init() *gin.Engine {
	defer logrus.Debug("router is ready")
	defer logrus.Debug("initialized gin engine")
	r := gin.New()
	r.HandleMethodNotAllowed = true // 开启 405 错误 handler
	r.RemoveExtraSlash = true       // 去掉 uri 末尾的斜杠 '/'
	registerMiddleware(r)
	registerRoutes(r)
	return r
}

func registerMiddleware(engine *gin.Engine) {
	defer logrus.Debug("registered route middleware")
	defer engine.Use(gin.Recovery())
	engine.Use(apmgin.Middleware(engine))
}

func registerRoutes(engine *gin.Engine) {
	defer logrus.Debug("registered route paths")
	pprof.Register(engine) // pprof for gin
	registerV1(engine)
}

func registerV1(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")
	{
		v1.GET("/hello", v1Ctrl.Hello)
	}
}
