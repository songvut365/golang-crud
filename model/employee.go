package model

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Salary    string `json:"salary"`
	Age       string `json:"age"`
}

type Employees struct {
	Employees []Employee `json:"employees"`
}
