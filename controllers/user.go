package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfo(ctx *gin.Context) {
	ReturnSuccess(ctx, http.StatusOK, "user info", "111111111", 1)
}

func GetUserList(ctx *gin.Context) {
	param := make(map[string]interface{})
	err := ctx.BindJSON(&param)
	if err == nil {
		ReturnSuccess(ctx, http.StatusOK, param["name"], param["id"], 1)
	} else {
		ReturnError(ctx, http.StatusNotFound, gin.H{"err": err})
	}
}
