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
        "/auth/logout": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Method that logout user on server side",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Logout",
                "operationId": "auth-logout",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Status"
                        }
                    }
                }
            }
        },
        "/auth/refresh-token": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Method that return new pair of access and refresh token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "RefreshToken",
                "operationId": "auth-refresh-token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/AuthResponse"
                        }
                    }
                }
            }
        },
        "/auth/sign-in": {
            "post": {
                "description": "Method allow to enter into account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "SignIn",
                "operationId": "auth-enter-account",
                "parameters": [
                    {
                        "description": "account info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/SignInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/AuthResponse"
                        }
                    }
                }
            }
        },
        "/auth/sign-up": {
            "post": {
                "description": "Method allow to create account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "SignUp",
                "operationId": "auth-create-account",
                "parameters": [
                    {
                        "description": "account info",
                        "name": "form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/AuthResponse"
                        }
                    }
                }
            }
        },
        "/budgets": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Method that return list of budgets",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Budgets"
                ],
                "operationId": "get-budgets",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/BudgetsResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Method that store budget",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Budgets"
                ],
                "operationId": "store-budgets",
                "parameters": [
                    {
                        "description": "Budget form",
                        "name": "form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/BudgetCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Status"
                        }
                    }
                }
            }
        },
        "/budgets/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Method that return budget by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Budgets"
                ],
                "operationId": "get-budget",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Budget ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/BudgetResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Method that allow to update budget by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Budgets"
                ],
                "operationId": "update-budgets",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Budget ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Budget form",
                        "name": "form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/BudgetUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Status"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/Error"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Method that allow to delete budget by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Budgets"
                ],
                "operationId": "delete-budgets",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Budget ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Status"
                        }
                    }
                }
            }
        },
        "/categories": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Method that return list of categories",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Categories"
                ],
                "operationId": "categories-get",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/CategoriesResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": [
                            "admin"
                        ]
                    }
                ],
                "description": "Method that store category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Categories"
                ],
                "operationId": "categories-store",
                "parameters": [
                    {
                        "description": "CategoryRequest form",
                        "name": "form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/CategoryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Status"
                        }
                    }
                }
            }
        },
        "/categories/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Method that return category by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Categories"
                ],
                "operationId": "category-get",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "CategoryRequest ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/CategoryResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": [
                            "admin"
                        ]
                    }
                ],
                "description": "Method that update category by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Categories"
                ],
                "operationId": "categories-update",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "CategoryRequest ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "CategoryRequest form",
                        "name": "form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/CategoryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Status"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": [
                            "admin"
                        ]
                    }
                ],
                "description": "Method that delete category by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Categories"
                ],
                "operationId": "categories-delete",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "CategoryRequest ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Status"
                        }
                    }
                }
            }
        },
        "/reports/pdf": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Method for generation PDF report",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "multipart/form-data"
                ],
                "tags": [
                    "Reports"
                ],
                "operationId": "reports-pdf",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Status"
                        }
                    }
                }
            }
        },
        "/reports/xlsx": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Method for generation XLSX report",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "multipart/form-data"
                ],
                "tags": [
                    "Reports"
                ],
                "operationId": "reports-xlsx",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Status"
                        }
                    }
                }
            }
        },
        "/roles": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": [
                            "admin"
                        ]
                    }
                ],
                "description": "Method that return list of roles",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Roles"
                ],
                "operationId": "get-roles",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/RolesResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": [
                            "admin"
                        ]
                    }
                ],
                "description": "Method that store role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Roles"
                ],
                "operationId": "store-roles",
                "parameters": [
                    {
                        "description": "Role form",
                        "name": "form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Role"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Status"
                        }
                    }
                }
            }
        },
        "/roles/permissions": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": [
                            "admin"
                        ]
                    }
                ],
                "description": "Method that return list of permissions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Roles"
                ],
                "operationId": "get-permissions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/PermissionsResponse"
                        }
                    }
                }
            }
        },
        "/roles/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": [
                            "admin"
                        ]
                    }
                ],
                "description": "Method that return role by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Roles"
                ],
                "operationId": "get-role",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Role ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/RoleResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": [
                            "admin"
                        ]
                    }
                ],
                "description": "Method that allow to update role by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Roles"
                ],
                "operationId": "update-roles",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Role ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Role form",
                        "name": "form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Role"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Status"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/Error"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": [
                            "admin"
                        ]
                    }
                ],
                "description": "Method that allow to delete role by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Roles"
                ],
                "operationId": "delete-roles",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Role ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Status"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/Error"
                        }
                    }
                }
            }
        },
        "/users/:id/assign-role": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": [
                            "admin"
                        ]
                    }
                ],
                "description": "Method that return list of permissions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Roles"
                ],
                "operationId": "assign-role-to-users",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Role form",
                        "name": "form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/AssignRoleToUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Status"
                        }
                    }
                }
            }
        },
        "/users/current": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Method allow to get current user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "operationId": "users-current",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/UserResponse"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Method allow to update users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "operationId": "users-update",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Users ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User form",
                        "name": "form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/UserUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Status"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "AssignRoleToUserRequest": {
            "type": "object",
            "properties": {
                "role_id": {
                    "type": "integer",
                    "example": 14
                }
            }
        },
        "AuthResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string",
                    "example": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aW.xkZXIiLCJpYXQiOjE3MjE4N"
                },
                "refresh_token": {
                    "type": "string",
                    "example": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aW.xkZXIiLCJpYXQiOjE3MjE4N"
                },
                "type": {
                    "type": "string",
                    "example": "Bearer"
                }
            }
        },
        "BudgetCreateRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number",
                    "example": 10.25
                },
                "category_slug": {
                    "type": "string",
                    "example": "test-category"
                },
                "title": {
                    "type": "string",
                    "example": "New car"
                },
                "type": {
                    "type": "string",
                    "example": "income"
                },
                "user_id": {
                    "type": "integer",
                    "example": 12
                }
            }
        },
        "BudgetResponse": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number",
                    "example": 1000
                },
                "id": {
                    "type": "integer",
                    "example": 10
                },
                "title": {
                    "type": "string",
                    "example": "New budget"
                },
                "type": {
                    "type": "string",
                    "example": "income"
                },
                "user_id": {
                    "type": "integer",
                    "example": 10
                }
            }
        },
        "BudgetUpdateRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number",
                    "example": 10.25
                },
                "title": {
                    "type": "string",
                    "example": "New car"
                },
                "type": {
                    "type": "string",
                    "example": "income"
                },
                "user_id": {
                    "type": "integer",
                    "example": 12
                }
            }
        },
        "BudgetsResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/BudgetResponse"
                    }
                }
            }
        },
        "CategoriesResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/CategoryResponse"
                    }
                }
            }
        },
        "CategoryRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "Transport"
                }
            }
        },
        "CategoryResponse": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "Transport"
                },
                "type": {
                    "type": "string",
                    "example": "income|expense"
                }
            }
        },
        "Error": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Something went wrong try again later!"
                }
            }
        },
        "PermissionsResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "prefix_create",
                        "prefix_read",
                        "..."
                    ]
                }
            }
        },
        "Role": {
            "type": "object",
            "properties": {
                "display_title": {
                    "type": "string",
                    "example": "UserTitle"
                },
                "permissions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "prefix_create",
                        "prefix_read"
                    ]
                }
            }
        },
        "RoleResponse": {
            "type": "object",
            "properties": {
                "display_title": {
                    "type": "string",
                    "example": "UserTitle"
                },
                "id": {
                    "type": "integer",
                    "example": 12
                },
                "permissions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/PermissionsResponse"
                    }
                },
                "title": {
                    "type": "string",
                    "example": "RoleTitle"
                }
            }
        },
        "RolesResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/RoleResponse"
                    }
                }
            }
        },
        "SignInRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "john.doe@example.com"
                },
                "password": {
                    "type": "string",
                    "example": "qwerty"
                }
            }
        },
        "SignUpRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "john.doe@example.com"
                },
                "id": {
                    "type": "integer",
                    "example": 14
                },
                "name": {
                    "type": "string",
                    "example": "John Doe"
                },
                "password": {
                    "type": "string",
                    "example": "qwerty"
                },
                "password_confirmation": {
                    "type": "string",
                    "example": "qwerty"
                }
            }
        },
        "Status": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "example": "Everything is OK"
                }
            }
        },
        "UserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "jon_doe@gmail.com"
                },
                "id": {
                    "type": "integer",
                    "example": 14
                },
                "name": {
                    "type": "string",
                    "example": "Jon Doe"
                },
                "role": {
                    "$ref": "#/definitions/RoleResponse"
                },
                "tg_id": {
                    "type": "integer",
                    "example": 131231231231323
                }
            }
        },
        "UserUpdateRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "jon_doe@gmail.com"
                },
                "name": {
                    "type": "string",
                    "example": "Jon Doe"
                },
                "password": {
                    "type": "string",
                    "example": "qwerty"
                },
                "tg_id": {
                    "type": "integer",
                    "example": 131231231231323
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
	Host:             "localhost:8080/api",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Expense Application",
	Description:      "API for income and expense applications that allows you to receive a report for a certain period",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
