// Package partner GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package partner

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/partners/login": {
            "post": {
                "description": "Авторизация партнера",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Partner"
                ],
                "summary": "Авторизация партнера",
                "parameters": [
                    {
                        "description": "body",
                        "name": "auth",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.TokenResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httperrors.Response"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests",
                        "schema": {
                            "$ref": "#/definitions/httperrors.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httperrors.Response"
                        }
                    }
                }
            }
        },
        "/partners/logout": {
            "get": {
                "description": "выход авторизации партнера",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Partner"
                ],
                "summary": "выход авторизации партнера"
            }
        },
        "/partners/orders/": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Обновление заказа по решению партнера",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Обновление заказа по решению партнера",
                "parameters": [
                    {
                        "description": "body",
                        "name": "market",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.OrderPartnerUpdateRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httperrors.Response"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests",
                        "schema": {
                            "$ref": "#/definitions/httperrors.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httperrors.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.FIO": {
            "type": "object",
            "properties": {
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "middleName": {
                    "type": "string"
                }
            }
        },
        "dto.Login": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.OrderPartnerUpdateRequest": {
            "type": "object",
            "properties": {
                "customer": {
                    "$ref": "#/definitions/dto.FIO"
                },
                "offers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Offers"
                    }
                },
                "referenceId": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "stateTitle": {
                    "type": "string"
                }
            }
        },
        "dto.TokenResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "response_token": {
                    "type": "string"
                }
            }
        },
        "httperrors.Details": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "application-определенный код ошибки",
                    "type": "integer"
                },
                "message": {
                    "description": "application-level сообщение, для дебага",
                    "type": "string"
                },
                "status": {
                    "description": "сообщение пользовательского уровня",
                    "type": "string"
                }
            }
        },
        "httperrors.Response": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/httperrors.Details"
                },
                "validation": {
                    "description": "ошибки валидации",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                }
            }
        },
        "models.Offers": {
            "type": "object",
            "properties": {
                "contractNumber": {
                    "type": "string"
                },
                "loanAmount": {
                    "type": "string"
                },
                "loanLength": {
                    "type": "string"
                },
                "monthlyPayment": {
                    "type": "integer"
                },
                "product": {
                    "type": "string"
                },
                "productType": {
                    "type": "string"
                }
            }
        },
        "models.Response": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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