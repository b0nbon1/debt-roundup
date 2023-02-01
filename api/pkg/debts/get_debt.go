package debts

import (
    "net/http"

    "github.com/b0nbon1/debt-roundup/pkg/common/models"
    "github.com/gin-gonic/gin"
)

func (h handler) GetDebt(ctx *gin.Context) {
    id := ctx.Param("id")

    var debt models.Debt

    if result := h.DB.First(&debt, id); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusOK, &debt)
}

