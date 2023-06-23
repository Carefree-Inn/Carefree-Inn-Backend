// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/comment": {
            "get": {
                "description": "获取帖子下的评论",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "获取帖子下的评论 api",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户token",
                        "name": "Authorzation",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "帖子id",
                        "name": "post_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "条数",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "评论",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "评论 api",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户token",
                        "name": "Authorzation",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "评论信息",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/comment.makeCommentRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除评论",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "删除评论 api",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户token",
                        "name": "Authorzation",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "评论信息",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/comment.deleteCommentRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal.Response"
                        }
                    }
                }
            }
        },
        "/like": {
            "post": {
                "description": "点赞",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "like"
                ],
                "summary": "点赞 api",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户token",
                        "name": "Authorzation",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "被点赞帖子相关信息",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/like.makeLikeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "取消点赞",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "like"
                ],
                "summary": "取消点赞 api",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户token",
                        "name": "Authorzation",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "帖子id",
                        "name": "post_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal.Response"
                        }
                    }
                }
            }
        },
        "/notification": {
            "get": {
                "description": "获取用户通知",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "notification"
                ],
                "summary": "获取用户通知（被点赞/评论） api",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户token",
                        "name": "Authorzation",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal.Response"
                        }
                    }
                }
            }
        },
        "/post": {
            "post": {
                "description": "创建帖子",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "创建帖子 api",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户token",
                        "name": "Authorzation",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "帖子信息",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/post.createPostRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除帖子",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "删除帖子 api",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户token",
                        "name": "Authorzation",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "帖子信息",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/post.deletePostRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal.Response"
                        }
                    }
                }
            }
        },
        "/post/category": {
            "get": {
                "description": "获取分区帖子信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "获取分区帖子信息 api",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "条数",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "description": "分类信息",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/post.getPostOfCategoryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal.Response"
                        }
                    }
                }
            }
        },
        "/post/category/all": {
            "get": {
                "description": "获取分区信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "获取分区信息 api",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal.Response"
                        }
                    }
                }
            }
        },
        "/post/search": {
            "post": {
                "description": "搜索帖子",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "搜索帖子 api",
                "parameters": [
                    {
                        "description": "搜索信息",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/post.searchRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal.Response"
                        }
                    }
                }
            }
        },
        "/post/tag": {
            "post": {
                "description": "获取tag下的帖子",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "获取tag下的帖子 api",
                "parameters": [
                    {
                        "description": "tag信息",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/post.getPostOfTag"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal.Response"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "用户通过学号登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "登陆 api",
                "parameters": [
                    {
                        "description": "用户信息",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.user"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal.Response"
                        }
                    }
                }
            }
        },
        "/user/profile": {
            "get": {
                "description": "获取用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "获取用户信息 api",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户token",
                        "name": "Authorzation",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal.Response"
                        }
                    }
                }
            },
            "put": {
                "description": "修改用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "修改用户信息 api",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户token",
                        "name": "Authorzation",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "需要修改的信息",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.userInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "comment.deleteCommentRequest": {
            "type": "object",
            "properties": {
                "comment_id": {
                    "type": "integer"
                }
            }
        },
        "comment.makeCommentRequest": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "from_user_id": {
                    "type": "string"
                },
                "is_top": {
                    "type": "boolean"
                },
                "post_id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "to_user_id": {
                    "type": "string"
                },
                "top_comment_id": {
                    "type": "integer"
                }
            }
        },
        "internal.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "like.makeLikeRequest": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "post_id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "post.createPostRequest": {
            "type": "object",
            "required": [
                "category",
                "content",
                "title"
            ],
            "properties": {
                "category": {
                    "type": "object",
                    "properties": {
                        "category_id": {
                            "type": "integer"
                        },
                        "title": {
                            "type": "string"
                        }
                    }
                },
                "content": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "tag": {
                                "type": "string"
                            }
                        }
                    }
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "post.deletePostRequest": {
            "type": "object",
            "required": [
                "post_id"
            ],
            "properties": {
                "post_id": {
                    "type": "integer"
                }
            }
        },
        "post.getPostOfCategoryRequest": {
            "type": "object",
            "required": [
                "category_id"
            ],
            "properties": {
                "category_id": {
                    "type": "integer"
                }
            }
        },
        "post.getPostOfTag": {
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
        "post.searchRequest": {
            "type": "object",
            "required": [
                "data",
                "search_type"
            ],
            "properties": {
                "data": {
                    "type": "string"
                },
                "search_type": {
                    "type": "string"
                }
            }
        },
        "user.user": {
            "type": "object",
            "required": [
                "account",
                "password"
            ],
            "properties": {
                "account": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "user.userInfo": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                },
                "avatar": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "sex": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "139.196.30.123",
	BasePath:         "/inn/api/v1",
	Schemes:          []string{},
	Title:            "Inn",
	Description:      "Inn",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
