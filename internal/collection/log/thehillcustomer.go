package log

import (
	"github.com/TcMits/ent-clean-template/config"
	"github.com/TcMits/ent-clean-template/internal/collection/base"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/pocketbase/pocketbase/core"
)

func NewTheHillCustomerLogCollection(
	app core.App,
	l logger.Interface,
	cfg *config.Config,
) error {
	input := defaultLogCollection(cfg)
	input.Name = cfg.TheHillCustomerProject.Name
	return base.NewCollection(app, l, input)
}
