package main

import (
	"log"

	"github.com/labstack/echo"
	_infra "github.com/takuya911/golang-study/infra"
	_userHandler "github.com/takuya911/golang-study/user/handler"
	_userRepo "github.com/takuya911/golang-study/user/repository"
	_userUsecase "github.com/takuya911/golang-study/user/usecase"
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
	// user services
	userRepo := _userRepo.NewUserRepository(db)
	userUsecase := _userUsecase.NewUserUsecase(userRepo)
	_userHandler.NewUserHandler(e, userUsecase)

	log.Fatal(e.Start(":8080"))
}
