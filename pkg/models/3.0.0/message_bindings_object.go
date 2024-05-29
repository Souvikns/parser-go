
package models

// MessageBindingsObject represents a MessageBindingsObject model.
type MessageBindingsObject struct {
  Http *MessageBindingsObjectHttp `json:"http,omitempty"`
  Ws interface{} `json:"ws,omitempty"`
  Amqp *MessageBindingsObjectAmqp `json:"amqp,omitempty"`
  Amqp1 interface{} `json:"amqp1,omitempty"`
  Mqtt *MessageBindingsObjectMqtt `json:"mqtt,omitempty"`
  Kafka *MessageBindingsObjectKafka `json:"kafka,omitempty"`
  Anypointmq *MessageBindingsObjectAnypointmq `json:"anypointmq,omitempty"`
  Nats interface{} `json:"nats,omitempty"`
  Jms *MessageBindingsObjectJms `json:"jms,omitempty"`
  Sns interface{} `json:"sns,omitempty"`
  Sqs interface{} `json:"sqs,omitempty"`
  Stomp interface{} `json:"stomp,omitempty"`
  Redis interface{} `json:"redis,omitempty"`
  Ibmmq *MessageBindingsObjectIbmmq `json:"ibmmq,omitempty"`
  Solace interface{} `json:"solace,omitempty"`
  Googlepubsub *MessageBindingsObjectGooglepubsub `json:"googlepubsub,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}