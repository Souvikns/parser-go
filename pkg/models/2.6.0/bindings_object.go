
package models

// BindingsObject represents a BindingsObject model.
type BindingsObject struct {
  Http interface{} `json:"http,omitempty"`
  Ws interface{} `json:"ws,omitempty"`
  Amqp interface{} `json:"amqp,omitempty"`
  Amqp1 interface{} `json:"amqp1,omitempty"`
  Mqtt interface{} `json:"mqtt,omitempty"`
  Mqtt5 interface{} `json:"mqtt5,omitempty"`
  Kafka interface{} `json:"kafka,omitempty"`
  Anypointmq interface{} `json:"anypointmq,omitempty"`
  Nats interface{} `json:"nats,omitempty"`
  Jms interface{} `json:"jms,omitempty"`
  Sns interface{} `json:"sns,omitempty"`
  Sqs interface{} `json:"sqs,omitempty"`
  Stomp interface{} `json:"stomp,omitempty"`
  Redis interface{} `json:"redis,omitempty"`
  Ibmmq interface{} `json:"ibmmq,omitempty"`
  Solace interface{} `json:"solace,omitempty"`
  Googlepubsub interface{} `json:"googlepubsub,omitempty"`
  Pulsar interface{} `json:"pulsar,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}