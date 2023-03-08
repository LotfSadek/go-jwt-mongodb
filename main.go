package main

import (
	"golang-jwt/controllers"
	"golang-jwt/repository"
	"golang-jwt/router"
	"golang-jwt/usecases"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	ar := repository.NewAuthRepo()
	ur := repository.NewUserRepo()
	uc := usecases.NewBaseRepo(ar, ur)
	ah := controllers.NewAuthHandler(uc)
	uh := controllers.NewUserHandler(uc)
	router.StartGinRouter(port, ah, uh)
}
