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
        "/labels": {
            "get": {
                "description": "Get all labels.",
                "tags": [
                    "labels"
                ],
                "summary": "Get labels.",
                "parameters": [
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
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/server.getLabelResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create label.",
                "tags": [
                    "labels"
                ],
                "summary": "Create label.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Label",
                        "name": "label",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.createLabelRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.getLabelResponse"
                        }
                    }
                }
            }
        },
        "/labels/{label_id}": {
            "get": {
                "description": "Get label.",
                "tags": [
                    "labels"
                ],
                "summary": "Get label.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Label ID",
                        "name": "label_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.getLabelResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update label.",
                "tags": [
                    "labels"
                ],
                "summary": "Update label.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Label ID",
                        "name": "label_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Label",
                        "name": "label",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.createLabelRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.getLabelResponse"
                        }
                    }
                }
            }
        },
        "/labels/{label_id}/tasks": {
            "get": {
                "description": "Get tasks for label.",
                "tags": [
                    "labels"
                ],
                "summary": "Get tasks for label.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Label ID",
                        "name": "label_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/server.getTaskResponse"
                            }
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Login with username and password.",
                "tags": [
                    "user"
                ],
                "summary": "System login.",
                "parameters": [
                    {
                        "description": "Login request",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.LoginRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/register": {
            "post": {
                "description": "Register with username and password.",
                "tags": [
                    "user"
                ],
                "summary": "System register.",
                "parameters": [
                    {
                        "description": "Register request",
                        "name": "register",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.RegisterRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/tasks": {
            "get": {
                "description": "Get all tasks.",
                "tags": [
                    "tasks"
                ],
                "summary": "Get tasks.",
                "parameters": [
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
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/server.getTaskResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create one task, time format - 2018-12-10T13:49:51.141Z.",
                "tags": [
                    "tasks"
                ],
                "summary": "Create task.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Task",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.createTaskRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/tasks/completed": {
            "get": {
                "description": "Get all completed tasks.",
                "tags": [
                    "tasks"
                ],
                "summary": "Get completed tasks.",
                "parameters": [
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
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/server.getTaskResponse"
                            }
                        }
                    }
                }
            }
        },
        "/tasks/uncompleted": {
            "get": {
                "description": "Get all uncompleted tasks.",
                "tags": [
                    "tasks"
                ],
                "summary": "Get uncompleted tasks.",
                "parameters": [
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
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/server.getTaskResponse"
                            }
                        }
                    }
                }
            }
        },
        "/tasks/{task_id}": {
            "get": {
                "description": "Get one task.",
                "tags": [
                    "tasks"
                ],
                "summary": "Get task.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Task ID",
                        "name": "task_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.getTaskResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update one task.",
                "tags": [
                    "tasks"
                ],
                "summary": "Update task.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Task ID",
                        "name": "task_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Task",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.createTaskRequest"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete one task.",
                "tags": [
                    "tasks"
                ],
                "summary": "Delete task.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Task ID",
                        "name": "task_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/tasks/{task_id}/comments": {
            "get": {
                "description": "Get comments for task by Task ID.",
                "tags": [
                    "comments"
                ],
                "summary": "Get task comments.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Task ID",
                        "name": "task_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/server.getCommentResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create comment for task by Task ID.",
                "tags": [
                    "comments"
                ],
                "summary": "Create task comment.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Task ID",
                        "name": "task_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Comment",
                        "name": "comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.createCommentRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/tasks/{task_id}/comments/{comment_id}": {
            "get": {
                "description": "Get task comment by Task ID and Comment ID.",
                "tags": [
                    "comments"
                ],
                "summary": "Get task comment.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Task ID",
                        "name": "task_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Comment ID",
                        "name": "comment_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.getCommentResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update task comment by Task ID and Comment ID.",
                "tags": [
                    "comments"
                ],
                "summary": "Update comment.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Task ID",
                        "name": "task_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Comment ID",
                        "name": "comment_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Comment",
                        "name": "comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.createCommentRequest"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete task comment by Task ID and Comment ID.",
                "tags": [
                    "comments"
                ],
                "summary": "Delete comment from task.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Task ID",
                        "name": "task_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Comment ID",
                        "name": "comment_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/tasks/{task_id}/complete": {
            "post": {
                "description": "Complete one task.",
                "tags": [
                    "tasks"
                ],
                "summary": "Complete task.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Task ID",
                        "name": "task_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/tasks/{task_id}/labels/{label_id}": {
            "post": {
                "description": "Add label to task.",
                "tags": [
                    "tasks"
                ],
                "summary": "Add label to task.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Task ID",
                        "name": "task_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Label ID",
                        "name": "label_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/tasks/{task_id}/labels/{label_id}/": {
            "delete": {
                "description": "Remove label from task.",
                "tags": [
                    "tasks"
                ],
                "summary": "Remove label from task.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Task ID",
                        "name": "task_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Label ID",
                        "name": "label_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/tasks/{task_id}/uncomplete": {
            "post": {
                "description": "uncomplete one task.",
                "tags": [
                    "tasks"
                ],
                "summary": "uncomplete task.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Task ID",
                        "name": "task_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "server.LoginRequest": {
            "type": "object",
            "required": [
                "login"
            ],
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                }
            }
        },
        "server.RegisterRequest": {
            "type": "object",
            "required": [
                "login"
            ],
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                }
            }
        },
        "server.createCommentRequest": {
            "type": "object",
            "required": [
                "description",
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "server.createLabelRequest": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "title": {
                    "type": "string"
                }
            }
        },
        "server.createTaskRequest": {
            "type": "object",
            "required": [
                "description",
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "dueDate": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "server.getCommentResponse": {
            "type": "object",
            "properties": {
                "comment_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "task_id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "server.getLabelResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "label_id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "server.getTaskResponse": {
            "type": "object",
            "properties": {
                "completedAt": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "dueDate": {
                    "type": "string"
                },
                "task_id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "TODO API",
	Description:      "This is a sample server TODO server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
