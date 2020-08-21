package main

import (
	"log"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	log.Fatal(e.Start(":8080"))
}
