package routers

import "github.com/gin-gonic/gin"

import "github.com/Manu-Opensource/CMS-Backend/routers/api"

func RouterInit() *gin.Engine {
    r := gin.New()

    r.GET("/api/ping", api.Ping)

    r.Run()

    return r
}
