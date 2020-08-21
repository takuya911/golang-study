package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/takuya911/golang-study/handler"
	"github.com/takuya911/golang-study/infra"
	"github.com/takuya911/golang-study/usecase"
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
	handler.NewUserHandler(
		e,
		usecase.NewUserUsecase(db),
	)

	log.Fatal(e.Start(":8080"))
}
