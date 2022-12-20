package cron

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/hyuti/logger-project/config"
	"github.com/hyuti/logger-project/pkg/infrastructure/logger"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func NewScheduler() *gocron.Scheduler {
  return gocron.NewScheduler(time.UTC)
}

func registerHook(s *gocron.Scheduler, handler core.App, projName string, amountToDel int, l logger.Interface, cfg *config.Config) {
	startAfter1Hour := time.Now().UTC().Add(time.Hour)

	s.Every(cfg.Schedule.Period).Seconds().StartAt(startAfter1Hour).Do(func() {
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

func RegisterHook(scheduler *gocron.Scheduler, handler core.App, l logger.Interface, cfg *config.Config) {
	registerHook(scheduler, handler, cfg.TheHillAdminProject.Name, cfg.TheHillAdminProject.AmountToDelete, l, cfg)
	registerHook(scheduler, handler, cfg.TheHillCustomerProject.Name, cfg.TheHillCustomerProject.AmountToDelete, l, cfg)
	registerHook(scheduler, handler, cfg.TheHillStoreProject.Name, cfg.TheHillStoreProject.AmountToDelete, l, cfg)
}
