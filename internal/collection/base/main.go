package base

import (
	"github.com/TcMits/ent-clean-template/internal/collection/permission"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
)

type (
	CollectionInput struct {
		Name       string
		Type       string
		ListRule   *string
		ViewRule   *string
		CreateRule *string
		UpdateRule *string
		DeleteRule *string
		Fields     []*schema.SchemaField
	}
)

func saveCollection(form *forms.CollectionUpsert) error {
	if err := form.Submit(); err != nil {
		return err
	}
	return nil
}

func DefaultCollectionInput() *CollectionInput {
	return &CollectionInput{
		Type:       models.CollectionTypeBase,
		ListRule:   permission.IsSuperUser,
		ViewRule:   permission.IsSuperUser,
		CreateRule: permission.IsSuperUser,
		UpdateRule: permission.IsSuperUser,
		DeleteRule: permission.IsSuperUser,
	}
}

func getOrCreateCollectionByName(app *pocketbase.PocketBase, name *string) (*models.Collection, error) {
	collection, _ := app.Dao().FindCollectionByNameOrId(*name)
	if collection == nil {
		collection = new(models.Collection)
	}
	return collection, nil
}

func NewCollection(
	app *pocketbase.PocketBase,
	l logger.Interface,
	input *CollectionInput,
) error {
	collection, err := getOrCreateCollectionByName(app, &input.Name)
	if err != nil {
		return err
	}
	form := forms.NewCollectionUpsert(app, collection)
	form.Name = input.Name
	form.Type = input.Type
	form.ListRule = input.ListRule
	form.ViewRule = input.ViewRule
	form.CreateRule = input.CreateRule
	form.UpdateRule = input.UpdateRule
	form.DeleteRule = input.DeleteRule
	fieldsAsMap := form.Schema.AsMap()
	for _, f := range input.Fields {
		of := form.Schema.GetFieldById(f.Id)
		if of != nil {
			delete(fieldsAsMap, of.Name)

			of.System = f.System
			of.Name = f.Name
			of.Type = f.Type
			of.Required = f.Required
			of.Unique = f.Unique
			of.Options = f.Options
		} else {
			form.Schema.AddField(f)
		}
	}
	for _, v := range fieldsAsMap {
		form.Schema.RemoveField(v.Id)
	}
	return saveCollection(form)
}
