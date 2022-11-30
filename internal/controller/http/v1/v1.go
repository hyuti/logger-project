package v1

import (
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func NewHandler() *pocketbase.PocketBase {
	handler := pocketbase.New()
	return handler
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func RegisterV1HTTPServices(
	handler *pocketbase.PocketBase,
	// adding more usecases here
	// logger
	l logger.Interface,
) {
	handler.OnBeforeServe().Add(func(e *core.ServeEvent) error {

		return nil
	})
}
