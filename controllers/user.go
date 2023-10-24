package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfo(ctx *gin.Context) {
	ReturnSuccess(ctx, http.StatusOK, "succes", "getuser", 1)
}

func GetUserList(ctx *gin.Context) {
	ReturnError(ctx, http.StatusBadRequest, "错误400")
}
