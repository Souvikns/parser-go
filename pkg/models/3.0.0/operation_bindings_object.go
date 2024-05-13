
package models

// OperationBindingsObject represents a OperationBindingsObject model.
type OperationBindingsObject struct {
  Http *OperationBindingsObjectHttp `json:"http"`
  Ws interface{} `json:"ws"`
  Amqp *OperationBindingsObjectAmqp `json:"amqp"`
  Amqp1 interface{} `json:"amqp1"`
  Mqtt *OperationBindingsObjectMqtt `json:"mqtt"`
  Kafka *OperationBindingsObjectKafka `json:"kafka"`
  Anypointmq interface{} `json:"anypointmq"`
  Nats *OperationBindingsObjectNats `json:"nats"`
  Jms interface{} `json:"jms"`
  Sns *OperationBindingsObjectSns `json:"sns"`
  Sqs *OperationBindingsObjectSqs `json:"sqs"`
  Stomp interface{} `json:"stomp"`
  Redis interface{} `json:"redis"`
  Ibmmq interface{} `json:"ibmmq"`
  Solace *OperationBindingsObjectSolace `json:"solace"`
  Googlepubsub interface{} `json:"googlepubsub"`
  AdditionalProperties map[string]interface{} `json:"-"`
}