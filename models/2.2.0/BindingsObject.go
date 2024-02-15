
package models

// BindingsObject represents a BindingsObject model.
type BindingsObject struct {
  Http interface{}
  Ws interface{}
  Amqp interface{}
  Amqp1 interface{}
  Mqtt interface{}
  Mqtt5 interface{}
  Kafka interface{}
  Anypointmq interface{}
  Nats interface{}
  Jms interface{}
  Sns interface{}
  Sqs interface{}
  Stomp interface{}
  Redis interface{}
  Ibmmq interface{}
  AdditionalProperties map[string]interface{}
}