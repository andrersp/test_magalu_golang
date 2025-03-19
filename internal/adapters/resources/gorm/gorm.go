package gorm

import (
	"fmt"
	"sync"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	instance *GormInstance
	once     sync.Once
)

type GormInstance struct {
	db  *gorm.DB
	url string
}

func (p *GormInstance) DB() *gorm.DB {
	return p.db
}

func (p *GormInstance) URL() string {
	return p.url
}

func Connect(pgOptions *PgOptions, pgConfig PgConfig) *GormInstance {
	once.Do(func() {
		instance = newDB(pgOptions, pgConfig)
	})

	return instance
}

func newDB(opt *PgOptions, pgConfig PgConfig) *GormInstance {
	opt.setDefaultValues()
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		opt.User, opt.Password, opt.Host, opt.Port, opt.DBName, opt.SSLMode,
	)

	gormConfig := gorm.Config{}

	if pgConfig.LogQuery {
		log.Info("Enabling query log")
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	if pgConfig.SingularTable {
		log.Info("Enabling singular tables")

		gormConfig.NamingStrategy = schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: pgConfig.SingularTable,
			NameReplacer:  nil,
			NoLowerCase:   false,
		}
	}

	dbConnection, err := gorm.Open(postgres.Open(dsn), &gormConfig)
	if err != nil {
		log.Errorf("error on %s connection: %s", "pg", err.Error())
	}

	sqlDB, err := dbConnection.DB()
	if err != nil {
		log.Errorf("error on %s connection: %s", "pg", err.Error())
	}

	sqlDB.SetConnMaxIdleTime(opt.MaxIdleTime)
	sqlDB.SetMaxOpenConns(opt.PoolSize)

	if err := sqlDB.Ping(); err != nil {
		log.Errorf("error on %s connection: %s", "pg", err.Error())
	}

	log.Infof("%s successfully connected %s", "pg", opt.Host)

	return &GormInstance{
		db:  dbConnection,
		url: dsn,
	}
}
