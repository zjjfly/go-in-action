package entities

type User struct {
	Name  string
	Email string
	//这个是未公开的字段,因为它以小写字母开头
	age int
}

type person struct {
	Name  string
	Email string
}

type Admin struct {
	person
	Rights int
}
