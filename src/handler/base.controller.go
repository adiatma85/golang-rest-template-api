package handler

import (
	"net/http"

	"github.com/adiatma85/go-tutorial-gorm/src/helper"
	"github.com/gin-gonic/gin"
)

type BaseController interface {
	Base(ctx *gin.Context)
}

type baseController struct {
}

func NewBaseHandler() BaseController {
	return &baseController{}
}

func (c *baseController) Base(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, helper.BuildSuccessResponse("success", nil))
}
