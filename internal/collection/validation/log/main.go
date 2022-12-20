package log

import (
	"fmt"

	"github.com/TcMits/ent-clean-template/config"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/models"
)

type payload struct {
	Token string `json:"token"`
}

func RegisterLogBeforeCreateValidation(ctx echo.Context, rec *models.Record, cfg *config.Config) error {
	clt := rec.Collection()
	if clt.Name == "the_hill_admin" || clt.Name == "the_hill_customer" {
		pl := new(payload)
		if err := ctx.Bind(pl); err != nil {
			return err
		}
		if pl.Token == "" {
			return fmt.Errorf("token is required")
		}
		if pl.Token != cfg.ApiKey {
			return fmt.Errorf("token is invalid")
		}
	}
	return nil
}
