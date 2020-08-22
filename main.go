package main

import (
	"log"

	"github.com/labstack/echo"
	_userHandler "github.com/takuya911/golang-study/handler"
	_infra "github.com/takuya911/golang-study/infra"
	_userRepo "github.com/takuya911/golang-study/repository"
	_userUsecase "github.com/takuya911/golang-study/usecase"
)

func main() {

	// mysql connect
	db, err := _infra.NewGormDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()

	// services
	userRepo := _userRepo.NewUserRepository(db)
	userUsecase := _userUsecase.NewUserUsecase(userRepo)
	_userHandler.NewUserHandler(e, userUsecase)

	log.Fatal(e.Start(":8080"))
}
