package models

func (Person) TableName() string {
	return "Person"
}

type Person struct {
	Id   int
	Name string
	Age  int
	Sex  int
}
