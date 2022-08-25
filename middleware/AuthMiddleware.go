package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"xjtu.teach/ginEssential/common"
	"xjtu.teach/ginEssential/model"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//获取authorization header
		tokenStrng := ctx.GetHeader("Authorization")

		//validate token formate
		if tokenStrng == "" || !strings.HasPrefix(tokenStrng, "Bearer ") {
			ctx.JSON(401, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		tokenStrng = tokenStrng[7:]

		token, claims, err := common.ParseToken(tokenStrng)

		if err != nil || !token.Valid {
			ctx.JSON(401, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		// 验证通过后获取claim 中的userID
		userId := claims.Userid
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		// 验证用户是否存在
		if user.ID == 0 {
			ctx.JSON(401, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		//用户存在 将user的信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
