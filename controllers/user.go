package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfo(ctx *gin.Context) {
	ReturnSuccess(ctx, http.StatusOK, "user info", "111111111", 1)
}

func GetUserList(ctx *gin.Context) {
	id := ctx.PostForm("id")
	name := ctx.DefaultPostForm("name", "neo")
	ReturnSuccess(ctx, http.StatusOK, name, id, 1)
}
