// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a sample server Petstore server.  You can find out more about     Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).      For this sample, you can use the api key ` + "`" + `special-key` + "`" + ` to test the authorization     filters.",
    "title": "Swagger Petstore",
    "termsOfService": "http://swagger.io/terms/",
    "contact": {
      "email": "apiteam@swagger.io"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "127.0.0.1:9999",
  "basePath": "/v2",
  "paths": {
    "/pet": {
      "put": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Update an existing pet",
        "operationId": "updatePet",
        "parameters": [
          {
            "description": "Pet object that needs to be added to the store",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Pet"
            }
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Pet not found"
          },
          "405": {
            "description": "Validation exception"
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Add a new pet to the store",
        "operationId": "addPet",
        "parameters": [
          {
            "description": "Pet object that needs to be added to the store",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Pet"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successfully added",
            "schema": {
              "$ref": "#/definitions/Pet"
            }
          },
          "405": {
            "description": "Invalid input"
          }
        }
      }
    },
    "/pet/{petId}": {
      "get": {
        "description": "Returns a single pet",
        "produces": [
          "application/json"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Find pet by ID",
        "operationId": "getPetById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of pet to return",
            "name": "petId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Pet"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Pet not found"
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Deletes a pet",
        "operationId": "deletePet",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "Pet id to delete",
            "name": "petId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "success"
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Pet not found"
          }
        }
      }
    }
  },
  "definitions": {
    "ApiResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "Pet": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "favoriteFoods": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string",
          "example": "doggie"
        },
        "nicknames": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    }
  },
  "tags": [
    {
      "description": "Everything about your Pets",
      "name": "pet",
      "externalDocs": {
        "description": "Find out more",
        "url": "http://swagger.io"
      }
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a sample server Petstore server.  You can find out more about     Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).      For this sample, you can use the api key ` + "`" + `special-key` + "`" + ` to test the authorization     filters.",
    "title": "Swagger Petstore",
    "termsOfService": "http://swagger.io/terms/",
    "contact": {
      "email": "apiteam@swagger.io"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "127.0.0.1:9999",
  "basePath": "/v2",
  "paths": {
    "/pet": {
      "put": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Update an existing pet",
        "operationId": "updatePet",
        "parameters": [
          {
            "description": "Pet object that needs to be added to the store",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Pet"
            }
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Pet not found"
          },
          "405": {
            "description": "Validation exception"
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Add a new pet to the store",
        "operationId": "addPet",
        "parameters": [
          {
            "description": "Pet object that needs to be added to the store",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Pet"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successfully added",
            "schema": {
              "$ref": "#/definitions/Pet"
            }
          },
          "405": {
            "description": "Invalid input"
          }
        }
      }
    },
    "/pet/{petId}": {
      "get": {
        "description": "Returns a single pet",
        "produces": [
          "application/json"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Find pet by ID",
        "operationId": "getPetById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of pet to return",
            "name": "petId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Pet"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Pet not found"
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Deletes a pet",
        "operationId": "deletePet",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "Pet id to delete",
            "name": "petId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "success"
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Pet not found"
          }
        }
      }
    }
  },
  "definitions": {
    "ApiResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "Pet": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "favoriteFoods": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string",
          "example": "doggie"
        },
        "nicknames": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    }
  },
  "tags": [
    {
      "description": "Everything about your Pets",
      "name": "pet",
      "externalDocs": {
        "description": "Find out more",
        "url": "http://swagger.io"
      }
    }
  ]
}`))
}
