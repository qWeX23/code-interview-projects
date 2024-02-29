package config

import (
	"joynet-assignment/db"
	"os"

	"github.com/sirupsen/logrus"
)

type Env struct {
	Db  db.DbInterface
	Log *logrus.Logger
}

func NewEnv() Env {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	sql, _ := db.NewDb(db.PgConfig{
		DbName:     os.Getenv("POSTGRES_DB"),
		DbPassword: os.Getenv("POSTGRES_PASSWORD"),
	}, log)
	return Env{
		Db:  sql,
		Log: log,
	}
}
