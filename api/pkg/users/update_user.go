package users

import (
    "net/http"

    "github.com/b0nbon1/debt-roundup/pkg/common/models"
    "github.com/gin-gonic/gin"
)

type UpdateUserRequestBody struct {
    Firstname       string `json:"firstanme"`
    Lastname        string `json:"lastname"`
    Email           string `json:"email"`
	Mobileno        string `json:"mobileno"`
}

func (h handler) UpdateUser(ctx *gin.Context) {
    id := ctx.Param("id")
    body := UpdateUserRequestBody{}

    if err := ctx.BindJSON(&body); err != nil {
        ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }

    var user models.User

    if result := h.DB.First(&user, id); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    user.Firstname = body.Firstname
    user.Lastname = body.Lastname
    user.Email = body.Email
    user.Mobileno = body.Mobileno

    h.DB.Save(&user)

    ctx.JSON(http.StatusOK, &user)
}

