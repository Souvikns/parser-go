
package models

// ServerBindingsObject represents a ServerBindingsObject model.
type ServerBindingsObject struct {
  Http interface{}
  Ws interface{}
  Amqp interface{}
  Amqp1 interface{}
  Mqtt *ServerBindingsObjectMqtt
  Kafka *ServerBindingsObjectKafka
  Anypointmq interface{}
  Nats interface{}
  Jms *ServerBindingsObjectJms
  Sns interface{}
  Sqs interface{}
  Stomp interface{}
  Redis interface{}
  Ibmmq *ServerBindingsObjectIbmmq
  Solace *ServerBindingsObjectSolace
  Googlepubsub interface{}
  Pulsar *ServerBindingsObjectPulsar
  AdditionalProperties map[string]interface{}
}