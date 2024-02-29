package models

// AsyncApi_3Dot_0Dot_0SchemaDot represents a AsyncApi_3Dot_0Dot_0SchemaDot model.
type AsyncApi_3Dot_0Dot_0SchemaDot struct {
	Asyncapi             string
	Id                   string
	Info                 *Info
	Servers              map[string]interface{}
	DefaultContentType   string
	Channels             map[string]interface{}
	Operations           map[string]interface{}
	Components           *Components
	AdditionalProperties map[string]interface{}
}
