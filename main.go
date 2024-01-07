package main

import (
	"duongdx/example/initializers"
	"duongdx/example/router"
	"fmt"
	"log"

	"github.com/labstack/echo"
)

func main() {
	env, err := initializers.LoadConfig()

	if err != nil {
		log.Fatal("Loading environment error")
		return
	}

	sql := initializers.SQL{
		Host:         env.DBHost,
		Port:         env.DBPort,
		UserName:     env.DBUserName,
		Password:     env.DBUserPassword,
		DatabaseName: env.DBName,
	}

	server := echo.New()
	router.UserInit(server, &sql)

	server.Logger.Fatal(server.Start(fmt.Sprintf(":%d", env.AppPort)))
}
