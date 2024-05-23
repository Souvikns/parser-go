
package models

// ServerBindingsObjectPulsar represents a ServerBindingsObjectPulsar model.
type ServerBindingsObjectPulsar struct {
  BindingVersion *ServerBindingsObjectPulsarBindingVersion `json:"bindingVersion,omitempty"`
  Tenant string `json:"tenant,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}