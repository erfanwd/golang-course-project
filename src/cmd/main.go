package main

import (
	"github.com/erfanwd/golang-course-project/api"
	"github.com/erfanwd/golang-course-project/config"
	"github.com/erfanwd/golang-course-project/data/cache"
	"github.com/erfanwd/golang-course-project/data/db"
	"github.com/erfanwd/golang-course-project/data/db/migrations"
	"github.com/erfanwd/golang-course-project/pkg/logging"
)

// @title Golang Web API
// @version 1.0
// @description golang web api
// @BasePath /api
// @host localhost:8080
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main(){
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)
	cache.InitRedis(cfg)
	defer cache.CloseRedis()

	gormDB, err := db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil{
		logger.Fatal(logging.Postgres, logging.StartUp, err.Error(), nil)
	}
	migrations.Up1()
	api.InitialServer(cfg, gormDB)
	
}