
package models

// MessageBindingsObject represents a MessageBindingsObject model.
type MessageBindingsObject struct {
  Http *MessageBindingsObjectHttp `json:"http"`
  Ws interface{} `json:"ws"`
  Amqp *MessageBindingsObjectAmqp `json:"amqp"`
  Amqp1 interface{} `json:"amqp1"`
  Mqtt *MessageBindingsObjectMqtt `json:"mqtt"`
  Kafka *MessageBindingsObjectKafka `json:"kafka"`
  Anypointmq *MessageBindingsObjectAnypointmq `json:"anypointmq"`
  Nats interface{} `json:"nats"`
  Jms *MessageBindingsObjectJms `json:"jms"`
  Sns interface{} `json:"sns"`
  Sqs interface{} `json:"sqs"`
  Stomp interface{} `json:"stomp"`
  Redis interface{} `json:"redis"`
  Ibmmq *MessageBindingsObjectIbmmq `json:"ibmmq"`
  Solace interface{} `json:"solace"`
  Googlepubsub *MessageBindingsObjectGooglepubsub `json:"googlepubsub"`
  AdditionalProperties map[string]interface{} `json:"-"`
}