
package models

// ChannelBindingsObject represents a ChannelBindingsObject model.
type ChannelBindingsObject struct {
  Http interface{} `json:"http"`
  Ws *ChannelBindingsObjectWs `json:"ws"`
  Amqp *ChannelBindingsObjectAmqp `json:"amqp"`
  Amqp1 interface{} `json:"amqp1"`
  Mqtt interface{} `json:"mqtt"`
  Kafka *ChannelBindingsObjectKafka `json:"kafka"`
  Anypointmq *ChannelBindingsObjectAnypointmq `json:"anypointmq"`
  Nats interface{} `json:"nats"`
  Jms *ChannelBindingsObjectJms `json:"jms"`
  Sns *ChannelBindingsObjectSns `json:"sns"`
  Sqs *ChannelBindingsObjectSqs `json:"sqs"`
  Stomp interface{} `json:"stomp"`
  Redis interface{} `json:"redis"`
  Ibmmq *ChannelBindingsObjectIbmmq `json:"ibmmq"`
  Solace interface{} `json:"solace"`
  Googlepubsub *ChannelBindingsObjectGooglepubsub `json:"googlepubsub"`
  Pulsar *ChannelBindingsObjectPulsar `json:"pulsar"`
  AdditionalProperties map[string]interface{} `json:"-"`
}