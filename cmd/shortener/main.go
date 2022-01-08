package main

import (
	"log"
	"net/http"
	"os"

	"github.com/EMus88/go-musthave-shortener-tpl/internal/app/service"
	"github.com/EMus88/go-musthave-shortener-tpl/internal/handler"
	"github.com/EMus88/go-musthave-shortener-tpl/internal/repository"
	"github.com/EMus88/go-musthave-shortener-tpl/internal/repository/models/file"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	var model file.Model

	r := repository.NewStorage()
	s := service.NewService(r, &model)
	h := handler.NewHandler(s)

	//load env variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	//load data from file
	s.LoadFromFile()

	//start server
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/:id", h.HandlerGet)
	router.POST("/", h.HandlerPostText)
	router.POST("/api/shorten", h.HandlerPostJSON)
	router.NoRoute(func(c *gin.Context) { c.String(http.StatusBadRequest, "Not allowed requset") })

	serverAddress := os.Getenv("SERVER_ADDRESS")
	router.Run(serverAddress)
}
