package log

import (
	"github.com/TcMits/ent-clean-template/internal/collection/base"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

const (
	minStatusCode float64 = 200
	maxStatusCode float64 = 599
)

func defaultLogCollection() *base.CollectionInput {
	input := base.DefaultCollectionInput()
	input.Name = "default_log"
	input.Fields = []*schema.SchemaField{
		{
			Id:       "id_method",
			Name:     "method",
			Type:     schema.FieldTypeSelect,
			Required: true,
			Options: &schema.SelectOptions{
				MaxSelect: 1,
				Values: []string{
					"GET",
					"OPTION",
					"PUT",
					"PATCH",
					"POST",
				},
			},
		},
		{
			Id:       "id_status",
			Name:     "status",
			Type:     schema.FieldTypeNumber,
			Required: true,
			Options: &schema.NumberOptions{
				Min: types.Pointer(minStatusCode),
				Max: types.Pointer(maxStatusCode),
			},
		},
		{
			Id:       "id_url",
			Name:     "url",
			Type:     schema.FieldTypeUrl,
			Required: true,
		},
		{
			Id:       "id_extra",
			Name:     "extra",
			Type:     schema.FieldTypeText,
			Required: false,
		},
	}
	return input
}
