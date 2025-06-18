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
Ablt to switch databases with ease