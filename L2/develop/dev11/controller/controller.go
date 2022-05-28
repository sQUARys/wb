package controller

import (
	"dev11/middleware"
	"dev11/structs"
	"log"
	"net/http"
	"regexp"
	"sync"
)

var (
	//POST
	createEvent = regexp.MustCompile(`^\/create_event[\/]?.+$`) // ^$ - полное совпадение должно быть
	updateEvent = regexp.MustCompile(`^\/update_event[\/]?.+$`) // \d+ - одна или более цифра доступная позже
	deleteEvent = regexp.MustCompile(`^\/delete_event[\/]?.+$`) // * - жабная функция
	//GET
	getEventForDay   = regexp.MustCompile(`^\/events_for_day[\/]$`)
	getEventForYear  = regexp.MustCompile(`^\/events_for_year[\/]$`)
	getEventForMonth = regexp.MustCompile(`^\/events_for_month[\/]$`)
	//Time
	dataFormat = "01/02/2006" // day / month/year
)

type Datastore struct {
	M map[string][]structs.EventInfo `json:"event-info-arr"`
	*sync.RWMutex
}

type UserHandler struct {
	Store *Datastore
}

func New() UserHandler {
	items := map[string][]structs.EventInfo{
		"10/10/2010": {{EventId: "1", EventName: "bob"}},
	}

	uh := UserHandler{
		Store: &Datastore{
			M:       items,
			RWMutex: &sync.RWMutex{},
		},
	}

	return uh
}

func (userH *UserHandler) ControllerHandler() {
	mux := http.NewServeMux()

	mux.Handle("/create_event/", userH)
	mux.Handle("/update_event/", userH)
	mux.Handle("/delete_event/", userH)

	mux.Handle("/events_for_day/", userH)
	mux.Handle("/events_for_year/", userH)
	mux.Handle("/events_for_month/", userH)

	wrappedMux := middleware.NewLogger(mux)

	log.Fatal(http.ListenAndServe("localhost:8080", wrappedMux))
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	switch {
	case r.Method == http.MethodPost && createEvent.MatchString(r.URL.Path):
		h.Create(w, r)
		return
	case r.Method == http.MethodPost && updateEvent.MatchString(r.URL.Path):
		h.Update(w, r)
		return
	case r.Method == http.MethodPost && deleteEvent.MatchString(r.URL.Path):
		h.Delete(w, r)
		return
	case r.Method == http.MethodGet && getEventForDay.MatchString(r.URL.Path):
		date, ok := ParseURL(r)
		if !ok {
			log.Print("Your body of request is empty")
			break
		}
		h.GetEventsForDay(w, date["date"][0])
		return
	case r.Method == http.MethodGet && getEventForMonth.MatchString(r.URL.Path):
		date, ok := ParseURL(r)
		if !ok {
			log.Print("Your body of request is empty")
			break
		}
		h.GetEventsForMonth(w, date["date"][0])
		return
	case r.Method == http.MethodGet && getEventForYear.MatchString(r.URL.Path):
		date, ok := ParseURL(r)
		if !ok {
			log.Print("Your body of request is empty")
			break
		}
		h.GetEventsForYear(w, date["date"][0])
		return
	default:
		NotFound(w)
		return
	}
}
