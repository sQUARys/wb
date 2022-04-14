package model

type Request struct {
	Id       int
	Delivery DeliveryJSON
	Thing    ThingJSON
}

type DeliveryJSON struct {
	Name    string
	Phone   string
	City    string
	Address string
}

type ThingJSON struct {
	Price    int
	ItemName string
	Brand    string
}
