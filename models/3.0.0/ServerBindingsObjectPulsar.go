
package models

// ServerBindingsObjectPulsar represents a ServerBindingsObjectPulsar model.
type ServerBindingsObjectPulsar struct {
  BindingVersion *ServerBindingsObjectPulsarBindingVersion `json:"bindingVersion"`
  Tenant string `json:"tenant"`
  AdditionalProperties map[string]interface{} `json:"-"`
}