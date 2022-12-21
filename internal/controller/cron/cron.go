package cron

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/hyuti/logger-project/config"
	"github.com/hyuti/logger-project/pkg/infrastructure/logger"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/types"
)

func NewScheduler() *gocron.Scheduler {
	return gocron.NewScheduler(time.UTC)
}

func registerHook(s *gocron.Scheduler, handler core.App, projName string, l logger.Interface, cfg *config.Config) {
	startAfter1Hour := time.Now().UTC().Add(time.Hour)

	s.Every(cfg.Schedule.Period).Seconds().StartAt(startAfter1Hour).Do(func() {
		l.Info("Running %s task...", projName)

		now := time.Now().UTC().Add(-time.Hour).Format(types.DefaultDateLayout)
		collection, _ := handler.Dao().FindCollectionByNameOrId(projName)
		query := handler.Dao().RecordQuery(collection).
			AndWhere(dbx.NotBetween("status", 500, 599)).
			AndWhere(dbx.NewExp("created <= {:date}", dbx.Params{"date": now})).
			OrderBy("created ASC")
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
	registerHook(scheduler, handler, cfg.TheHillAdminProject.Name, l, cfg)
	registerHook(scheduler, handler, cfg.TheHillCustomerProject.Name, l, cfg)
	registerHook(scheduler, handler, cfg.TheHillStoreProject.Name, l, cfg)
}
