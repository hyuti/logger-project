package createuser

import (
	"fmt"

	"github.com/hyuti/logger-project/config"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/spf13/cobra"
)

func getOrCreateAdmin(app core.App, email *string) (*models.Admin, error) {
	ins, _ := app.Dao().FindAdminByEmail(*email)
	if ins == nil {
		ins = new(models.Admin)
	}
	return ins, nil
}

func CreateUser(app core.App, cfg *config.Config) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		admin, err := getOrCreateAdmin(app, &cfg.Admin.Username)
		if err != nil {
			fmt.Printf("error creating admin: %s\n", err)
		}

		form := forms.NewAdminUpsert(app, admin)
		form.Email = cfg.Admin.Username
		form.Password = cfg.Admin.Password
		form.PasswordConfirm = cfg.Admin.Password

		err = form.Submit()

		if err != nil {
			fmt.Printf("error creating admin: %s\n", err)
		}
	}
}
