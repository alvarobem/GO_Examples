package main

type person struct {
	name    string
	surName string
	age     int
}

type Person interface {
	talk() string
	walk() string
	getName() string
}

type PersonBuilder interface {
	Name(string) PersonBuilder
	SurName(string) PersonBuilder
	Age(int) PersonBuilder
	Build() Person
}

func NewPersonBuilder() PersonBuilder {
	return &personBuilder{}
}

type personBuilder struct {
	name    string
	surName string
	age     int
}

func (persBuilder *personBuilder) SurName(surName string) PersonBuilder {
	persBuilder.surName = surName
	return persBuilder
}

func (persBuilder *personBuilder) Name(name string) PersonBuilder {
	persBuilder.name = name
	return persBuilder
}

func (persBuilder *personBuilder) Age(age int) PersonBuilder {
	persBuilder.age = age
	return persBuilder
}

func (persBuilder *personBuilder) Build() Person {
	return &person{
		age:     persBuilder.age,
		name:    persBuilder.name,
		surName: persBuilder.surName,
	}
}

func (this person) walk() string {
	return this.name + " is walking"
}

func (this person) talk() string {
	return this.name + " is talking"
}

func (this person) getName() string {
	return this.name
}
