package debts

import (
    "net/http"

    "github.com/b0nbon1/debt-roundup/pkg/common/models"
    "github.com/gin-gonic/gin"
)

type UpdateDebtRequestBody struct {
    Title       string `json:"title"`
    Lender      string `json:"lender"`
    Description string `json:"description"`
}

func (h handler) UpdateDebt(ctx *gin.Context) {
    id := ctx.Param("id")
    body := UpdateDebtRequestBody{}

    if err := ctx.BindJSON(&body); err != nil {
        ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }

    var debt models.Debt

    if result := h.DB.First(&debt, id); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    debt.Title = body.Title
    debt.Lender = body.Lender
    debt.Description = body.Description

    h.DB.Save(&debt)

    ctx.JSON(http.StatusOK, &debt)
}

