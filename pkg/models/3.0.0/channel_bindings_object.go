
package models

// ChannelBindingsObject represents a ChannelBindingsObject model.
type ChannelBindingsObject struct {
  Http interface{} `json:"http,omitempty"`
  Ws *ChannelBindingsObjectWs `json:"ws,omitempty"`
  Amqp *ChannelBindingsObjectAmqp `json:"amqp,omitempty"`
  Amqp1 interface{} `json:"amqp1,omitempty"`
  Mqtt interface{} `json:"mqtt,omitempty"`
  Kafka *ChannelBindingsObjectKafka `json:"kafka,omitempty"`
  Anypointmq *ChannelBindingsObjectAnypointmq `json:"anypointmq,omitempty"`
  Nats interface{} `json:"nats,omitempty"`
  Jms *ChannelBindingsObjectJms `json:"jms,omitempty"`
  Sns *ChannelBindingsObjectSns `json:"sns,omitempty"`
  Sqs *ChannelBindingsObjectSqs `json:"sqs,omitempty"`
  Stomp interface{} `json:"stomp,omitempty"`
  Redis interface{} `json:"redis,omitempty"`
  Ibmmq *ChannelBindingsObjectIbmmq `json:"ibmmq,omitempty"`
  Solace interface{} `json:"solace,omitempty"`
  Googlepubsub *ChannelBindingsObjectGooglepubsub `json:"googlepubsub,omitempty"`
  Pulsar *ChannelBindingsObjectPulsar `json:"pulsar,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}