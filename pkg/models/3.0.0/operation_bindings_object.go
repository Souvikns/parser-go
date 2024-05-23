
package models

// OperationBindingsObject represents a OperationBindingsObject model.
type OperationBindingsObject struct {
  Http *OperationBindingsObjectHttp `json:"http,omitempty"`
  Ws interface{} `json:"ws,omitempty"`
  Amqp *OperationBindingsObjectAmqp `json:"amqp,omitempty"`
  Amqp1 interface{} `json:"amqp1,omitempty"`
  Mqtt *OperationBindingsObjectMqtt `json:"mqtt,omitempty"`
  Kafka *OperationBindingsObjectKafka `json:"kafka,omitempty"`
  Anypointmq interface{} `json:"anypointmq,omitempty"`
  Nats *OperationBindingsObjectNats `json:"nats,omitempty"`
  Jms interface{} `json:"jms,omitempty"`
  Sns *OperationBindingsObjectSns `json:"sns,omitempty"`
  Sqs *OperationBindingsObjectSqs `json:"sqs,omitempty"`
  Stomp interface{} `json:"stomp,omitempty"`
  Redis interface{} `json:"redis,omitempty"`
  Ibmmq interface{} `json:"ibmmq,omitempty"`
  Solace *OperationBindingsObjectSolace `json:"solace,omitempty"`
  Googlepubsub interface{} `json:"googlepubsub,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}