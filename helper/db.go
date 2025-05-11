package helper

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

func NewDb() (*sql.DB, func(), error) {
	db, err := sql.Open("mysql", "root:andhikad@/userdb")
	if err != nil {
		NewLoggerConfigure("db.log", logrus.ErrorLevel, err.Error(), logrus.ErrorLevel)
		return nil, nil, err
	}

	cleanup := func() {
		db.Close()
	}

	return db, cleanup, nil
}
