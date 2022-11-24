package debts

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type handler struct {
    DB *gorm.DB
}


func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
    h := &handler{
        DB: db,
    }

    routes := router.Group("/debts")
    routes.POST("/", h.AddDebt)
    routes.GET("/", h.GetDebts)
    routes.GET("/:id", h.GetDebt)
    routes.PATCH("/:id", h.UpdateDebt)
    routes.DELETE("/:id", h.DeleteDebt)
}
