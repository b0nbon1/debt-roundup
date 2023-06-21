package users

import (
    "net/http"

    "github.com/b0nbon1/debt-roundup/pkg/common/models"
    "github.com/gin-gonic/gin"
)

func (h handler) GetUser(ctx *gin.Context) {
    id := ctx.Param("id")

    var user models.User

    if result := h.DB.Preload("Loans").Preload("Assets").First(&user, id); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusOK, &user)
}

