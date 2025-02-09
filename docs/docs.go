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
        "/like/": {
            "get": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "Get liked videos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "like"
                ],
                "summary": "Get liked videos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.LikeVideoResponse"
                            }
                        }
                    }
                }
            }
        },
        "/like/{videoID}": {
            "get": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "Get like status of video",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "like"
                ],
                "summary": "Get like status of video",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Video ID",
                        "name": "videoID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"is_liked\": \"true\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "Add like to video",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "like"
                ],
                "summary": "Add like to video",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Video ID",
                        "name": "videoID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\": \"Video liked successfully\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "Remove like from video",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "like"
                ],
                "summary": "Remove like from video",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Video ID",
                        "name": "videoID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\": \"Video like removed successfully\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ping"
                ],
                "summary": "ping",
                "responses": {
                    "200": {
                        "description": "{\"message\": \"pong\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/piped/opensearch/suggestions": {
            "get": {
                "description": "Suggestions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "piped"
                ],
                "summary": "Suggestions",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Query",
                        "name": "query",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Suggestions",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/piped/search": {
            "get": {
                "description": "Search",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "piped"
                ],
                "summary": "Search",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Query",
                        "name": "q",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Search results",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/piped/streams/{videoID}": {
            "get": {
                "description": "Get video streams by video ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "piped"
                ],
                "summary": "Get video streams",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Video ID",
                        "name": "videoID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Video streams",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/": {
            "delete": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "Delete user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete user",
                "responses": {
                    "200": {
                        "description": "{\"message\": \"User deleted successfully\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/devices": {
            "get": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "Get devices",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get devices",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.DeviceListResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "Delete devices",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete devices",
                "parameters": [
                    {
                        "description": "Devices to be deleted",
                        "name": "devices",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "integer"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\": \"Devices deleted successfully\", \"deleted\": [1, 2, 3]}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/email": {
            "patch": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "Edit email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Edit email",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateEmailRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\": \"Email updated successfully\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"access_token\": \"access_token\", \"refresh_token\": \"refresh_token\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/logout": {
            "post": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "Logout user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Logout user",
                "responses": {
                    "200": {
                        "description": "{\"message\": \"Logged out successfully\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/refresh_token": {
            "post": {
                "security": [
                    {
                        "RefreshToken": []
                    }
                ],
                "description": "Refresh token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Refresh token",
                "responses": {
                    "200": {
                        "description": "{\"access_token\": \"access_token\", \"refresh_token\": \"refresh_token\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/signup": {
            "post": {
                "description": "Signup",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Signup",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SignupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\": \"Registration successful\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.DeviceListResponse": {
            "type": "object",
            "properties": {
                "current_device_id": {
                    "type": "integer"
                },
                "devices": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.DeviceResponse"
                    }
                }
            }
        },
        "models.DeviceResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "ip_address": {
                    "type": "string"
                },
                "last_active": {
                    "type": "string"
                },
                "user_agent": {
                    "type": "string"
                }
            }
        },
        "models.LikeVideoResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "thumbnail_url": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.LoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.SignupRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.UpdateEmailRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "AccessToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "RefreshToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "AltTube API",
	Description:      "This is the API documentation for the AltTube application.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
