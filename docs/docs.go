// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/export": {
            "post": {
                "description": "Create a new export record",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "export-apis"
                ],
                "summary": "Create a new export record",
                "parameters": [
                    {
                        "description": "export information",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.ExportPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ExportResponse"
                        }
                    }
                }
            }
        },
        "/export/active/{Id}": {
            "put": {
                "description": "Active a export record",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "export-apis"
                ],
                "summary": "Active a export record",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "export id",
                        "name": "Id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ExportDetailResponse"
                        }
                    }
                }
            }
        },
        "/export/deactive/{Id}": {
            "put": {
                "description": "DeActive a export record",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "export-apis"
                ],
                "summary": "DeActive a export record",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "export id",
                        "name": "Id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ExportDetailResponse"
                        }
                    }
                }
            }
        },
        "/export/update/{Id}": {
            "put": {
                "description": "Update a export record",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "export-apis"
                ],
                "summary": "Update a export record",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "export id",
                        "name": "Id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "export information",
                        "name": "export",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.UpdateExport"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ExportDetailResponse"
                        }
                    }
                }
            }
        },
        "/export/{Id}": {
            "get": {
                "description": "Get a export record",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "export-apis"
                ],
                "summary": "Get a export record",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "export id",
                        "name": "Id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ExportDetailResponse"
                        }
                    }
                }
            }
        },
        "/exports": {
            "get": {
                "description": "Get list of export records",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "export-apis"
                ],
                "summary": "Get list of export records",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "pageSize",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ListExportsResponse"
                        }
                    }
                }
            }
        },
        "/import": {
            "post": {
                "description": "Create a new import record",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "import-apis"
                ],
                "summary": "Create a new import record",
                "parameters": [
                    {
                        "description": "import information",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.ImportPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ImportResponse"
                        }
                    }
                }
            }
        },
        "/import/active/{Id}": {
            "put": {
                "description": "Active a import record",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "import-apis"
                ],
                "summary": "Active a import record",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "import id",
                        "name": "Id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ImportDetailResponse"
                        }
                    }
                }
            }
        },
        "/import/deactive/{Id}": {
            "put": {
                "description": "DeActive a import record",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "import-apis"
                ],
                "summary": "DeActive a import record",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "import id",
                        "name": "Id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ImportDetailResponse"
                        }
                    }
                }
            }
        },
        "/import/update/{Id}": {
            "put": {
                "description": "Update a import record",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "import-apis"
                ],
                "summary": "Update a import record",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "import id",
                        "name": "Id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "import information",
                        "name": "import",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.UpdateImport"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ImportDetailResponse"
                        }
                    }
                }
            }
        },
        "/import/{Id}": {
            "get": {
                "description": "Get a import record",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "import-apis"
                ],
                "summary": "Get a import record",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "import id",
                        "name": "Id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ImportDetailResponse"
                        }
                    }
                }
            }
        },
        "/imports": {
            "get": {
                "description": "Get list of import records",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "import-apis"
                ],
                "summary": "Get list of import records",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "pageSize",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ListImportsResponse"
                        }
                    }
                }
            }
        },
        "/inventories": {
            "get": {
                "description": "Get list of inventory records",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "inventory-apis"
                ],
                "summary": "Get list of inventory records",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "pageSize",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ListInventoriesResponse"
                        }
                    }
                }
            }
        },
        "/inventory/{Id}": {
            "get": {
                "description": "Get a inventory record",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "inventory-apis"
                ],
                "summary": "Get a inventory record",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "inventory id",
                        "name": "Id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.InventoryDetailResponse"
                        }
                    }
                }
            }
        },
        "/product": {
            "post": {
                "description": "Create a new product",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "product-apis"
                ],
                "summary": "Create a new product",
                "parameters": [
                    {
                        "description": "product information",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.ProductPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ProductResponse"
                        }
                    }
                }
            }
        },
        "/product/active/{Id}": {
            "put": {
                "description": "Active a product",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "product-apis"
                ],
                "summary": "Active a product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "product id",
                        "name": "Id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ProductDetailResponse"
                        }
                    }
                }
            }
        },
        "/product/deactive/{Id}": {
            "put": {
                "description": "DeActive a product",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "product-apis"
                ],
                "summary": "DeActive a product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "product id",
                        "name": "Id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ProductDetailResponse"
                        }
                    }
                }
            }
        },
        "/product/update/{Id}": {
            "put": {
                "description": "Update a product",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "product-apis"
                ],
                "summary": "Update a product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "product id",
                        "name": "Id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "product information",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.UpdateProduct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ProductDetailResponse"
                        }
                    }
                }
            }
        },
        "/product/{Id}": {
            "get": {
                "description": "Get a product",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "product-apis"
                ],
                "summary": "Get a product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "product id",
                        "name": "Id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ProductDetailResponse"
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "description": "Get list of products",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "product-apis"
                ],
                "summary": "Get list of products",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "pageSize",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ListProductResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.ExportDetailResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/model.ExportRecord"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "controller.ExportResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "controller.ImportDetailResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/model.ImportRecord"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "controller.ImportResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "controller.InventoryDetailResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/model.Inventory"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "controller.ListExportsResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ExportRecord"
                    }
                },
                "message": {
                    "type": "string"
                },
                "meta": {
                    "$ref": "#/definitions/model.Meta"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "controller.ListImportsResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ImportRecord"
                    }
                },
                "message": {
                    "type": "string"
                },
                "meta": {
                    "$ref": "#/definitions/model.Meta"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "controller.ListInventoriesResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Inventory"
                    }
                },
                "message": {
                    "type": "string"
                },
                "meta": {
                    "$ref": "#/definitions/model.Meta"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "controller.ListProductResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Product"
                    }
                },
                "message": {
                    "type": "string"
                },
                "meta": {
                    "$ref": "#/definitions/model.Meta"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "controller.ProductDetailResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/model.Product"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "controller.ProductResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "controller.UpdateExport": {
            "type": "object",
            "properties": {
                "product_id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "controller.UpdateImport": {
            "type": "object",
            "properties": {
                "price": {
                    "type": "number"
                },
                "product_id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "controller.UpdateProduct": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "model.ExportRecord": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "locked": {
                    "type": "boolean"
                },
                "price": {
                    "type": "number"
                },
                "product_id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.ImportRecord": {
            "type": "object",
            "properties": {
                "Locked": {
                    "type": "boolean"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "product_id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.Inventory": {
            "type": "object",
            "properties": {
                "average_price": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "model.Meta": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "model.Product": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "service.ExportPayload": {
            "type": "object",
            "properties": {
                "product_id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "service.ImportPayload": {
            "type": "object",
            "properties": {
                "price": {
                    "type": "number"
                },
                "product_id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "service.ProductPayload": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "2.0",
	Host:        "localhost:3000",
	BasePath:    "/api/tda/student-management",
	Schemes:     []string{},
	Title:       "Swagger scheduling-api project API",
	Description: "This is list api for scheduling-api project",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
