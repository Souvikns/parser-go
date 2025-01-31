package v3

import "encoding/json"

type Info struct {
	title          string
	version        string
	description    string
	termsOfService string
	contact        Contact
	// license        *string
	// tags           *string
}

type InfoJSON struct {
	Title          string      `json:"title"`
	Version        string      `json:"version"`
	Description    string      `json:"description,omitempty"`
	TermsOfService string      `json:"termsOfService,omitempty"`
	Contact        ContactJSON `json:"contact,omitempty"`
}

func (i *Info) UnmarshalJSON(data []byte) error {
	var temp InfoJSON
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	i.parse(temp)
	return nil
}

func (i *Info) parse(info InfoJSON) {
	i.title = info.Title
	i.version = info.Version
	i.description = info.Description
	i.termsOfService = info.TermsOfService
	i.contact.parse(info.Contact)
}

func (i Info) Title() string {
	return i.title
}

func (i Info) Version() string {
	return i.version
}

func (i Info) Contact() Contact {
	return i.contact
}