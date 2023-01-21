package postgresConn

import (
	"fmt"
	"sync"
	"time"

	// "context"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	// _ "github.com/newrelic/go-agent/v3/integrations/nrpgx"
	// "gin_proj/pkg/common/request"
)

type singleton struct {
	sync.Once
	postresDB *gorm.DB
}

// var (
// 	singleton sync.Once
// 	postgresDB = &gorm.DB{}
// )

var (
	singletonDB singleton
)

type PostgresClient struct{}

func NewPostgresClient() *PostgresClient {
	return &PostgresClient{}
}

func (pc PostgresClient) InitDB() error {
	singletonDB.Do(func() {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			viper.GetString("database.postgres.host"),
			viper.GetString("database.postgres.user"),
			viper.GetString("database.postgres.password"),
			viper.GetString("database.postgres.db_name"),
			viper.GetInt("database.postgres.port"),
		)
		postgresDB, err := gorm.Open(postgres.New(postgres.Config{
			// DriverName: "nrpgx",
			DSN: dsn,
		}), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			panic(fmt.Sprintf("not able to connect to database. Error:- %s", err.Error()))
		}
		singletonDB.postresDB = postgresDB

		db, err := postgresDB.DB()
		if err != nil {
			panic(fmt.Sprintf("error occurred while getting db instance object. Error:- %s", err.Error()))
		}

		db.SetMaxIdleConns(viper.GetInt("database.postgres.max_idle_connections"))
		db.SetMaxOpenConns(viper.GetInt("database.postgres.max_open_connections"))
		db.SetConnMaxLifetime(time.Duration(viper.GetInt("database.postgres.max_connection_lifetime")) * time.Minute)
	})
	return nil
}

// func (pc PostgresClient)GetClient(rCtx *request.Context) *gorm.DB {
// 	return singletonDB.postresDB
// }

// func (pc PostgresClient)GetSession(rCtx *request.Context) *gorm.Session {
// 	return nil
// }

func (pc PostgresClient) GetClient() (*gorm.DB, error) {
	return singletonDB.postresDB, nil
}

func (pc PostgresClient) GetSession() (*gorm.Session, error) {
	return nil, nil
}
