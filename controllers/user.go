package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserList struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
}

func GetUserInfo(ctx *gin.Context) {
	ReturnSuccess(ctx, http.StatusOK, "user info", "111111111", 1)
}

func GetUserList(ctx *gin.Context) {
	userlist := &UserList{}
	err := ctx.BindJSON(&userlist)
	if err == nil {
		ReturnSuccess(ctx, http.StatusOK, userlist.Name, userlist.ID, 1)
	} else {
		ReturnError(ctx, http.StatusNotFound, gin.H{"err": err})
	}
}
