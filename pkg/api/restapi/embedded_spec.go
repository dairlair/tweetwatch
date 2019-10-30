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
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Tweetwatch Server",
    "version": "0.0.1"
  },
  "paths": {
    "/login": {
      "post": {
        "security": [],
        "operationId": "login",
        "parameters": [
          {
            "description": "New User",
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Credentials"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Logged in successfully",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "422": {
            "description": "Invalid credentials",
            "schema": {
              "$ref": "#/definitions/DefaultError"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/DefaultError"
            }
          }
        }
      }
    },
    "/signup": {
      "post": {
        "security": [],
        "operationId": "signup",
        "parameters": [
          {
            "description": "New User",
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Credentials"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "User created",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "422": {
            "description": "Email already taken",
            "schema": {
              "$ref": "#/definitions/DefaultError"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/DefaultError"
            }
          }
        }
      }
    },
    "/status": {
      "get": {
        "operationId": "getStatus",
        "responses": {
          "200": {
            "description": "Current JWT is valid",
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        }
      }
    },
    "/topics": {
      "get": {
        "operationId": "getUserTopics",
        "responses": {
          "200": {
            "description": "Returns topics list of current user",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Topic"
              }
            }
          }
        }
      },
      "post": {
        "operationId": "createTopic",
        "parameters": [
          {
            "description": "New Topic",
            "name": "topic",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/CreateTopic"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Topic created",
            "schema": {
              "$ref": "#/definitions/Topic"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/DefaultError"
            }
          }
        }
      }
    },
    "/topics/{topicId}": {
      "put": {
        "summary": "Update desired topic by Topic ID",
        "operationId": "updateTopic",
        "parameters": [
          {
            "type": "integer",
            "description": "Numeric ID of the topic to update",
            "name": "topicId",
            "in": "path",
            "required": true
          },
          {
            "description": "Updated topic data",
            "name": "topic",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/CreateTopic"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Topic updated",
            "schema": {
              "$ref": "#/definitions/Topic"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/DefaultError"
            }
          }
        }
      },
      "delete": {
        "summary": "Delete desired topic by Topic ID",
        "operationId": "deleteTopic",
        "parameters": [
          {
            "type": "integer",
            "description": "Numeric ID of the topic to delete",
            "name": "topicId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Topic deleted",
            "schema": {
              "$ref": "#/definitions/DefaultSuccess"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/DefaultError"
            }
          }
        }
      }
    },
    "/topics/{topicId}/streams": {
      "get": {
        "summary": "Returns list of streams inside the topic",
        "operationId": "getStreams",
        "parameters": [
          {
            "type": "integer",
            "name": "topicId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Streams list",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Stream"
              }
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/DefaultError"
            }
          }
        }
      },
      "post": {
        "operationId": "createStream",
        "parameters": [
          {
            "type": "integer",
            "description": "Numeric ID of the topic to create stream",
            "name": "topicId",
            "in": "path",
            "required": true
          },
          {
            "description": "Stream to create",
            "name": "stream",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/CreateStream"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Stream created",
            "schema": {
              "$ref": "#/definitions/Stream"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/DefaultError"
            }
          }
        }
      }
    },
    "/topics/{topicId}/streams/{streamId}": {
      "put": {
        "summary": "Update desired stream by Topic ID and Stream ID",
        "operationId": "updateStream",
        "parameters": [
          {
            "type": "integer",
            "description": "Numeric ID of the topic to update",
            "name": "topicId",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "description": "Numeric ID of the stream to update",
            "name": "streamId",
            "in": "path",
            "required": true
          },
          {
            "description": "Updated stream data",
            "name": "stream",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/CreateStream"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Stream updated",
            "schema": {
              "$ref": "#/definitions/Stream"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/DefaultError"
            }
          }
        }
      },
      "delete": {
        "summary": "Delete desired stream by Topic ID and Stream ID",
        "operationId": "deleteStream",
        "parameters": [
          {
            "type": "integer",
            "description": "Numeric ID of the topic to update",
            "name": "topicId",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "description": "Numeric ID of the stream to update",
            "name": "streamId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Stream deleted",
            "schema": {
              "$ref": "#/definitions/DefaultSuccess"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/DefaultError"
            }
          }
        }
      }
    },
    "/topics/{topicId}/tweets": {
      "get": {
        "summary": "Returns list of tweets retrieved for this topic",
        "operationId": "getTopicTweets",
        "parameters": [
          {
            "type": "integer",
            "name": "topicId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Tweets list",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Tweet"
              }
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/DefaultError"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "CreateStream": {
      "required": [
        "track"
      ],
      "properties": {
        "track": {
          "type": "string"
        }
      }
    },
    "CreateTopic": {
      "required": [
        "name",
        "isActive"
      ],
      "properties": {
        "isActive": {
          "type": "boolean"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "Credentials": {
      "required": [
        "email",
        "password"
      ],
      "properties": {
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string",
          "format": "password"
        }
      }
    },
    "DefaultError": {
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "DefaultSuccess": {
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "Stream": {
      "required": [
        "id",
        "track",
        "createdAt"
      ],
      "properties": {
        "createdAt": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "track": {
          "type": "string"
        }
      }
    },
    "Topic": {
      "required": [
        "id",
        "name",
        "createdAt",
        "isActive"
      ],
      "properties": {
        "createdAt": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "isActive": {
          "type": "boolean"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "Tweet": {
      "required": [
        "id",
        "twitteId",
        "fullText",
        "createdAt"
      ],
      "properties": {
        "createdAt": {
          "type": "string"
        },
        "fullText": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "twitteId": {
          "type": "integer"
        }
      }
    },
    "User": {
      "required": [
        "id",
        "email",
        "token"
      ],
      "properties": {
        "email": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "token": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "JWT": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    },
    "isRegistered": {
      "type": "basic"
    }
  },
  "security": [
    {
      "JWT": []
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Tweetwatch Server",
    "version": "0.0.1"
  },
  "paths": {
    "/login": {
      "post": {
        "security": [],
        "operationId": "login",
        "parameters": [
          {
            "description": "New User",
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Credentials"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Logged in successfully",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "422": {
            "description": "Invalid credentials",
            "schema": {
              "$ref": "#/definitions/DefaultError"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/DefaultError"
            }
          }
        }
      }
    },
    "/signup": {
      "post": {
        "security": [],
        "operationId": "signup",
        "parameters": [
          {
            "description": "New User",
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Credentials"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "User created",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "422": {
            "description": "Email already taken",
            "schema": {
              "$ref": "#/definitions/DefaultError"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/DefaultError"
            }
          }
        }
      }
    },
    "/status": {
      "get": {
        "operationId": "getStatus",
        "responses": {
          "200": {
            "description": "Current JWT is valid",
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        }
      }
    },
    "/topics": {
      "get": {
        "operationId": "getUserTopics",
        "responses": {
          "200": {
            "description": "Returns topics list of current user",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Topic"
              }
            }
          }
        }
      },
      "post": {
        "operationId": "createTopic",
        "parameters": [
          {
            "description": "New Topic",
            "name": "topic",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/CreateTopic"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Topic created",
            "schema": {
              "$ref": "#/definitions/Topic"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/DefaultError"
            }
          }
        }
      }
    },
    "/topics/{topicId}": {
      "put": {
        "summary": "Update desired topic by Topic ID",
        "operationId": "updateTopic",
        "parameters": [
          {
            "type": "integer",
            "description": "Numeric ID of the topic to update",
            "name": "topicId",
            "in": "path",
            "required": true
          },
          {
            "description": "Updated topic data",
            "name": "topic",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/CreateTopic"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Topic updated",
            "schema": {
              "$ref": "#/definitions/Topic"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/DefaultError"
            }
          }
        }
      },
      "delete": {
        "summary": "Delete desired topic by Topic ID",
        "operationId": "deleteTopic",
        "parameters": [
          {
            "type": "integer",
            "description": "Numeric ID of the topic to delete",
            "name": "topicId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Topic deleted",
            "schema": {
              "$ref": "#/definitions/DefaultSuccess"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/DefaultError"
            }
          }
        }
      }
    },
    "/topics/{topicId}/streams": {
      "get": {
        "summary": "Returns list of streams inside the topic",
        "operationId": "getStreams",
        "parameters": [
          {
            "type": "integer",
            "name": "topicId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Streams list",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Stream"
              }
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/DefaultError"
            }
          }
        }
      },
      "post": {
        "operationId": "createStream",
        "parameters": [
          {
            "type": "integer",
            "description": "Numeric ID of the topic to create stream",
            "name": "topicId",
            "in": "path",
            "required": true
          },
          {
            "description": "Stream to create",
            "name": "stream",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/CreateStream"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Stream created",
            "schema": {
              "$ref": "#/definitions/Stream"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/DefaultError"
            }
          }
        }
      }
    },
    "/topics/{topicId}/streams/{streamId}": {
      "put": {
        "summary": "Update desired stream by Topic ID and Stream ID",
        "operationId": "updateStream",
        "parameters": [
          {
            "type": "integer",
            "description": "Numeric ID of the topic to update",
            "name": "topicId",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "description": "Numeric ID of the stream to update",
            "name": "streamId",
            "in": "path",
            "required": true
          },
          {
            "description": "Updated stream data",
            "name": "stream",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/CreateStream"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Stream updated",
            "schema": {
              "$ref": "#/definitions/Stream"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/DefaultError"
            }
          }
        }
      },
      "delete": {
        "summary": "Delete desired stream by Topic ID and Stream ID",
        "operationId": "deleteStream",
        "parameters": [
          {
            "type": "integer",
            "description": "Numeric ID of the topic to update",
            "name": "topicId",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "description": "Numeric ID of the stream to update",
            "name": "streamId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Stream deleted",
            "schema": {
              "$ref": "#/definitions/DefaultSuccess"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/DefaultError"
            }
          }
        }
      }
    },
    "/topics/{topicId}/tweets": {
      "get": {
        "summary": "Returns list of tweets retrieved for this topic",
        "operationId": "getTopicTweets",
        "parameters": [
          {
            "type": "integer",
            "name": "topicId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Tweets list",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Tweet"
              }
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/DefaultError"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "CreateStream": {
      "required": [
        "track"
      ],
      "properties": {
        "track": {
          "type": "string"
        }
      }
    },
    "CreateTopic": {
      "required": [
        "name",
        "isActive"
      ],
      "properties": {
        "isActive": {
          "type": "boolean"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "Credentials": {
      "required": [
        "email",
        "password"
      ],
      "properties": {
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string",
          "format": "password"
        }
      }
    },
    "DefaultError": {
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "DefaultSuccess": {
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "Stream": {
      "required": [
        "id",
        "track",
        "createdAt"
      ],
      "properties": {
        "createdAt": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "track": {
          "type": "string"
        }
      }
    },
    "Topic": {
      "required": [
        "id",
        "name",
        "createdAt",
        "isActive"
      ],
      "properties": {
        "createdAt": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "isActive": {
          "type": "boolean"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "Tweet": {
      "required": [
        "id",
        "twitteId",
        "fullText",
        "createdAt"
      ],
      "properties": {
        "createdAt": {
          "type": "string"
        },
        "fullText": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "twitteId": {
          "type": "integer"
        }
      }
    },
    "User": {
      "required": [
        "id",
        "email",
        "token"
      ],
      "properties": {
        "email": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "token": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "JWT": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    },
    "isRegistered": {
      "type": "basic"
    }
  },
  "security": [
    {
      "JWT": []
    }
  ]
}`))
}
