package users

import (
    "net/http"

    "github.com/b0nbon1/debt-roundup/pkg/common/models"
    "github.com/gin-gonic/gin"
)

func (h handler) GetUsers(ctx *gin.Context) {
    var users []models.User

    if result := h.DB.Preload("Loans").Preload("Assets").Find(&users); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusOK, &users)
}

