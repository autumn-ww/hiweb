// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/autumnzw/hiweb"
	"net/http"

	"github.com/alecthomas/template"
)

func init() {
	http.HandleFunc("/swag/", hiweb.Handler(
		hiweb.URL("./swagger.json", "example"), //The url pointing to API definition"
	))
}

var doc = `{
    "openapi": "3.0.1",
    "info": {
        "title": "example",
        "version": "v1"
    },
    "paths": {
        "/Token/GenToken": {
            "post": {
                "tags": [
                    "Token"
                ],
                "summary": "",
                "requestBody": {
                    "content": {
                        "application/*+json": {
                            "schema": {
                                "$ref": "#/components/schemas/UserCredentials"
                            }
                        },
                        "application/json": {
                            "schema": {
                                "$ref": "#/components/schemas/UserCredentials"
                            }
                        },
                        "application/json-patch+json": {
                            "schema": {
                                "$ref": "#/components/schemas/UserCredentials"
                            }
                        },
                        "text/json": {
                            "schema": {
                                "$ref": "#/components/schemas/UserCredentials"
                            }
                        }
                    }
                },
                "responses": {
                    "200": {
                        "description": "Success"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "403": {
                        "description": "Forbidden"
                    }
                }
            }
        },
        "/User/GetUser": {
            "post": {
                "tags": [
                    "User"
                ],
                "summary": "get user",
                "responses": {
                    "200": {
                        "description": "Success"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "403": {
                        "description": "Forbidden"
                    }
                },
                "security": [
                    {
                        "oauth2": []
                    }
                ]
            }
        }
    },
    "components": {
        "schemas": {
            "UserCredentials": {
                "type": "object",
                "properties": {
                    "password": {
                        "type": "string"
                    },
                    "username": {
                        "type": "string"
                    }
                },
                "additionalProperties": false
            }
        },
        "securitySchemes": {
            "oauth2": {
                "type": "apiKey",
                "description": "JWT授权(数据将在请求头中进行传输) 直接在下框中输入Bearer {token}（注意两者之间是一个空格）\"",
                "name": "Authorization",
                "in": "header"
            }
        }
    }
}
`

type swaggerInfo struct {
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo

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
	hiweb.SwaggerRegister(&s{})

	token := Token{}

	hiweb.Route("/Token/GenToken", &token, "", "post:GenToken", hiweb.RouteOption{IsAuth: false})

	user := User{}

	hiweb.Route("/User/GetUser", &user, "", "post:GetUser", hiweb.RouteOption{IsAuth: true})

}