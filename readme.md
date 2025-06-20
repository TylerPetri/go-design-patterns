## FACTORY

Creating `type Product struct`\
Creating `func (p *Product) New(){}` that creates a product rather than manually creating a product in main.go\
This is the Factory design, when you see a `func New()`

## ABSTRACT FACTORY

Another creational design pattern but provides an interface for creating families of related objects\
Ie: building Dog might be called from DB, and building Cat might be called from API\
Client side doesn't need to know, abstract factory design takes care of it all

## REPOSITORY

An intermediary layer between an application's business logic and data storage\
Ablt to switch databases with ease\
Able to test without connecting to database

## SINGLETON

Allows us to restruct the instantiation of something to a singular instance.\
Useful when exaxctly one object is needed to coordinate actions across a system.\
For information that will never change.\
Like when new devs want to work on something related to DB and make calls, but 30 use at same time and exhausts\
Use `connection.New()` in ./configuration/config.go, won't create a new instance (prone to exhausting DB), it will use the current that was developped.

## BUILDER  

Allows to chain methods, ie:\
`SetSpecies("dog") SetBreed("Germain Shepherd Dog") Set... Build()`

## BUILER FLUENT INTERFACE

Nicer and more readable than sending a big body of params to just one function, you can see which methods and only send to specific methods (not need all params filled)\
myAddress := CreateAddress().SetStreet("Main St.").SetNumber(11).SetCity("Gotham").SetCountry("USA")\
func CreateAddress() *Address { return &Address{} }\
func (a *Address) SetStreet(streetName string) *Address { a.street = streetName; return a }