package clients

type Client struct {
	Id              int32  `json:"id"`
	CompanyName     string `json:"company_name"`
	ContactPerson   string `json:"contact_person"`
	Email           string `json:"email"`
	TelephoneNumber string `json:"telephone_number"`
}

func (c *Client) SetId(id int32) {
	c.Id = id
}

func NewClient(id int32, companyName, contactPerson, email, telephoneNumber string) (*Client, error) {
	return &Client{
		Id:              id,
		CompanyName:     companyName,
		ContactPerson:   contactPerson,
		Email:           email,
		TelephoneNumber: telephoneNumber}, nil
}

func CreateClient(companyName, contactPerson, email, telephoneNumber string) (*Client, error) {
	return &Client{
		CompanyName:     companyName,
		ContactPerson:   contactPerson,
		Email:           email,
		TelephoneNumber: telephoneNumber}, nil
}
