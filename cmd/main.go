package cmd

import (
	"github.com/TcMits/ent-clean-template/cmd/createuser"
	"github.com/TcMits/ent-clean-template/config"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
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
