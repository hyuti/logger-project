package base

import (
	"fmt"

	"github.com/hyuti/logger-project/internal/collection/permission"
	"github.com/hyuti/logger-project/pkg/infrastructure/logger"
	"github.com/pocketbase/pocketbase/core"
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

func getOrCreateCollectionByName(app core.App, name *string) (*models.Collection, error) {
	collection, _ := app.Dao().FindCollectionByNameOrId(*name)
	if collection == nil {
		collection = new(models.Collection)
	}
	return collection, nil
}

func runValidation(input *CollectionInput) error {
	st := map[string]bool{}
	for _, f := range input.Fields {
		if f.Id == "" {
			return fmt.Errorf("%s schema field:id can not be empty string", f.Name)
		}
		_, ok := st[f.Id]
		if ok {
			return fmt.Errorf("%s schema field:id must be unique", f.Name)
		}
		st[f.Id] = true
	}
	return nil
}

func NewCollection(
	app core.App,
	l logger.Interface,
	input *CollectionInput,
) error {
	err := runValidation(input)
	if err != nil {
		return err
	}
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
