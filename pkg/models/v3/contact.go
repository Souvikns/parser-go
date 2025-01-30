package v3

type Contact struct {
	name  *string
	url   *string
	email *string
}

type ContactJSON struct {
	Name  *string `json:"name"`
	Url   *string `json:"url"`
	Email *string `json:"email"`
}

func (c *Contact) parse(contact ContactJSON) {
	c.name = contact.Name
	c.url = contact.Url
	c.email = contact.Email
}

func (c Contact) Name() *string {
	return c.name
}

func (c Contact) Url() *string {
	return c.url
}

func (c Contact) Email() *string {
	return c.url
}