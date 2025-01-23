package deliveries

type Driver struct {
	Id          int32  `json:"id"`
	Surname     string `json:"surname"`
	MiddleName  string `json:"middleName"`
	FirstName   string `json:"firstName"`
	NumberPhone string `json:"numberPhone"`
}

func NewDriver(id int32, surname string, middleName string, firstName string, numberPhone string) Driver {
	return Driver{
		Id:          id,
		Surname:     surname,
		MiddleName:  middleName,
		FirstName:   firstName,
		NumberPhone: numberPhone,
	}
}

func CreateDriver(surname string, middleName string, firstName string, numberPhone string) Driver {
	return Driver{
		Surname:     surname,
		MiddleName:  middleName,
		FirstName:   firstName,
		NumberPhone: numberPhone,
	}
}
