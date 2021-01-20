// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "117503445",
            "url": "https://github.com/117503445",
            "email": "t117503445@gmail.com"
        },
        "license": {
            "name": "GNU General Public License v3.0",
            "url": "https://github.com/TGclub/Wizz-Home-Page/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "更改请求中的 Username 和 Password 进行登录。登陆成功以后，返回json中token字段比如说是\"token\":\"123\"，就在右上角Authorize按钮点一下，输入Bearer 123，大小写、空格敏感。然后就能使用需要身份验证的接口啦。",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "身份验证"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "description": "登录值",
                        "name": "loginVals",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Middlewares.login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"expire\":\"2020-02-05T23:11:41+08:00\",\"token\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODA5MTU1MDEsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU4MDMxMDcwMX0.GWlmyTfCkXQYwgbtuTgVSTUSJXDcoDb_bptgRpt4HCU\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/backup": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "备份"
                ],
                "summary": "下载数据库"
            }
        },
        "/image/BackGroundImageUrls": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图片"
                ],
                "summary": "获取所有背景图片的url",
                "responses": {
                    "200": {
                        "description": "[\"http://q52qkptnh.bkt.clouddn.com/1.png\",\"http://q52qkptnh.bkt.clouddn.com/2.png\",\"http://q52qkptnh.bkt.clouddn.com/3.png\"]",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/image/PlaceAndDomain": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图片"
                ],
                "summary": "获取七牛云空间的地区和域名",
                "responses": {
                    "200": {
                        "description": "{\"domain\":\"q52qkptnh.bkt.clouddn.com\",\"place\": \"华东}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/image/UpToken": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图片"
                ],
                "summary": "获取一个上传文件的upToken",
                "parameters": [
                    {
                        "type": "string",
                        "description": "要上传的文件名,如 abc.png",
                        "name": "fileName",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "upToken",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/logs": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "日志"
                ],
                "summary": "分页获取日志",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 5,
                        "description": "一页含有的log数量",
                        "name": "pageSize",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "第几页",
                        "name": "pageIndex",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"LogTotalNum\":10,\"logs\":[{\"id\":2,\"TimeStamp\":1580620678,\"RequestMethod\":\"PUT\",\"Username\":\"admin\",\"RequestURI\":\"/api/products/1\",\"ResponseCode\":\"200\",\"ModelName\":\"一起来开黑\"}],\"pageTotalNum\":10}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/members": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "成员"
                ],
                "summary": "获取所有成员",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Member"
                            }
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
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "成员"
                ],
                "summary": "添加一个成员",
                "parameters": [
                    {
                        "description": "成员",
                        "name": "member",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httpModels.NoIdMember"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Member"
                        }
                    }
                }
            }
        },
        "/members/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "成员"
                ],
                "summary": "获取一个成员",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "成员id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Member"
                        }
                    },
                    "404": {
                        "description": "{\"message\":\"Member not found\"}",
                        "schema": {
                            "type": "string"
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
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "成员"
                ],
                "summary": "更改一个成员",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "成员id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "成员",
                        "name": "member",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httpModels.NoIdMember"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Member"
                        }
                    },
                    "404": {
                        "description": "{\"message\": \"Member not found\"}",
                        "schema": {
                            "type": "string"
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
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "成员"
                ],
                "summary": "删除一个成员",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "成员id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\": \"delete success\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "{\"message\": \"Member not found\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/passage": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "介绍"
                ],
                "summary": "添加一个介绍",
                "parameters": [
                    {
                        "description": "介绍",
                        "name": "passage",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httpModels.NoIdPassage"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Passage"
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "产品"
                ],
                "summary": "获取所有产品",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Product"
                            }
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
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "产品"
                ],
                "summary": "添加一个产品",
                "parameters": [
                    {
                        "description": "产品",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httpModels.NoIdProduct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    }
                }
            }
        },
        "/products/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "产品"
                ],
                "summary": "获取一个产品",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "产品id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    },
                    "404": {
                        "description": "{\"message\":\"Product not found\"}",
                        "schema": {
                            "type": "string"
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
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "产品"
                ],
                "summary": "更改一个产品",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "产品id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "产品",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httpModels.NoIdProduct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    },
                    "404": {
                        "description": "{\"message\": \"Product not found\"}",
                        "schema": {
                            "type": "string"
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
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "产品"
                ],
                "summary": "删除一个产品",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "产品id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\": \"delete success\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "{\"message\": \"Product not found\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/stories": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "历史事件"
                ],
                "summary": "获取所有历史事件",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Story"
                            }
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
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "历史事件"
                ],
                "summary": "添加一个历史事件",
                "parameters": [
                    {
                        "description": "历史事件",
                        "name": "story",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httpModels.NoIdStory"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Story"
                        }
                    }
                }
            }
        },
        "/stories/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "历史事件"
                ],
                "summary": "获取一个历史事件",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "历史事件id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Story"
                        }
                    },
                    "404": {
                        "description": "{\"message\":\"Story not found\"}",
                        "schema": {
                            "type": "string"
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
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "历史事件"
                ],
                "summary": "更改一个历史事件",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "历史事件id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "历史事件",
                        "name": "story",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httpModels.NoIdStory"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Story"
                        }
                    },
                    "404": {
                        "description": "{\"message\": \"Story not found\"}",
                        "schema": {
                            "type": "string"
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
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "历史事件"
                ],
                "summary": "删除一个历史事件",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "历史事件id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\": \"delete success\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "{\"message\": \"Story not found\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "Middlewares.login": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "example": "admin"
                },
                "username": {
                    "type": "string",
                    "example": "admin"
                }
            }
        },
        "httpModels.NoIdMember": {
            "type": "object",
            "properties": {
                "describe": {
                    "description": "个人简介",
                    "type": "string",
                    "example": "超级帅的前端dalao"
                },
                "memberType": {
                    "description": "成员类型,0 - 学生,1 - 导师",
                    "type": "integer",
                    "enum": [
                        0,
                        1
                    ],
                    "example": 0
                },
                "name": {
                    "description": "成员姓名",
                    "type": "string",
                    "example": "tengye"
                },
                "schoolYear": {
                    "description": "入学年份",
                    "type": "integer",
                    "example": 2017
                },
                "teacherInfo": {
                    "description": "如果是老师,则有若干头衔",
                    "type": "string",
                    "example": "创新校长\n制霸西电"
                },
                "urlAvatar": {
                    "description": "成员头像图片的url",
                    "type": "string",
                    "example": "http://wuygewfuyd/weiug.jpg"
                }
            }
        },
        "httpModels.NoIdPassage": {
            "type": "object",
            "properties": {
                "passage": {
                    "type": "string",
                    "example": "这是为之介绍"
                }
            }
        },
        "httpModels.NoIdProduct": {
            "type": "object",
            "properties": {
                "describe": {
                    "description": "项目介绍",
                    "type": "string",
                    "example": "大学生喜闻乐见的跨校交友平台，一键匹配聊天约游戏，一起来开黑帮你聊天\u0026游戏两不误~"
                },
                "littleDescribe": {
                    "description": "一句话介绍",
                    "type": "string",
                    "example": "一键约游戏"
                },
                "name": {
                    "description": "产品名称",
                    "type": "string",
                    "example": "一起来开黑"
                },
                "partner": {
                    "description": "合作方",
                    "type": "string",
                    "example": "腾讯"
                },
                "projectType": {
                    "description": "项目类型,0 - 校企合作,1 - 校园合作,2 - 校内自研",
                    "type": "integer",
                    "enum": [
                        0,
                        1
                    ]
                },
                "urlAvatar": {
                    "description": "项目图标的 Url",
                    "type": "string",
                    "example": "http://qiniu.com/UrlAvatar.png"
                },
                "urlBackground": {
                    "description": "项目背景图片的 Url",
                    "type": "string",
                    "example": "http://qiniu.com/UrlBackground.png"
                },
                "urlPartnerLogo": {
                    "description": "合作方Logo的 Url",
                    "type": "string",
                    "example": "http://qiniu.com/UrlPartnerLogo.png"
                },
                "urlProCode": {
                    "description": "二维码的 Url",
                    "type": "string",
                    "example": "http://qiniu.com/UrlProCode.png"
                },
                "urlScreenshot": {
                    "description": "项目截图的 Url",
                    "type": "string",
                    "example": "http://qiniu.com/UrlScreenshot.png"
                }
            }
        },
        "httpModels.NoIdStory": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "为之诞生"
                },
                "storyDescribe": {
                    "type": "string",
                    "example": "11月,前身 TgClub 诞生"
                },
                "timeStamp": {
                    "type": "integer",
                    "example": 1580397149
                }
            }
        },
        "models.Member": {
            "type": "object",
            "properties": {
                "describe": {
                    "description": "个人简介",
                    "type": "string",
                    "example": "超级帅的前端dalao"
                },
                "id": {
                    "description": "member's id",
                    "type": "integer",
                    "example": 1
                },
                "memberType": {
                    "description": "成员类型,0 - 学生,1 - 导师",
                    "type": "integer",
                    "enum": [
                        0,
                        1
                    ],
                    "example": 0
                },
                "name": {
                    "description": "成员姓名",
                    "type": "string",
                    "example": "tengye"
                },
                "schoolYear": {
                    "description": "入学年份",
                    "type": "integer",
                    "example": 2017
                },
                "teacherInfo": {
                    "description": "如果是老师,则有若干头衔",
                    "type": "string",
                    "example": "创新校长\n制霸西电"
                },
                "urlAvatar": {
                    "description": "成员头像图片的url",
                    "type": "string",
                    "example": "http://wuygewfuyd/weiug.jpg"
                }
            }
        },
        "models.Passage": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID is passage's id",
                    "type": "integer",
                    "example": 1
                },
                "passage": {
                    "type": "string",
                    "example": "这是为之介绍"
                }
            }
        },
        "models.Product": {
            "type": "object",
            "properties": {
                "describe": {
                    "description": "项目介绍",
                    "type": "string",
                    "example": "大学生喜闻乐见的跨校交友平台，一键匹配聊天约游戏，一起来开黑帮你聊天\u0026游戏两不误~"
                },
                "id": {
                    "description": "product's id",
                    "type": "integer",
                    "example": 1
                },
                "littleDescribe": {
                    "description": "一句话介绍",
                    "type": "string",
                    "example": "一键约游戏"
                },
                "name": {
                    "description": "产品名称",
                    "type": "string",
                    "example": "一起来开黑"
                },
                "partner": {
                    "description": "合作方",
                    "type": "string",
                    "example": "腾讯"
                },
                "projectType": {
                    "description": "项目类型,0 - 校企合作,1 - 校园合作,2 - 校内自研",
                    "type": "integer",
                    "enum": [
                        0,
                        1
                    ]
                },
                "urlAvatar": {
                    "description": "项目图标的 Url",
                    "type": "string",
                    "example": "http://qiniu.com/UrlAvatar.png"
                },
                "urlBackground": {
                    "description": "项目背景图片的 Url",
                    "type": "string",
                    "example": "http://qiniu.com/UrlBackground.png"
                },
                "urlPartnerLogo": {
                    "description": "合作方Logo的 Url",
                    "type": "string",
                    "example": "http://qiniu.com/UrlPartnerLogo.png"
                },
                "urlProCode": {
                    "description": "二维码的 Url",
                    "type": "string",
                    "example": "http://qiniu.com/UrlProCode.png"
                },
                "urlScreenshot": {
                    "description": "项目截图的 Url",
                    "type": "string",
                    "example": "http://qiniu.com/UrlScreenshot.png"
                }
            }
        },
        "models.Story": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID is story's id",
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "为之诞生"
                },
                "storyDescribe": {
                    "type": "string",
                    "example": "11月,前身 TgClub 诞生"
                },
                "timeStamp": {
                    "type": "integer",
                    "example": 1580397149
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
	Version:     "1.0",
	Host:        "homepage.backend.wizzstudio.com",
	BasePath:    "/api",
	Schemes:     []string{"https"},
	Title:       "Wizz-Home-Page API",
	Description: "Wizz's HomePage Backend",
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
