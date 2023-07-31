package main

import (
	"os"

	"github.com/goark/errs"
	"github.com/goark/gocli/exitcode"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zerologadapter"
	"github.com/jackc/pgx/v4/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Run() exitcode.ExitCode {
	// get logger
	zlogger := CreateLogger()

	// create gorm.DB instance for PostgreSQL service
	cfg, err := pgx.ParseConfig(os.Getenv("ELEPHANTSQL_URL"))
	if err != nil {
		zlogger.Error().Interface("error", errs.Wrap(err)).Send()
		return exitcode.Abnormal
	}
	cfg.Logger = zerologadapter.NewLogger(zlogger)
	cfg.LogLevel = pgx.LogLevelDebug
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: stdlib.OpenDB(*cfg),
	}), &gorm.Config{
		Logger: logger.Discard,
	})
	defer func() {
		if sqlDb, err := db.DB(); err == nil {
			sqlDb.Close()
		}
	}()

	// query
	var results []map[string]interface{}
	tx := db.Table("tablename").Find(&results) // "tablename" is not exist
	if tx.Error != nil {
		zlogger.Error().Interface("error", errs.Wrap(tx.Error)).Send()
		return exitcode.Abnormal
	}

	return exitcode.Normal
}
