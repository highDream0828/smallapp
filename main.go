package main

import (
    "net/http"
    
    "github.com/labstack/echo/v4"
	"github.com/smallapp/apihandlers/auth"
)

func main() {
    e := echo.New()
    e.GET("/", auth.Signup)
    e.Logger.Fatal(e.Start(":1323"))
}