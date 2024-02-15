
package models

// BindingsMinusPulsarMinus_0Dot_1Dot_0MinusChannelPersistence represents an enum of BindingsMinusPulsarMinus_0Dot_1Dot_0MinusChannelPersistence.
type BindingsMinusPulsarMinus_0Dot_1Dot_0MinusChannelPersistence uint

const (
  BindingsMinusPulsarMinus_0Dot_1Dot_0MinusChannelPersistencePersistent BindingsMinusPulsarMinus_0Dot_1Dot_0MinusChannelPersistence = iota
  BindingsMinusPulsarMinus_0Dot_1Dot_0MinusChannelPersistenceNonMinusPersistent
)

// Value returns the value of the enum.
func (op BindingsMinusPulsarMinus_0Dot_1Dot_0MinusChannelPersistence) Value() any {
	if op >= BindingsMinusPulsarMinus_0Dot_1Dot_0MinusChannelPersistence(len(BindingsMinusPulsarMinus_0Dot_1Dot_0MinusChannelPersistenceValues)) {
		return nil
	}
	return BindingsMinusPulsarMinus_0Dot_1Dot_0MinusChannelPersistenceValues[op]
}

var BindingsMinusPulsarMinus_0Dot_1Dot_0MinusChannelPersistenceValues = []any{"persistent","non-persistent"}
var ValuesToBindingsMinusPulsarMinus_0Dot_1Dot_0MinusChannelPersistence = map[any]BindingsMinusPulsarMinus_0Dot_1Dot_0MinusChannelPersistence{
  BindingsMinusPulsarMinus_0Dot_1Dot_0MinusChannelPersistenceValues[BindingsMinusPulsarMinus_0Dot_1Dot_0MinusChannelPersistencePersistent]: BindingsMinusPulsarMinus_0Dot_1Dot_0MinusChannelPersistencePersistent,
  BindingsMinusPulsarMinus_0Dot_1Dot_0MinusChannelPersistenceValues[BindingsMinusPulsarMinus_0Dot_1Dot_0MinusChannelPersistenceNonMinusPersistent]: BindingsMinusPulsarMinus_0Dot_1Dot_0MinusChannelPersistenceNonMinusPersistent,
}