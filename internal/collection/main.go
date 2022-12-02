package collection

import (
	"github.com/TcMits/ent-clean-template/internal/collection/log"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func RegisterCollections(
	handler *pocketbase.PocketBase,
	// logger
	l logger.Interface,
) {
	handler.OnBeforeServe().Add(func(_ *core.ServeEvent) error {
		errs := []error{}
		var err error
		// add more collections here
		err = log.NewTheHillCustomerLogCollection(handler, l)
		if err != nil {
			errs = append(errs, err)
		}
		err = log.NewTheHillAdminLogCollection(handler, l)
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
