// Package app configures and runs application.
package app

import (
	"fmt"

	"github.com/TcMits/ent-clean-template/cmd/createuser"
	"github.com/TcMits/ent-clean-template/config"
	v1 "github.com/TcMits/ent-clean-template/internal/controller/http/v1"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/spf13/cobra"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// repository

	// HTTP Server
	handler := v1.NewHandler()

	// Register commands
	handler.RootCmd.AddCommand(&cobra.Command{
		Use: "createadmin",
		Run: createuser.CreateUser(handler, cfg),
	})

	v1.RegisterV1HTTPServices(handler, l)

	if err := handler.Start(); err != nil {
		l.Fatal(fmt.Errorf("app - Run - handler.Start: %w", err))
	}
}
