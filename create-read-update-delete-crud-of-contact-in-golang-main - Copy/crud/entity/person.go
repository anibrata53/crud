package entity

//Paitent object for REST(CRUD)
type Paitent struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     int    `json:"phone"`
	Password  string `json:"password"`
	Acctype   string `json:"acctype"`
	Diseases  string `json:"diseases"`
	Medicine  string `json:"medicine"`
}

