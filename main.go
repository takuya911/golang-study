package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/takuya911/golang-study/infra"
)

func main() {

	// mysql connect
	db, err := infra.NewGormDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	//
	e := echo.New()
	log.Fatal(e.Start(":8080"))
}
