
package models

// ServerBindingsObject represents a ServerBindingsObject model.
type ServerBindingsObject struct {
  Http interface{} `json:"http"`
  Ws interface{} `json:"ws"`
  Amqp interface{} `json:"amqp"`
  Amqp1 interface{} `json:"amqp1"`
  Mqtt *ServerBindingsObjectMqtt `json:"mqtt"`
  Kafka *ServerBindingsObjectKafka `json:"kafka"`
  Anypointmq interface{} `json:"anypointmq"`
  Nats interface{} `json:"nats"`
  Jms *ServerBindingsObjectJms `json:"jms"`
  Sns interface{} `json:"sns"`
  Sqs interface{} `json:"sqs"`
  Stomp interface{} `json:"stomp"`
  Redis interface{} `json:"redis"`
  Ibmmq *ServerBindingsObjectIbmmq `json:"ibmmq"`
  Solace *ServerBindingsObjectSolace `json:"solace"`
  Googlepubsub interface{} `json:"googlepubsub"`
  Pulsar *ServerBindingsObjectPulsar `json:"pulsar"`
  AdditionalProperties map[string]interface{} `json:"-"`
}