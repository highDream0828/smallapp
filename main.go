package main

import (
    
    "github.com/labstack/echo/v4"
	"github.com/highdream0828/smallapp/apihandlers/auth"
	"github.com/highdream0828/smallapp/data/migrations"
	"github.com/highdream0828/smallapp/data/dbspeeds"
)

func main() {
    e := echo.New()
    dbspeeds.Connect()
    migrations.Up()
    e.POST("/users", auth.Register)
    e.POST("/users/login", auth.Login)
    e.Logger.Fatal(e.Start(":1323"))
}