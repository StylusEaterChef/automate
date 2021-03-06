package api

func init() {
	Swagger.Add("iam_v2_introspect", `{
  "swagger": "2.0",
  "info": {
    "title": "automate-gateway/api/iam/v2/introspect.proto",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/apis/iam/v2/introspect": {
      "get": {
        "operationId": "IntrospectAll",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.IntrospectResp"
            }
          }
        },
        "tags": [
          "Authorization"
        ]
      },
      "post": {
        "operationId": "Introspect",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.IntrospectResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.IntrospectReq"
            }
          }
        ],
        "tags": [
          "Authorization"
        ]
      }
    },
    "/apis/iam/v2/introspect_some": {
      "post": {
        "operationId": "IntrospectSome",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.IntrospectResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.IntrospectSomeReq"
            }
          }
        ],
        "tags": [
          "Authorization"
        ]
      }
    }
  },
  "definitions": {
    "chef.automate.api.iam.v2.IntrospectReq": {
      "type": "object",
      "properties": {
        "path": {
          "type": "string"
        },
        "parameters": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "chef.automate.api.iam.v2.IntrospectResp": {
      "type": "object",
      "properties": {
        "endpoints": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/chef.automate.api.iam.v2.MethodsAllowed"
          }
        }
      }
    },
    "chef.automate.api.iam.v2.IntrospectSomeReq": {
      "type": "object",
      "properties": {
        "paths": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "chef.automate.api.iam.v2.MethodsAllowed": {
      "type": "object",
      "properties": {
        "get": {
          "type": "boolean",
          "format": "boolean"
        },
        "put": {
          "type": "boolean",
          "format": "boolean"
        },
        "post": {
          "type": "boolean",
          "format": "boolean"
        },
        "delete": {
          "type": "boolean",
          "format": "boolean"
        },
        "patch": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    }
  }
}
`)
}
