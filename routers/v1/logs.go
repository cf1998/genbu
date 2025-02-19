/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package v1

import (
	"genbu/controllers/operationLogs"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func InitLogRouters(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	{
		r.Use(authMiddleware.MiddlewareFunc())
		r.GET("/log", operationLogs.GetOperationLogList)
	}
	return r
}
