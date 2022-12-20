// Package app configures and runs application.
package app

import (
	"fmt"

	"github.com/hyuti/logger-project/cmd"
	"github.com/hyuti/logger-project/config"
	"github.com/hyuti/logger-project/internal/collection"
	"github.com/hyuti/logger-project/internal/controller/cron"
	v1 "github.com/hyuti/logger-project/internal/controller/http/v1"
	_ "github.com/hyuti/logger-project/migrations"
	"github.com/pocketbase/pocketbase/core"

	"github.com/hyuti/logger-project/pkg/infrastructure/logger"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)
	// HTTP Server
	handler := v1.NewHandler()
  scheduler := cron.NewScheduler()

	cmd.RegisterCMD(handler, l, cfg)

	migratecmd.MustRegister(handler, handler.RootCmd, &migratecmd.Options{
		Automigrate: true,
	})

	handler.OnRecordBeforeCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		return collection.Validate(e.HttpContext, e.Record, l, cfg)
	})

	handler.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		cron.RegisterHook(scheduler, e.App, l, cfg)
    scheduler.StartAsync()
		return nil
	})

	if err := handler.Start(); err != nil {
		l.Fatal(fmt.Errorf("app - Run - handler.Start: %w", err))
	}
}
