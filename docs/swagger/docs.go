// Package swagger Code generated by swaggo/swag. DO NOT EDIT
package swagger

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
        "/users/me": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Find user by me - ✅",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Find user by me",
                "responses": {
                    "200": {
                        "description": "사용자 조회 성공",
                        "schema": {
                            "$ref": "#/definitions/presenter.FindUserByMeResponse"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Patch user by me - ✅",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Patch user by me",
                "parameters": [
                    {
                        "description": "사용자 수정 요청",
                        "name": "PatchUserByMeRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.PatchUserByMeRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "사용자 수정 성공"
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.PatchUserByMeRequest": {
            "type": "object",
            "properties": {
                "nickname": {
                    "type": "string"
                },
                "resolution": {
                    "type": "string"
                }
            }
        },
        "domain.Provider": {
            "type": "string",
            "enum": [
                "GOOGLE",
                "APPLE",
                "KAKAO"
            ],
            "x-enum-varnames": [
                "ProviderGoogle",
                "ProviderApple",
                "ProviderKakao"
            ]
        },
        "presenter.FindUserByMeResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "nickname": {
                    "type": "string"
                },
                "omitemtpy": {
                    "type": "string"
                },
                "provider": {
                    "$ref": "#/definitions/domain.Provider"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
