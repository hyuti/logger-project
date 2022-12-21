package cmd

import (
	"github.com/hyuti/logger-project/cmd/createuser"
	"github.com/hyuti/logger-project/config"
	"github.com/hyuti/logger-project/pkg/infrastructure/logger"
	"github.com/pocketbase/pocketbase"
	"github.com/spf13/cobra"
)

func RegisterCMD(
	handler *pocketbase.PocketBase,
	l logger.Interface,
	cfg *config.Config,
) {
	handler.RootCmd.AddCommand(&cobra.Command{
		Use: "createadmin",
		Run: createuser.CreateUser(handler, cfg),
	})
}
