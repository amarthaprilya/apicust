// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/customers": {
            "get": {
                "description": "Get list of customers",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customers"
                ],
                "summary": "Get all customers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Customer"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new customer entry",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customers"
                ],
                "summary": "Create a new customer",
                "parameters": [
                    {
                        "description": "Customer Data",
                        "name": "customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Customer"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Customer"
                        }
                    }
                }
            }
        },
        "/customers/{id}": {
            "get": {
                "description": "Get a single customer by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customers"
                ],
                "summary": "Get customer by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Customer ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Customer"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customers"
                ],
                "summary": "Update customer by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Customer ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Customer Data",
                        "name": "customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Customer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Customer"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a customer",
                "tags": [
                    "customers"
                ],
                "summary": "Delete customer by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Customer ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Customer": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:8080/",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Sweager Service API",
	Description:      "Sweager service API in Go using Gin framework",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
