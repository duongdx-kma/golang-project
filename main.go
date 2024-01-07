package main

import (
	"database/sql"
	"duongdx/example/initializers"
	"duongdx/example/router"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?multiStatements=true",
		sql.UserName,
		sql.Password,
		sql.Host,
		sql.Port,
		sql.DatabaseName,
	)

	log.Printf("Running on mainnnnnnnnnnnn")

	MigrateRunning(env.DBDriver, "file://databases/migrations", dbSource)

	server := echo.New()
	router.UserInit(server, &sql)

	server.Logger.Fatal(server.Start(fmt.Sprintf(":%d", env.AppPort)))
}

func MigrateRunning(sqlDriver string, migrationsURL string, dbSource string) {
	db, err := sql.Open(sqlDriver, dbSource)

	log.Println("Running on migrate", sqlDriver, migrationsURL, dbSource)
	// root:password@tcp(database:3306)/db_business?multiStatements=true
	if err != nil {
		log.Fatal("Open connect mysql failed", err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal("Cannot create new migrate instance", err)
	}

	migration, err := migrate.NewWithDatabaseInstance(
		migrationsURL,
		sqlDriver,
		driver,
	)

	if err := migration.Up(); err != nil {
		log.Fatal("Cannot migrate database", err)
	}

	log.Println("database migrated successfully !")
}
