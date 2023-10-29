package repositories

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

type DataSource struct {
	*sqlx.DB
	dbDriver         string
	dbUrl            string
	maxIdleConns     int
	maxOpenConns     int
	maxConnsLifetime time.Duration
}

func NewDatabaseConnectionWithConnectionPool(dbDriver string, dbUrl string, maxIdle int, maxOpenConnection int, _maxLifeTime int) (*DataSource, error) {
	db, err := sqlx.Connect(dbDriver, dbUrl)
	if err != nil {
		log.Info(err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	maxLifeTime := time.Duration(_maxLifeTime) * time.Second

	db.SetMaxIdleConns(maxIdle)
	db.SetMaxOpenConns(maxOpenConnection)
	db.SetConnMaxLifetime(maxLifeTime)
	return &DataSource{db, dbDriver, dbUrl, maxIdle, maxOpenConnection, maxLifeTime}, nil
}

func NewDatabaseConnection(dbDriver string, dbUrl string) (*DataSource, error) {
	return NewDatabaseConnectionWithConnectionPool(dbDriver, dbUrl, 10, 10, 1)
}
