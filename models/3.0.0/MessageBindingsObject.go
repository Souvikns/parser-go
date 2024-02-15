
package models

// MessageBindingsObject represents a MessageBindingsObject model.
type MessageBindingsObject struct {
  Http *MessageBindingsObjectHttp
  Ws interface{}
  Amqp *MessageBindingsObjectAmqp
  Amqp1 interface{}
  Mqtt *MessageBindingsObjectMqtt
  Kafka *MessageBindingsObjectKafka
  Anypointmq *MessageBindingsObjectAnypointmq
  Nats interface{}
  Jms *MessageBindingsObjectJms
  Sns interface{}
  Sqs interface{}
  Stomp interface{}
  Redis interface{}
  Ibmmq *MessageBindingsObjectIbmmq
  Solace interface{}
  Googlepubsub *MessageBindingsObjectGooglepubsub
  AdditionalProperties map[string]interface{}
}