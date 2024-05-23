
package models

// ServerBindingsObject represents a ServerBindingsObject model.
type ServerBindingsObject struct {
  Http interface{} `json:"http,omitempty"`
  Ws interface{} `json:"ws,omitempty"`
  Amqp interface{} `json:"amqp,omitempty"`
  Amqp1 interface{} `json:"amqp1,omitempty"`
  Mqtt *ServerBindingsObjectMqtt `json:"mqtt,omitempty"`
  Kafka *ServerBindingsObjectKafka `json:"kafka,omitempty"`
  Anypointmq interface{} `json:"anypointmq,omitempty"`
  Nats interface{} `json:"nats,omitempty"`
  Jms *ServerBindingsObjectJms `json:"jms,omitempty"`
  Sns interface{} `json:"sns,omitempty"`
  Sqs interface{} `json:"sqs,omitempty"`
  Stomp interface{} `json:"stomp,omitempty"`
  Redis interface{} `json:"redis,omitempty"`
  Ibmmq *ServerBindingsObjectIbmmq `json:"ibmmq,omitempty"`
  Solace *ServerBindingsObjectSolace `json:"solace,omitempty"`
  Googlepubsub interface{} `json:"googlepubsub,omitempty"`
  Pulsar *ServerBindingsObjectPulsar `json:"pulsar,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}