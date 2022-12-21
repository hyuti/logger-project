package log

import (
	"github.com/hyuti/logger-project/config"
	"github.com/hyuti/logger-project/internal/collection/base"
	"github.com/hyuti/logger-project/pkg/infrastructure/logger"
	"github.com/pocketbase/pocketbase/core"
)

func NewTheHillStoreLogCollection(
	app core.App,
	l logger.Interface,
	cfg *config.Config,
) error {
	input := defaultLogCollection(cfg)
	input.Name = cfg.TheHillStoreProject.Name
	return base.NewCollection(app, l, input)
}
