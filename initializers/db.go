package initializers

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type SQL struct {
	DB           *sqlx.DB
	Host         string
	Port         int
	UserName     string
	Password     string
	DatabaseName string
}

func (s *SQL) Connect() {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		s.UserName, s.Password, s.Host, s.Port, s.DatabaseName)

	s.DB = sqlx.MustConnect("mysql", dataSource)

	if err := s.DB.Ping(); err != nil {
		log.Error(err.Error())
		return
	}

	// fmt.Println("Connect database OK")
}

func (s *SQL) Close() {
	if err := s.DB.Close(); err != nil {
		return
	}
}
