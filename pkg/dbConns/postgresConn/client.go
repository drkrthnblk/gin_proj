package postgresConn

import (
	"sync"
	"time"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/driver/postgres"
	"github.com/spf13/viper"
	// _ "github.com/newrelic/go-agent/v3/integrations/nrpgx"
)

var (
	singleton sync.Once
	postgresDB = &gorm.DB{}
)

type PostgresClient struct {}

func NewPostgresClient *PostgresClient {
	return &PostgresClient{}
}

func (pc PostgresClient)InitDB() error {
	singleton.Do(func(){
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslnode=disable",
			viper.GetString("database.postres.source.host"),
			viper.GetString("database.postres.source.user"),
			viper.GetString("database.postres.source.password"),
			viper.GetString("database.postres.source.db_name"),
			viper.GetInt("database.postres.source.port"),
		)
		postgresDB, err := gorm.Open(postgres.New(postgres.config{
			DriverName: "nrpgx",
			DSN: dsn,
		}), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			panic(fmt.Spritf("not able to connect to database. Error:- %s", err.Error()))
		}

		db, err := postgresDB.DB()
		if err != nil {
			panic(fmt.Spritf("error occurred while getting db instance object. Error:- %s", err.Error()))
		}

		db.SetMaxIdleConns(viper.GetInt("database.postgres.max_idle_connections"))
		db.SetMaxOpenConns(viper.GetInt("database.postgres.max_open_connections"))
		db.SetConnMaxLifetime(time.Duration(viper.GetInt("database.postgres.max_connection_lifetime")) * time.Minute)
	})
	return postgresDB
}

func (pc PostgresClient)GetClient(rCtx *request.Context) *gorm.DB {
	return postgresDB
}

func (pc PostgresClient)GetSession(rCtx *request.Context) *gorm.Session {
	return nil
}