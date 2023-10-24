package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfo(ctx *gin.Context) {
	id := ctx.Param("id")
	name := ctx.Param("name")
	ReturnSuccess(ctx, http.StatusOK, name, id, 1)
}

func GetUserList(ctx *gin.Context) {
	ReturnError(ctx, http.StatusBadRequest, "错误400")
}
