package createuser

import (
	"fmt"

	"github.com/TcMits/ent-clean-template/config"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/spf13/cobra"
)

func CreateUser(app core.App, cfg *config.Config) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		admin := new(models.Admin)
		form := forms.NewAdminUpsert(app, admin)
		form.Email = cfg.Admin.Username
		form.Password = cfg.Admin.Password
		form.PasswordConfirm = cfg.Admin.Password
		fmt.Println(form)
		submitErr := form.Submit()

		if submitErr != nil {
			fmt.Printf("error creating admin: %s\n", submitErr)
		}
	}
}
