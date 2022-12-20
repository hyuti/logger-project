// Package app configures and runs application.
package app

import (
	"fmt"

	"github.com/TcMits/ent-clean-template/cmd"
	"github.com/TcMits/ent-clean-template/config"
	"github.com/TcMits/ent-clean-template/internal/collection"
	v1 "github.com/TcMits/ent-clean-template/internal/controller/http/v1"
	"github.com/TcMits/ent-clean-template/internal/hook"
	_ "github.com/TcMits/ent-clean-template/migrations"
	"github.com/pocketbase/pocketbase/core"

	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)
	// HTTP Server
	handler := v1.NewHandler()

	cmd.RegisterCMD(handler, l, cfg)
	migratecmd.MustRegister(handler, handler.RootCmd, &migratecmd.Options{
		Automigrate: true,
	})
	handler.OnRecordBeforeCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		return collection.Validate(e.HttpContext, e.Record, l, cfg)
	})
	handler.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		collection.RegisterCollections(e.App, l, cfg)
		v1.RegisterV1HTTPServices(e.App, l)
		hook.RegisterHook(e.App, l, cfg)
		return nil
	})
	if err := handler.Start(); err != nil {
		l.Fatal(fmt.Errorf("app - Run - handler.Start: %w", err))
	}
}
