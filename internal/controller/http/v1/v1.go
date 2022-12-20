package v1

import (
	"github.com/pocketbase/pocketbase"
)

func NewHandler() *pocketbase.PocketBase {
	handler := pocketbase.New()
	return handler
}

