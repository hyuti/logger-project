package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `[
			{
				"id": "_pb_users_auth_",
				"created": "2022-11-30 08:25:14.622Z",
				"updated": "2022-11-30 10:47:11.875Z",
				"name": "users",
				"type": "auth",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "users_name",
						"name": "name",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "users_avatar",
						"name": "avatar",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 5242880,
							"mimeTypes": [
								"image/jpg",
								"image/jpeg",
								"image/png",
								"image/svg+xml",
								"image/gif"
							],
							"thumbs": null
						}
					}
				],
				"listRule": "id = @request.auth.id",
				"viewRule": "id = @request.auth.id",
				"createRule": "",
				"updateRule": "id = @request.auth.id",
				"deleteRule": "id = @request.auth.id",
				"options": {
					"allowEmailAuth": true,
					"allowOAuth2Auth": true,
					"allowUsernameAuth": true,
					"exceptEmailDomains": null,
					"manageRule": null,
					"minPasswordLength": 8,
					"onlyEmailDomains": null,
					"requireEmail": false
				}
			},
			{
				"id": "k4nslxjwrpkhvhw",
				"created": "2022-11-30 08:30:40.032Z",
				"updated": "2022-12-20 09:42:04.108Z",
				"name": "the_hill_customer",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "id_method",
						"name": "method",
						"type": "select",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"GET",
								"OPTION",
								"PUT",
								"PATCH",
								"POST"
							]
						}
					},
					{
						"system": false,
						"id": "id_status",
						"name": "status",
						"type": "number",
						"required": true,
						"unique": false,
						"options": {
							"min": 100,
							"max": 599
						}
					},
					{
						"system": false,
						"id": "id_extra",
						"name": "extra",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "id_url",
						"name": "url",
						"type": "url",
						"required": true,
						"unique": false,
						"options": {
							"exceptDomains": null,
							"onlyDomains": null
						}
					}
				],
				"listRule": null,
				"viewRule": null,
				"createRule": "",
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "6qjohvzbtcv2ukl",
				"created": "2022-11-30 10:45:04.318Z",
				"updated": "2022-12-20 09:42:04.109Z",
				"name": "the_hill_admin",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "id_method",
						"name": "method",
						"type": "select",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"GET",
								"OPTION",
								"PUT",
								"PATCH",
								"POST"
							]
						}
					},
					{
						"system": false,
						"id": "id_status",
						"name": "status",
						"type": "number",
						"required": true,
						"unique": false,
						"options": {
							"min": 100,
							"max": 599
						}
					},
					{
						"system": false,
						"id": "id_url",
						"name": "url",
						"type": "url",
						"required": true,
						"unique": false,
						"options": {
							"exceptDomains": null,
							"onlyDomains": null
						}
					},
					{
						"system": false,
						"id": "id_extra",
						"name": "extra",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"listRule": null,
				"viewRule": null,
				"createRule": "",
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "12d2gtjk73cvlf2",
				"created": "2022-12-19 11:10:35.834Z",
				"updated": "2022-12-20 09:42:04.110Z",
				"name": "the_hill_store",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "id_method",
						"name": "method",
						"type": "select",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"GET",
								"OPTION",
								"PUT",
								"PATCH",
								"POST"
							]
						}
					},
					{
						"system": false,
						"id": "id_status",
						"name": "status",
						"type": "number",
						"required": true,
						"unique": false,
						"options": {
							"min": 100,
							"max": 599
						}
					},
					{
						"system": false,
						"id": "id_url",
						"name": "url",
						"type": "url",
						"required": true,
						"unique": false,
						"options": {
							"exceptDomains": null,
							"onlyDomains": null
						}
					},
					{
						"system": false,
						"id": "id_extra",
						"name": "extra",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"listRule": null,
				"viewRule": null,
				"createRule": "",
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		return nil
	})
}
