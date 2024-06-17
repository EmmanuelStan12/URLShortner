package services

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

func InitDBMock(t *testing.T) (*sql.DB, *gorm.DB, sqlmock.Sqlmock) {
	sqldb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	// Set expectation for the initial GORM version check query
	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow("8.0.23"))

	gormdb, err := gorm.Open(
		mysql.New(mysql.Config{Conn: sqldb}),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		t.Fatal(err)
	}

	return sqldb, gormdb, mock
}
