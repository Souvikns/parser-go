
package models
import (  
  "encoding/json"
)
// BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocol represents an enum of BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocol.
type BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocol uint

const (
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolHttp BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocol = iota
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolHttps
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolEmail
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolEmailMinusJson
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolSms
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolSqs
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolApplication
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolLambda
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolFirehose
)

// Value returns the value of the enum.
func (op BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocol) Value() any {
	if op >= BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocol(len(BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolValues)) {
		return nil
	}
	return BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolValues[op]
}

var BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolValues = []any{"http","https","email","email-json","sms","sqs","application","lambda","firehose"}
var ValuesToBindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocol = map[any]BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocol{
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolValues[BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolHttp]: BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolHttp,
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolValues[BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolHttps]: BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolHttps,
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolValues[BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolEmail]: BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolEmail,
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolValues[BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolEmailMinusJson]: BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolEmailMinusJson,
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolValues[BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolSms]: BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolSms,
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolValues[BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolSqs]: BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolSqs,
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolValues[BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolApplication]: BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolApplication,
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolValues[BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolLambda]: BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolLambda,
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolValues[BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolFirehose]: BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocolFirehose,
}

 
func (op *BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocol) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToBindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocol[v]
  return nil
}

func (op BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocol) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}