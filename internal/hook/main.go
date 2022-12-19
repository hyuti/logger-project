package hook

import (
	"time"

	"github.com/TcMits/ent-clean-template/config"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/go-co-op/gocron"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func registerHook(s *gocron.Scheduler, handler core.App, projName string, amountToDel int, l logger.Interface, cfg *config.Config) {
	s.Every(cfg.Schedule.Period).Seconds().Do(func() {
		l.Info("Running %s task...", projName)
		collection, _ := handler.Dao().FindCollectionByNameOrId(projName)
		query := handler.Dao().RecordQuery(collection).
			Where(dbx.NotBetween("status", 500, 599)).
			OrderBy("created ASC").
			Limit(int64(amountToDel))
		rows := []dbx.NullStringMap{}
		if err := query.All(&rows); err != nil {
			l.Fatal("deleting %s logs failed due to %s", projName, err)
		}
		records := models.NewRecordsFromNullStringMaps(collection, rows)
		for _, r := range records {
			err := handler.Dao().DeleteRecord(r)
			if err != nil {
				l.Fatal("deleting %s logs failed due to %s", projName, err)
			}
		}
	})
}

func RegisterHook(handler core.App, l logger.Interface, cfg *config.Config) *gocron.Scheduler {
	s := gocron.NewScheduler(time.UTC)

	registerHook(s, handler, cfg.TheHillAdminProject.Name, cfg.TheHillAdminProject.AmountToDelete, l, cfg)
	registerHook(s, handler, cfg.TheHillCustomerProject.Name, cfg.TheHillCustomerProject.AmountToDelete, l, cfg)
	registerHook(s, handler, cfg.TheHillStoreProject.Name, cfg.TheHillStoreProject.AmountToDelete, l, cfg)

	s.StartAsync()
	return s
}
