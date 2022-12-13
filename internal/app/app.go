// Package app configures and runs application.
package app

import (
	"fmt"

	"github.com/TcMits/ent-clean-template/cmd"
	"github.com/TcMits/ent-clean-template/config"
	"github.com/TcMits/ent-clean-template/internal/collection"
	v1 "github.com/TcMits/ent-clean-template/internal/controller/http/v1"

	_ "github.com/TcMits/ent-clean-template/migrations"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// repository

	// HTTP Server
	handler := v1.NewHandler()

	cmd.RegisterCMD(handler, l, cfg)
	v1.RegisterV1HTTPServices(handler, l)
	collection.RegisterCollections(handler, l, cfg)
	collection.RegisterValidation(handler, l, cfg)

	if err := handler.Start(); err != nil {
		l.Fatal(fmt.Errorf("app - Run - handler.Start: %w", err))
	}
}
