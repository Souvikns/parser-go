package v3

import "encoding/json"

type Contact struct {
	name  string
	url   string
	email string
}

type ContactJSON struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Email string `json:"email"`
}

func (c *Contact) UnmarshalJSON(data []byte) error {
	var temp ContactJSON
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	c.parse(temp)
	return nil
}

func (c *Contact) parse(contact ContactJSON) {
	c.name = contact.Name
	c.url = contact.Url
	c.email = contact.Email
}

func (c Contact) Name() string {
	return c.name
}

func (c Contact) Url() string {
	return c.url
}

func (c Contact) Email() string {
	return c.url
}