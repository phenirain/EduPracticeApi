package clients

type ClientView struct {
	Query string
	View  ClientViewDb
}

type ClientViewDb struct {
	Id              int32  `db:"c_id"`
	CompanyName     string `db:"c_company_name"`
	ContactPerson   string `db:"c_contact_person"`
	Email           string `db:"c_email"`
	TelephoneNumber string `db:"c_telephone_number"`
}

func MustNewClientView() *ClientView {
	return &ClientView{
		Query: `SELECT id as c_id, company_name, contact_person, email, telephone_number FROM clients`,
		View:  ClientViewDb{},
	}
}