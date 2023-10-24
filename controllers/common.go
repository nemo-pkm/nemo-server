package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JSONStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}

type JSONErrStruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
}

// return 200
func ReturnSuccess(ctx *gin.Context, code int, msg interface{}, data interface{}, count int64) {
	json := &JSONStruct{
		Code:  code,
		Msg:   msg,
		Data:  data,
		Count: count,
	}
	ctx.JSON(http.StatusOK, json)
}

// return 400
func ReturnError(ctx *gin.Context, code int, msg interface{}) {
	json := &JSONErrStruct{
		Code: code,
		Msg:  msg,
	}
	ctx.JSON(http.StatusOK, json)
}
