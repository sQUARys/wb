package main

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

type Visitor interface {
	VisitMuseum(m *Museum) string
	VisitCinema(c *Cinema) string
}

type Place interface {
	Accept(v Visitor) string
}
type Persons struct {
}

func (p *Persons) VisitMuseum(m *Museum) string {
	return m.BuyMuseumTickets()
}

func (p *Persons) VisitCinema(c *Cinema) string {
	return c.BuyCinemaTickets()
}

type City struct {
	places []Place
}

func (c *City) Add(p Place) {
	c.places = append(c.places, p)
}

func (c *City) Accept(v Visitor) string {
	var res string
	for _, p := range c.places {
		res += p.Accept(v)
	}
	return res
}

type Museum struct{}

func (m *Museum) Accept(v Visitor) string {
	return v.VisitMuseum(m)
}

func (m *Museum) BuyMuseumTickets() string {
	return "Museum tiskets buying..."
}

type Cinema struct{}

func (c *Cinema) Accept(v Visitor) string {
	return v.VisitCinema(c)
}

func (c *Cinema) BuyCinemaTickets() string {
	return "Cinema tiskets buying..."
}

func main() {
	city := new(City)

	city.Add(&Museum{})
	city.Add(&Cinema{})

	res := city.Accept(&Persons{})
	fmt.Println(res)
}
