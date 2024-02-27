package parser

import (
	"testing"

	"github.com/Souvikns/parser-go/scheme"
	"github.com/stretchr/testify/assert"
)

const Spec = `
{
	"asyncapi": "3.0.0",
	"info": {
	  "title": "Account Service",
	  "version": "1.0.0",
	  "description": "This service is in charge of processing user signups"
	},
	"servers": {
	  "mosquitto": {
		"host": "test.mosquitto.org",
		"protocol": "mqtts",
		"security": [
		  {
			"$ref": "#/components/securitySchemes/X509Certificate"
		  }
		],
		"bindings": {
		  "mqtt": {
			"clientId": "guest",
			"cleanSession": true,
			"keepAlive": 60,
			"bindingVersion": "0.2.0"
		  }
		}
	  },
	  "websockets": {
		"host": "localhost:3005",
		"protocol": "ws"
	  }
	},
	"channels": {
	  "userSignedUp": {
		"address": "/user/signedup",
		"messages": {
		  "UserSignedUp": {
			"$ref": "#/components/messages/UserSignedUp"
		  }
		},
		"bindings": {
		  "ws": {
			"bindingVersion": "0.1.0",
			"query": {
			  "type": "object",
			  "additionalProperties": false,
			  "properties": {
				"salutation": {
				  "type": "string",
				  "default": "Hello"
				}
			  }
			},
			"headers": {
			  "type": "object",
			  "properties": {
				"my-custom-header": {
				  "type": "string",
				  "const": "custom value"
				}
			  }
			}
		  }
		}
	  },
	  "serverAnnounce": {
		"address": "server/announce",
		"messages": {
		  "ServerAnnounce": {
			"$ref": "#/components/messages/ServerAnnounce"
		  }
		}
	  }
	},
	"operations": {
	  "recieveUserSignedUp": {
		"action": "receive",
		"channel": {
		  "$ref": "#/channels/userSignedUp"
		},
		"bindings": {
		  "mqtt": {
			"qos": 2,
			"retain": true,
			"bindingVersion": "0.2.0"
		  }
		},
		"messages": [
		  {
			"$ref": "#/channels/userSignedUp/messages/UserSignedUp"
		  }
		]
	  },
	  "sendSignedUpUser": {
		"action": "send",
		"channel": {
		  "$ref": "#/channels/userSignedUp"
		},
		"messages": [
		  {
			"$ref": "#/channels/userSignedUp/messages/UserSignedUp"
		  }
		]
	  },
	  "sendServerAnnounce": {
		"action": "send",
		"channel": {
		  "$ref": "#/channels/serverAnnounce"
		},
		"messages": [
		  {
			"$ref": "#/channels/serverAnnounce/messages/ServerAnnounce"
		  }
		]
	  }
	},
	"components": {
	  "securitySchemes": {
		"userAndPassword": {
		  "type": "userPassword"
		},
		"X509Certificate": {
		  "type": "X509"
		}
	  },
	  "messages": {
		"UserSignedUp": {
		  "payload": {
			"type": "object",
			"additionalProperties": false,
			"required": [
			  "displayName",
			  "email"
			],
			"properties": {
			  "displayName": {
				"type": "string",
				"description": "Name of the user"
			  },
			  "email": {
				"type": "string",
				"format": "email",
				"description": "Email of the user"
			  }
			}
		  }
		},
		"EmailSent": {
		  "payload": {
			"type": "object",
			"additionalProperties": false,
			"required": [
			  "email",
			  "timestamp"
			],
			"properties": {
			  "email": {
				"type": "string",
				"format": "email",
				"description": "Email of the user"
			  },
			  "timestamp": {
				"type": "string",
				"format": "date-time"
			  }
			}
		  }
		},
		"ServerAnnounce": {
		  "payload": {
			"type": "object",
			"additionalProperties": false,
			"required": [
			  "id"
			],
			"properties": {
			  "id": {
				"type": "string"
			  }
			}
		  }
		}
	  }
	}
  }
`

func TestLoadDocument(t *testing.T) {
	got, _ := loadDocument([]byte(Spec))
	want := "3.0.0"
	assert.Contains(t, got["asyncapi"], want)
}

func TestExtractVersion(t *testing.T) {
	doc, _ := loadDocument([]byte(Spec))
	got, _ := extractVersion(doc)
	want := "3.0.0"
	assert.Contains(t, got, want)
}

func TestValidateScheme(t *testing.T) {
	doc, _ := loadDocument([]byte(Spec))
	version, _ := extractVersion(doc)
	schemes, _ := scheme.LoadSchemes()
	schema, _ := schemes.Get(version)
	err := validateSchema(schema.Definition, Spec)
	assert.Nil(t, err)
}