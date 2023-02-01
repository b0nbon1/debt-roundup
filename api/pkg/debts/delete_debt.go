package debts

import (
    "net/http"

    "github.com/b0nbon1/debt-roundup/pkg/common/models"
    "github.com/gin-gonic/gin"
)

func (h handler) DeleteDebt(ctx *gin.Context) {
    id := ctx.Param("id")

    var debt models.Debt

    if result := h.DB.First(&debt, id); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    h.DB.Delete(&debt)

    ctx.Status(http.StatusOK)
}
