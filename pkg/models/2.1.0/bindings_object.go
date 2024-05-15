
package models

// BindingsObject represents a BindingsObject model.
type BindingsObject struct {
  Http interface{} `json:"http"`
  Ws interface{} `json:"ws"`
  Amqp interface{} `json:"amqp"`
  Amqp1 interface{} `json:"amqp1"`
  Mqtt interface{} `json:"mqtt"`
  Mqtt5 interface{} `json:"mqtt5"`
  Kafka interface{} `json:"kafka"`
  Nats interface{} `json:"nats"`
  Jms interface{} `json:"jms"`
  Sns interface{} `json:"sns"`
  Sqs interface{} `json:"sqs"`
  Stomp interface{} `json:"stomp"`
  Redis interface{} `json:"redis"`
  Ibmmq interface{} `json:"ibmmq"`
  AdditionalProperties map[string]interface{} `json:"-"`
}