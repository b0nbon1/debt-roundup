package debts

import (
    "net/http"
    "github.com/b0nbon1/debt-roundup/pkg/common/models"
    "github.com/gin-gonic/gin"
)

type AddDebtRequestBody struct {
    Title       string `json:"title"`
    Loaner      int `json:"loaner"`
    Description string `json:"description"`
    Loanee      int `json:"loanee"`
    Amount      float64 `json:"amount"`
}

func (h handler) AddDebt(ctx *gin.Context) {
    body := AddDebtRequestBody{}

    if err := ctx.BindJSON(&body); err != nil {
        ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }

    var debt models.Debt

    debt.Title = body.Title
    debt.Loaner = body.Loaner
    debt.Loanee = body.Loanee
    debt.Description = body.Description
    debt.Amount = body.Amount

    if result := h.DB.Create(&debt); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusCreated, &debt)
}

