package magazine

type Subscriber struct {
	Name    string
	Rate    float64
	Active  bool
	Address //Annonymous struct field
}

type Employee struct {
	Name    string
	Salary  float64
	Address // Annonymous struct field
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
