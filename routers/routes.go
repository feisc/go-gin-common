package routers

import (
	"github.com/gin-gonic/gin"
	"gitlab/go-gin/common/controller"
	"gitlab/go-gin/common/logger"
	"gitlab/go-gin/common/middleware"
	"net/http"
)

func SetupRouter(mode string, isHttps bool, httpsPort int) *gin.Engine {
	// release mode
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	// new a engine
	r := gin.New()
	// logger, recover and request id
	r.Use(logger.GinLogger(),
		logger.GinRecovery(true),
		middleware.RequestIDHandler())
	// tls handler
	if isHttps {
		r.Use(middleware.TLSHandler(httpsPort))
	}

	r.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// /api/v1
	api := r.Group("/api")
	v1Group := api.Group("/v1")

	// init route
	initRoute_v1(v1Group)

	// 404
	r.NoRoute(controller.NotFound)

	return r
}

func initRoute_v1(g *gin.RouterGroup) {
	g.POST("/mysql/batchExecTest", controller.ErrorWrapper(controller.BatchMysqlHandlerTest))
}
