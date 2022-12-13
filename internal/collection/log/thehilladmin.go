package log

import (
	"github.com/TcMits/ent-clean-template/config"
	"github.com/TcMits/ent-clean-template/internal/collection/base"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/pocketbase/pocketbase"
)

func NewTheHillAdminLogCollection(
	app *pocketbase.PocketBase,
	l logger.Interface,
	cfg *config.Config,
) error {
	input := defaultLogCollection(cfg)
	input.Name = "the_hill_admin"
	return base.NewCollection(app, l, input)
}
