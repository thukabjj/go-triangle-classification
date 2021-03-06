// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://www.linkedin.com/in/arthur-alves-da-costa/",
        "contact": {
            "name": "Arthur Alves",
            "url": "https://twitter.com/prayformercy_tv",
            "email": "arthur.alvesdeveloper@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/triangle/v1/classifier": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Takes the triangle request JSON and identifies the triangle's type, and stores it in DB. Return saved JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "triangle"
                ],
                "summary": "Identify a tiangle type",
                "parameters": [
                    {
                        "description": "TriangleEntrypointRequest JSON information",
                        "name": "TraingleRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.TriangleEntrypointRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.TriangleEntrypointResponse"
                        }
                    },
                    "401": {
                        "description": "User not authorized!",
                        "schema": {
                            "$ref": "#/definitions/middleware.Error"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/middleware.Error"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Takes the username and the password from the Header and valid this information. Return JSON with the JWT information.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Make the authentication of an username",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "JWT informations",
                        "schema": {
                            "$ref": "#/definitions/entity.AuthenticationEntrypointResponse"
                        }
                    },
                    "401": {
                        "description": "User not authorized!",
                        "schema": {
                            "$ref": "#/definitions/middleware.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.AuthenticationEntrypointResponse": {
            "type": "object",
            "properties": {
                "expirationTime": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "entity.TriangleEntrypointRequest": {
            "type": "object",
            "required": [
                "sideA",
                "sideB",
                "sideC"
            ],
            "properties": {
                "sideA": {
                    "type": "number"
                },
                "sideB": {
                    "type": "number"
                },
                "sideC": {
                    "type": "number"
                }
            }
        },
        "entity.TriangleEntrypointResponse": {
            "type": "object",
            "properties": {
                "type": {
                    "type": "string"
                }
            }
        },
        "middleware.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "datails": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "status": {
                    "type": "string"
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

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Triangle API",
	Description:      "A triangle management service API in Go using Gin framework.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
