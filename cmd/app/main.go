package main

import (
	"github.com/MorselShogiew/Users-service-billing/application"
	"github.com/MorselShogiew/Users-service-billing/config"
	"github.com/MorselShogiew/Users-service-billing/logger"
	"github.com/MorselShogiew/Users-service-billing/logger/opt"
	"github.com/MorselShogiew/Users-service-billing/provider"
	"github.com/MorselShogiew/Users-service-billing/repos"
	"github.com/MorselShogiew/Users-service-billing/service/api"

	"github.com/google/uuid"

	_ "github.com/lib/pq"
)

func main() {
	conf := config.LoadConfig()
	conf.InstanceID = uuid.New()
	opts := makeLoggerOpts(conf)
	l := logger.New(opts)
	p := provider.New(conf, l)

	repositories := repos.New(p, l)

	resizePhotoService := api.New(l, repositories)

	app := application.New(conf, l, resizePhotoService)
	app.Start()
}

func makeLoggerOpts(c *config.Config) *opt.LoggerOpts {
	return &opt.LoggerOpts{
		Opts: &opt.GeneralOpts{
			InstanceID: c.InstanceID,
			Env:        c.Environment,
			AppName:    c.ApplicationName,
			Level:      c.Logger.Level,
		},
		StdLoggerOpts: &opt.StdLoggerOpts{
			LogFile:  c.Logger.LoggerStd.LogFile,
			Stdout:   c.Logger.LoggerStd.Stdout,
			Disabled: c.Logger.LoggerStd.Disabled,
		},
	}
}
