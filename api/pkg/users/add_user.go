package users

import (
    "net/http"
    "github.com/b0nbon1/debt-roundup/pkg/common/models"
    "github.com/gin-gonic/gin"
)

type AddUserRequestBody struct {
    Firstname       string `json:"firstanme"`
    Lastname        string `json:"lastname"`
    Email           string `json:"email"`
	Mobileno        string `json:"mobileno"`
}

func (h handler) AddUser(ctx *gin.Context) {
    body := AddUserRequestBody{}

    if err := ctx.BindJSON(&body); err != nil {
        ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }

    var user models.User

    user.Firstname = body.Firstname
    user.Lastname = body.Lastname
    user.Email = body.Email
    user.Mobileno = body.Mobileno

    if result := h.DB.Create(&user); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusCreated, &user)
}

