package v3

type Info struct {
	title          string
	version        string
	description    *string
	termsOfService *string
	contact        *Contact
	// license        *string
	// tags           *string
}

type InfoJSON struct {
	Title          string      `json:"title"`
	Version        string      `json:"version"`
	Description    *string     `json:"description"`
	TermsOfService *string     `json:"termsOfService"`
	Contact        ContactJSON `json:"contact"`
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
