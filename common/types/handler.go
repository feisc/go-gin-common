package types

import "github.com/gin-gonic/gin"

type Handler func(c *gin.Context) error

type ControllerActionResult string

const (
	CtrlActSuccess ControllerActionResult = "success"
	CtrlActFail    ControllerActionResult = "fail"
)
