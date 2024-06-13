// Package swagger Code generated by swaggo/swag. DO NOT EDIT
package swagger

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://github.com/sreok/kube-go",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/deployments": {
            "get": {
                "description": "获取Deployments列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kubernetes"
                ],
                "summary": "获取Deployments列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "命名空间",
                        "name": "namespace",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/api/gateways": {
            "get": {
                "description": "获取gateway信息",
                "tags": [
                    "Kubernetes"
                ],
                "summary": "获取gateway信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "命名空间",
                        "name": "namespace",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/api/groups": {
            "get": {
                "description": "获取APIGroups信息",
                "tags": [
                    "Kubernetes"
                ],
                "summary": "获取APIGroups信息",
                "responses": {}
            }
        },
        "/api/namespaces": {
            "get": {
                "description": "获取命名空间",
                "tags": [
                    "Kubernetes"
                ],
                "summary": "获取命名空间",
                "responses": {}
            }
        },
        "/api/nodes": {
            "get": {
                "description": "获取node列表",
                "tags": [
                    "Kubernetes"
                ],
                "summary": "获取node列表",
                "responses": {}
            }
        },
        "/api/pod/list": {
            "get": {
                "description": "获取pod列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kubernetes"
                ],
                "summary": "获取pod列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "命名空间",
                        "name": "namespace",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/api/services": {
            "get": {
                "description": "获取Service列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kubernetes"
                ],
                "summary": "获取Service列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "命名空间",
                        "name": "namespace",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/api/vm/list": {
            "get": {
                "description": "获取vm列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kubernetes"
                ],
                "summary": "获取vm列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "命名空间",
                        "name": "namespace",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/api/vmi/list": {
            "get": {
                "description": "获取vmi列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kubernetes"
                ],
                "summary": "获取vmi列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "命名空间",
                        "name": "namespace",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "kube-go",
	Description:      "基于client-go和gin框架后端api服务",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
