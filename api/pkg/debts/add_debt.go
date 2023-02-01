package debts

import (
    "net/http"

    "github.com/b0nbon1/debt-roundup/pkg/common/models"
    "github.com/gin-gonic/gin"
)

type AddDebtRequestBody struct {
    Title       string `json:"title"`
    Lender      string `json:"lender"`
    Description string `json:"description"`
}

func (h handler) AddDebt(ctx *gin.Context) {
    body := AddDebtRequestBody{}

    if err := ctx.BindJSON(&body); err != nil {
        ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }

    var debt models.Debt

    debt.Title = body.Title
    debt.Lender = body.Lender
    debt.Description = body.Description

    if result := h.DB.Create(&debt); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusCreated, &debt)
}

