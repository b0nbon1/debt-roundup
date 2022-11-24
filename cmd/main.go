package main

import (
	"fmt"

	"github.com/b0nbon1/debt-roundup/pkg/common/db"
	"github.com/b0nbon1/debt-roundup/pkg/debts"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
    viper.SetConfigFile(".env")
    viper.ReadInConfig()

    // add env variables as needed
    port := viper.Get("PORT").(string)
    dbUrl := viper.Get("DB_URL").(string)

    fmt.Println(port, dbUrl)

    router := gin.Default()
    dbHandler := db.Init(dbUrl)

    debts.RegisterRoutes(router, dbHandler)

    router.GET("/", func(ctx *gin.Context) {
    ctx.JSON(200, gin.H{
      "message" : "Welcome to debt rounderup",
    })
  })
  router.Run(port)
}

