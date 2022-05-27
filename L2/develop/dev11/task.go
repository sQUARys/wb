package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"sync"
	"time"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через querystring, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

//http://localhost:8080/create_event/?user_id=3&date=2019-09-09

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

type Input struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Date string `json:"date"`
}

type Date struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

type EventInfo struct {
	EventId   string `json:"event-id"`
	EventName string `json:"event-name"`
}

type datastore struct {
	m map[Date][]EventInfo
	*sync.RWMutex
}

type userHandler struct {
	store *datastore
}

func main() {
	mux := http.NewServeMux()
	userH := &userHandler{
		store: &datastore{
			m: map[Date][]EventInfo{
				{Day: 10, Month: 10, Year: 2010}: {{EventId: "1", EventName: "bob"}},
			},
			RWMutex: &sync.RWMutex{},
		},
	}

	mux.Handle("/create_event/", userH)
	mux.Handle("/update_event/", userH)
	mux.Handle("/delete_event/", userH)

	mux.Handle("/events_for_day/", userH)
	mux.Handle("/events_for_year/", userH)
	mux.Handle("/events_for_month/", userH)

	http.ListenAndServe("localhost:8080", mux)
}

func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var queryMap url.Values

	if r.Method == http.MethodGet {
		r.ParseForm()
		queryMap = r.Form
	}

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
		h.GetEventsForDay(w, queryMap["date"][0])
		return
	case r.Method == http.MethodGet && getEventForMonth.MatchString(r.URL.Path):
		h.GetEventsForMonth(w, queryMap["date"][0])
		return
	case r.Method == http.MethodGet && getEventForYear.MatchString(r.URL.Path):
		h.GetEventsForYear(w, queryMap["date"][0])
		return
	default:
		//notFound(w, r)
		return
	}
}

func (h *userHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var u Input
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		//internalServerError(w, r)
		return
	}

	date := ParseDate(w, u.Date)

	h.store.Lock()
	delete(h.store.m, date)
	h.store.Unlock()

	fmt.Println("DELETE , ", h.store.m)
	w.WriteHeader(http.StatusOK)

}

func (h *userHandler) Update(w http.ResponseWriter, r *http.Request) {
	var u Input
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		//internalServerError(w, r)
		return
	}
	jsonBytes, err := json.Marshal(u)

	if err != nil {
		//internalServerError(w, r)
		return
	}

	date := ParseDate(w, u.Date)
	newEvent := EventInfo{
		EventId:   u.ID,
		EventName: u.Name,
	}

	h.store.Lock()
	h.store.m[date] = append(h.store.m[date], newEvent)

	//h.store.m[date] = newEvent
	h.store.Unlock()

	fmt.Println("Update", h.store.m)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	var u Input
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		//internalServerError(w, r)
		return
	}
	jsonBytes, err := json.Marshal(u)

	if err != nil {
		//internalServerError(w, r)
		return
	}

	date := ParseDate(w, u.Date)

	newEvent := EventInfo{
		EventId:   u.ID,
		EventName: u.Name,
	}

	h.store.Lock()
	h.store.m[date] = append(h.store.m[date], newEvent)
	h.store.Unlock()

	fmt.Println("Create", h.store.m)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func ParseDate(w http.ResponseWriter, date string) Date {
	currDate, errorDate := time.Parse(dataFormat, date)
	if errorDate != nil {
		w.WriteHeader(400)
		return Date{}
	}

	year, day, month := currDate.Date()
	dateStruct := Date{
		Day:   int(day),
		Month: month,
		Year:  year,
	}
	return dateStruct
}

func (h *userHandler) GetEventsForDay(w http.ResponseWriter, date string) {
	var jsonBytes []byte
	var err error
	parsedDate := ParseDate(w, date)
	fmt.Println(parsedDate)
	for key, _ := range h.store.m {
		if key == parsedDate {
			jsonBytes, err = json.Marshal(h.store.m[key])
			if err != nil {
				//internalServerError(w, r)
				return
			}
			break
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *userHandler) GetEventsForMonth(w http.ResponseWriter, date string) {
	var jsonBytes []byte
	var err error
	var memory [][]EventInfo
	parsedDate := ParseDate(w, date)
	for key, _ := range h.store.m {
		if key.Month == parsedDate.Month && parsedDate.Year == key.Year {
			memory = append(memory, h.store.m[key])
		}
	}
	jsonBytes, err = json.Marshal(memory)
	if err != nil {
		//internalServerError(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *userHandler) GetEventsForYear(w http.ResponseWriter, date string) {
	var jsonBytes []byte
	var err error
	var memory [][]EventInfo
	parsedDate := ParseDate(w, date)

	for key, _ := range h.store.m {
		if parsedDate.Year == key.Year {
			memory = append(memory, h.store.m[key])
		}
	}

	jsonBytes, err = json.Marshal(memory)
	if err != nil {
		//internalServerError(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
