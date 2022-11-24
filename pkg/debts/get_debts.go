package debts

import (
    "net/http"

    "github.com/b0nbon1/debt-roundup/pkg/common/models"
    "github.com/gin-gonic/gin"
)

func (h handler) GetDebts(ctx *gin.Context) {
    var debts []models.Debt

    if result := h.DB.Find(&debts); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusOK, &debts)
}

