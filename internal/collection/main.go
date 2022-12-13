package collection

import (
	"github.com/TcMits/ent-clean-template/config"
	"github.com/TcMits/ent-clean-template/internal/collection/log"
	valLog "github.com/TcMits/ent-clean-template/internal/collection/validation/log"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func RegisterCollections(
	handler *pocketbase.PocketBase,
	l logger.Interface,
	cfg *config.Config,
) {
	handler.OnBeforeServe().Add(func(_ *core.ServeEvent) error {
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

		if len(errs) > 0 {
			errAsStr := ""
			for _, e := range errs {
				errAsStr += " - "
				errAsStr += e.Error()
			}
			l.Fatal("internal.collection.RegisterCollections:%s", errAsStr)
		}
		return nil
	})
}

func registerCreateValidation(
	handler *pocketbase.PocketBase,
	l logger.Interface,
	cfg *config.Config,
) {
	handler.OnRecordBeforeCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		err := valLog.RegisterLogBeforeCreateValidation(e.HttpContext, e.Record, cfg)
		if err != nil {
			return err
		}
		return nil
	})
}

func RegisterValidation(
	handler *pocketbase.PocketBase,
	l logger.Interface,
	cfg *config.Config,
) {
	registerCreateValidation(handler, l, cfg)
}
