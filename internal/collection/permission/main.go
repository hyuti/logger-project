package permission

import (
	"github.com/pocketbase/pocketbase/tools/types"
)

var (
	AllowAny        *string = types.Pointer("")
	IsAuthenticated *string = types.Pointer("@request.auth.id != ''")
	IsSuperUser     *string = nil
)
