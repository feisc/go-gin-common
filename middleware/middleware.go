package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/unrolled/secure"
	"gitlab/go-gin/common/util"
	"go.uber.org/zap"
)

func TLSHandler(port int) gin.HandlerFunc {
	return func(c *gin.Context) {
		secMidd := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     fmt.Sprintf(":%d", port),
		})
		err := secMidd.Process(c.Writer, c.Request)
		if err != nil {
			zap.L().Error("tls handler", zap.Error(err))
			return
		}

		c.Next()
	}
}

func RequestIDHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		rid := uuid.New().String()
		c.Set(util.RidFlag, rid)
		c.Next()
	}
}
