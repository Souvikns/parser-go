
package models

// OperationBindingsObject represents a OperationBindingsObject model.
type OperationBindingsObject struct {
  Http *OperationBindingsObjectHttp
  Ws interface{}
  Amqp *OperationBindingsObjectAmqp
  Amqp1 interface{}
  Mqtt *OperationBindingsObjectMqtt
  Kafka *OperationBindingsObjectKafka
  Anypointmq interface{}
  Nats *OperationBindingsObjectNats
  Jms interface{}
  Sns *OperationBindingsObjectSns
  Sqs *OperationBindingsObjectSqs
  Stomp interface{}
  Redis interface{}
  Ibmmq interface{}
  Solace *OperationBindingsObjectSolace
  Googlepubsub interface{}
  AdditionalProperties map[string]interface{}
}