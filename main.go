package main

import (
	"database/sql"
	"duongdx/example/initializers"
	"duongdx/example/router"
	mySocket "duongdx/example/websocket"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	env, err := initializers.LoadConfig()

	if err != nil {
		log.Println("Loading environment error")
		return
	}

	if env.AppEnvironment == "prod" {
		log.Println("This is a prod environment -> will call data from secret manager")
		secretData := initializers.GetSecretManager(env.SecretManagerKey, env.Region)
		env.DBHost = secretData.Host
		env.DBPort = secretData.Port
		env.DBUserPassword = secretData.Password
		env.DBUserName = secretData.Username
	}

	sql := initializers.SQL{
		Host:         env.DBHost,
		Port:         env.DBPort,
		UserName:     env.DBUserName,
		Password:     env.DBUserPassword,
		DatabaseName: env.DBName,
	}

	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?multiStatements=true",
		sql.UserName,
		sql.Password,
		sql.Host,
		sql.Port,
		sql.DatabaseName,
	)

	MigrateRunning(env.DBDriver, "file://databases/migrations", dbSource)

	server := echo.New()
	server.GET("/", func(ctx echo.Context) error {
		log.Println("runninggggggggggg")

		return ctx.JSON(http.StatusOK, "App Healthy")
	})
	server.Use(middleware.CORS())
	router.UserInit(server, &sql)
	router.InitSocket(server, &sql)
	// socket here
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.GET("ws", func(ctx echo.Context) error {
		mySocket.NewWebSocket(ctx, &sql)

		return nil
	})
	server.Logger.Fatal(server.Start(fmt.Sprintf(":%d", env.AppPort)))
}

func MigrateRunning(sqlDriver string, migrationsURL string, dbSource string) {
	db, err := sql.Open(sqlDriver, dbSource)

	log.Println("Running on migrate", sqlDriver, migrationsURL, dbSource)
	// root:password@tcp(database:3306)/db_business?multiStatements=true
	if err != nil {
		log.Println("Open connect mysql failed", err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Println("Cannot create new migrate instance", err)
	}

	migration, err := migrate.NewWithDatabaseInstance(
		migrationsURL,
		sqlDriver,
		driver,
	)

	if err := migration.Up(); err != nil {
		log.Println("Cannot migrate database: ", err)
	}

	log.Println("database migrated successfully !")
}
