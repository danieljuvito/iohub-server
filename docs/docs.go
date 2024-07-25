// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Daniel Juvito",
            "url": "danieljuvito.github.io",
            "email": "danieljuvito@outlook.com"
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
        "/feed": {
            "get": {
                "description": "Retrieve user feeds.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get user feeds",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/user.FeedResponse"
                            }
                        }
                    }
                }
            }
        },
        "/follow/{id}": {
            "patch": {
                "description": "Follow a specific user by ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Follow a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully followed user",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/identity": {
            "get": {
                "description": "Retrieve user identity information.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get user identity",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.IdentityResponse"
                        }
                    }
                }
            }
        },
        "/log-in": {
            "post": {
                "description": "Authenticates a user based on email and password.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sessions"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "User login request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/session.LogInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful login",
                        "schema": {
                            "$ref": "#/definitions/session.LogInResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {}
                    }
                }
            }
        },
        "/log-out": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Logs out the currently authenticated user.",
                "tags": [
                    "Sessions"
                ],
                "summary": "User logout",
                "responses": {
                    "204": {
                        "description": "No content"
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "Responds with \"pong\" as a health check.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Ping endpoint",
                "responses": {
                    "200": {
                        "description": "pong",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/sign-up": {
            "post": {
                "description": "Register a new user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "User sign-up",
                "parameters": [
                    {
                        "description": "Sign-up request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/user.SignUpResponse"
                        }
                    }
                }
            }
        },
        "/stories": {
            "get": {
                "description": "Retrieve a list of stories.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stories"
                ],
                "summary": "Get a list of stories",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/story.ListResponse"
                            }
                        }
                    }
                }
            }
        },
        "/story": {
            "post": {
                "description": "Creates a new story with the provided content.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stories"
                ],
                "summary": "Create a story",
                "parameters": [
                    {
                        "description": "Story content",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/story.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Story created successfully",
                        "schema": {
                            "$ref": "#/definitions/story.CreateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {}
                    }
                }
            }
        },
        "/story/{id}": {
            "get": {
                "description": "Retrieves a story based on the provided ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stories"
                ],
                "summary": "Get a story by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Story ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Story retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/story.GetResponse"
                        }
                    },
                    "404": {
                        "description": "Story not found",
                        "schema": {}
                    }
                }
            }
        },
        "/unfollow/{id}": {
            "patch": {
                "description": "Unfollow a specific user by ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Unfollow a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully unfollowed user",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Retrieve a list of users.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get a list of users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/user.ListResponse"
                            }
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "Retrieve user details by ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get user details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.GetResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "session.LogInRequest": {
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
        "session.LogInResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "story.CreateRequest": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                }
            }
        },
        "story.CreateResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "story.GetResponse": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "story.ListResponse": {
            "type": "object",
            "properties": {
                "story_id": {
                    "type": "string"
                }
            }
        },
        "user.FeedResponse": {
            "type": "object",
            "properties": {
                "is_followed": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "story_count": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "user.GetResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "user.IdentityResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "user.ListResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "is_followed": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "user.SignUpRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "user.SignUpResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
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
	Title:            "IOHub Server",
	Description:      "This is API server for IOHub App.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
