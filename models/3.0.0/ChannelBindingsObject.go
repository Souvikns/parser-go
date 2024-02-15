
package models

// ChannelBindingsObject represents a ChannelBindingsObject model.
type ChannelBindingsObject struct {
  Http interface{}
  Ws *ChannelBindingsObjectWs
  Amqp *ChannelBindingsObjectAmqp
  Amqp1 interface{}
  Mqtt interface{}
  Kafka *ChannelBindingsObjectKafka
  Anypointmq *ChannelBindingsObjectAnypointmq
  Nats interface{}
  Jms *ChannelBindingsObjectJms
  Sns *ChannelBindingsObjectSns
  Sqs *ChannelBindingsObjectSqs
  Stomp interface{}
  Redis interface{}
  Ibmmq *ChannelBindingsObjectIbmmq
  Solace interface{}
  Googlepubsub *ChannelBindingsObjectGooglepubsub
  Pulsar *ChannelBindingsObjectPulsar
  AdditionalProperties map[string]interface{}
}