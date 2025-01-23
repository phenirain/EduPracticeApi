package clients

type Client struct {
	Id            int32  `json:"id"`
	CompanyName   string `json:"companyName"`
	ContactPerson string `json:"contactPerson"`
	Email         string `json:"email"`
	Number        string `json:"number"`
}
