package clients

type Client struct {
	Id            int32  `json:"id"`
	CompanyName   string `json:"companyName"`
	ContactPerson string `json:"contactPerson"`
	Email         string `json:"email"`
	Number        string `json:"number"`
}

func NewClient(id int32, companyName string, contactPerson string, email string, number string) (*Client, error) {
	return &Client{
		Id:            id,
		CompanyName:   companyName,
		ContactPerson: contactPerson,
		Email:         email,
		Number:        number}, nil
}

func CreateClient(companyName string, contactPerson string, email string, number string) (*Client, error) {
	return &Client{
		CompanyName:   companyName,
		ContactPerson: contactPerson,
		Email:         email,
		Number:        number}, nil
}
