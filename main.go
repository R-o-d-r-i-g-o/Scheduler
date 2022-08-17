package main

import (
	"fmt"
	"scheduler/configs"
	"scheduler/router"
	"scheduler/server"

	"scheduler/database"
	model "scheduler/database/entity"
	"time"
)

func init() {

	configs.Load()
	database.StartDatabase()
	model.Handler(database.GetGormDB())
}

func main() {

	currentServer := server.CreateServer()
	router.Avaible(currentServer.GetServerEngine())
	currentServer.GetServerEngine().Run(

		configs.Server.HOST + ":" +
			configs.Server.PORT,
	)

	currentTime := time.Now()

	format1 := "2006-01-02"
	format2 := "15:04:05"

	timeTest1, _ := time.Parse(format1, currentTime.Format(format1))
	timeTest2, _ := time.Parse(format2, currentTime.Format(format2))

	fmt.Println(currentTime.Format("2006-01-02"))
	fmt.Println(currentTime.Format("15:04:05"))
	fmt.Println(timeTest1)
	fmt.Println(timeTest2)
}

// 01   -> month with zero prefix
// 02   -> day with zero prefix
// 06   -> year (last two digits)
// 15   -> hour (24h based)
// 04   -> minutes with zero prefix
// 05   -> seconds with zero prefix
// 2006 -> long year
