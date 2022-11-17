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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user/login": {
            "post": {
                "description": "用户登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "用户登录请求",
                        "name": "userLoginRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserLoginReq"
                        }
                    }
                ],
                "responses": {
                    "20000": {
                        "description": "登录成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Status"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/service.User"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "40001": {
                        "description": "请求参数格式错误",
                        "schema": {
                            "$ref": "#/definitions/response.Status"
                        }
                    },
                    "40002": {
                        "description": "请求参数不符合要求",
                        "schema": {
                            "$ref": "#/definitions/response.Status"
                        }
                    }
                }
            }
        },
        "/user/logout": {
            "get": {
                "description": "退出登录",
                "tags": [
                    "用户"
                ],
                "summary": "退出登录",
                "responses": {
                    "20000": {
                        "description": "退出成功",
                        "schema": {
                            "$ref": "#/definitions/response.Status"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "用户注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "用户注册请求",
                        "name": "userRegisterRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserRegisterReq"
                        }
                    }
                ],
                "responses": {
                    "20000": {
                        "description": "注册成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Status"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/service.User"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "40001": {
                        "description": "请求参数格式错误",
                        "schema": {
                            "$ref": "#/definitions/response.Status"
                        }
                    },
                    "40002": {
                        "description": "请求参数不符合要求",
                        "schema": {
                            "$ref": "#/definitions/response.Status"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.UserLoginReq": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "example": "required;min:6;max:20"
                },
                "username": {
                    "type": "string",
                    "example": "required;min:2;max:20"
                }
            }
        },
        "request.UserRegisterReq": {
            "type": "object",
            "properties": {
                "avatar_url": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "example": "required;min:6;max:20"
                },
                "username": {
                    "type": "string",
                    "example": "required;min:2;max:20"
                }
            }
        },
        "response.Status": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "service.User": {
            "type": "object",
            "properties": {
                "avatarUrl": {
                    "type": "string",
                    "example": "https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png"
                },
                "createdAt": {
                    "type": "string",
                    "example": "2021-01-01 00:00:00"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "password": {
                    "type": "string",
                    "example": "123456"
                },
                "updatedAt": {
                    "type": "string",
                    "example": "2021-01-01 00:00:00"
                },
                "username": {
                    "type": "string",
                    "example": "admin"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "DJ-Blog API文档",
	Description:      "DJ-Blog API文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
