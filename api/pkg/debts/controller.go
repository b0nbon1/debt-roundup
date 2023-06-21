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

    routes := router.Group("/api/v1/debts")
    routes.POST("/", h.AddDebt)
    routes.GET("/", h.GetDebts)
    routes.GET("/:id", h.GetDebt)
    routes.PUT("/:id", h.UpdateDebt)
    routes.DELETE("/:id", h.DeleteDebt)
}
