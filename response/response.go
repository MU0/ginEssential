package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, htttpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(htttpStatus, gin.H{"code": code, "data": data, "msg": msg})

}

func Success(ctx *gin.Context, code int, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)

}

func Fail(ctx *gin.Context, code int, data gin.H, msg string) {
	Response(ctx, http.StatusBadRequest, 200, data, msg)

}