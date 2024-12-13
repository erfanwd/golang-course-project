package main

import (
	"github.com/erfanwd/golang-course-project/api"
	"github.com/erfanwd/golang-course-project/config"
	"github.com/erfanwd/golang-course-project/data/cache"
	"github.com/erfanwd/golang-course-project/data/db"
	"github.com/erfanwd/golang-course-project/data/db/migrations"
	"github.com/erfanwd/golang-course-project/pkg/logging"
)



func main(){
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)
	cache.InitRedis(cfg)
	defer cache.CloseRedis()

	err := db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil{
		logger.Fatal(logging.Postgres, logging.StartUp, err.Error(), nil)
	}
	migrations.Up1()
	api.InitialServer(cfg)
	
}