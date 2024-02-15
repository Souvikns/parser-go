
package models

// ServerScheme represents an enum of ServerScheme.
type ServerScheme uint

const (
  ServerSchemeKafka ServerScheme = iota
  ServerSchemeKafkaMinusSecure
  ServerSchemeAmqp
  ServerSchemeAmqps
  ServerSchemeMqtt
  ServerSchemeMqtts
  ServerSchemeSecureMinusMqtt
  ServerSchemeWs
  ServerSchemeWss
  ServerSchemeStomp
  ServerSchemeStomps
)

// Value returns the value of the enum.
func (op ServerScheme) Value() any {
	if op >= ServerScheme(len(ServerSchemeValues)) {
		return nil
	}
	return ServerSchemeValues[op]
}

var ServerSchemeValues = []any{"kafka","kafka-secure","amqp","amqps","mqtt","mqtts","secure-mqtt","ws","wss","stomp","stomps"}
var ValuesToServerScheme = map[any]ServerScheme{
  ServerSchemeValues[ServerSchemeKafka]: ServerSchemeKafka,
  ServerSchemeValues[ServerSchemeKafkaMinusSecure]: ServerSchemeKafkaMinusSecure,
  ServerSchemeValues[ServerSchemeAmqp]: ServerSchemeAmqp,
  ServerSchemeValues[ServerSchemeAmqps]: ServerSchemeAmqps,
  ServerSchemeValues[ServerSchemeMqtt]: ServerSchemeMqtt,
  ServerSchemeValues[ServerSchemeMqtts]: ServerSchemeMqtts,
  ServerSchemeValues[ServerSchemeSecureMinusMqtt]: ServerSchemeSecureMinusMqtt,
  ServerSchemeValues[ServerSchemeWs]: ServerSchemeWs,
  ServerSchemeValues[ServerSchemeWss]: ServerSchemeWss,
  ServerSchemeValues[ServerSchemeStomp]: ServerSchemeStomp,
  ServerSchemeValues[ServerSchemeStomps]: ServerSchemeStomps,
}
