package structs

type Input struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Date string `json:"date"`
}

type Date struct {
	Day   int
	Month int
	Year  int
}

type EventInfo struct {
	EventId    string `json:"event-id"`
	EventName  string `json:"event-name"`
	EventDescr string `json:"event-name, omitempty"`
}
