// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/api/v1/tasks": {
            "get": {
                "description": "get task list",
                "tags": [
                    "Task"
                ],
                "summary": "get task list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.Task"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/tasks/:task_id/runs": {
            "get": {
                "description": "get task run list",
                "tags": [
                    "Task"
                ],
                "summary": "get task run list",
                "responses": {}
            }
        },
        "/api/v1/tasks/:task_id/runs/:run_id": {
            "get": {
                "description": "get run detail",
                "tags": [
                    "Task"
                ],
                "summary": "get run detail",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Run"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "executor.ExeLog": {
            "type": "object",
            "properties": {
                "line": {
                    "description": "行内容",
                    "type": "string"
                },
                "output_at": {
                    "description": "输出时间",
                    "type": "string"
                },
                "type": {
                    "description": "内容类型，0标准输出，1标准错误输出",
                    "type": "integer"
                }
            }
        },
        "response.Run": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "exit_code": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "logs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/executor.ExeLog"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "response.Task": {
            "type": "object",
            "properties": {
                "bash": {
                    "type": "string"
                },
                "cron": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}
