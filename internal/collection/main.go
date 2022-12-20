package collection

import (
	"github.com/TcMits/ent-clean-template/config"
	"github.com/TcMits/ent-clean-template/internal/collection/log"
	valLog "github.com/TcMits/ent-clean-template/internal/collection/validation/log"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func RegisterCollections(
	handler core.App,
	l logger.Interface,
	cfg *config.Config,
) {
	errs := []error{}
	var err error
	// add more collections here
	err = log.NewTheHillCustomerLogCollection(handler, l, cfg)
	if err != nil {
		errs = append(errs, err)
	}
	err = log.NewTheHillAdminLogCollection(handler, l, cfg)
	if err != nil {
		errs = append(errs, err)
	}
	err = log.NewTheHillStoreLogCollection(handler, l, cfg)
	if err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		errAsStr := ""
		for _, e := range errs {
			errAsStr += " - "
			errAsStr += e.Error()
		}
		l.Fatal("internal.collection.RegisterCollections:%s", errAsStr)
	}
}

func registerCreateValidation(
	ctx echo.Context,
	rec *models.Record,
	l logger.Interface,
	cfg *config.Config,
) error {
	err := valLog.RegisterLogBeforeCreateValidation(ctx, rec, cfg)
	if err != nil {
		return err
	}
	return nil
}

func Validate(
	ctx echo.Context,
	rec *models.Record,
	l logger.Interface,
	cfg *config.Config,
) error {
	return registerCreateValidation(ctx, rec, l, cfg)
}
