package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON. DONE
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области. DONE
	4. Реализовать middleware для логирования запросов DONE
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через querystring, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400.
В случае остальных ошибок сервер должен возвращать HTTP 500.
Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

//http://localhost:8080/create_event/?user_id=3&date=2019-09-09

// Сделать валидацию параметров методов  /create_event и /update_event.
// Разобраться что такое бизнес логика

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
	Day   int
	Month int
	Year  int
}

type EventInfo struct {
	EventId    string `json:"event-id"`
	EventName  string `json:"event-name"`
	EventDescr string `json:"event-name, omitempty"`
}

type datastore struct {
	m map[string][]EventInfo `json:"event-info-arr"`
	*sync.RWMutex
}

type userHandler struct {
	store *datastore
}

func main() {
	mux := http.NewServeMux()
	userH := &userHandler{
		store: &datastore{
			m: map[string][]EventInfo{
				"10/10/2010": {{EventId: "1", EventName: "bob"}},
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

	wrappedMux := NewLogger(mux)

	log.Fatal(http.ListenAndServe("localhost:8080", wrappedMux))

}

//Logging middleware

type Logger struct {
	handler http.Handler
}

//ServeHTTP handles the request by passing it to the real
//handler and logging the request details
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}

//NewLogger constructs a new Logger middleware handler
func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}

//Validation

type Validator struct {
	err error
}

func (i *Input) isValid(option string) error {
	validator := Validator{}

	switch option {
	case "del":
		validator.MustHaveDate(*i)
	case "crt":
		validator.MustHaveId(*i)
		validator.MustHaveName(*i)
	case "upd":
		validator.MustHaveNameOrId(*i)
		validator.MustHaveDate(*i)
	}
	return validator.IsValid()
}

func (v *Validator) IsValid() error {
	return v.err
}

func (v *Validator) MustHaveId(input Input) bool {
	if v.err != nil {
		return false
	}
	if input.ID == "" {
		v.err = fmt.Errorf("Must have an id field")
		return false
	}
	return true
}

func (v *Validator) MustHaveName(input Input) bool {
	if v.err != nil {
		return false
	}
	if input.Name == "" {
		v.err = fmt.Errorf("Must have a name field")
		return false
	}
	return true
}

func (v *Validator) MustHaveNameOrId(input Input) bool {
	if v.err != nil {
		return false
	}
	if input.Name == "" && input.ID == "" {
		v.err = fmt.Errorf("Must have a name or id fields")
		return false
	}
	return true
}

func (v *Validator) MustHaveDate(input Input) bool {
	if v.err != nil {
		return false
	}
	if input.Date == "" {
		v.err = fmt.Errorf("Must have a date field")
		return false
	}
	return true
}

//Server methods

func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
		//notFound(w, r)
		return
	}
}

func ParseURL(r *http.Request) (url.Values, bool) {
	log.Printf("Handle : ", r.URL.Path)
	ok := true

	r.ParseForm()
	queryMap := r.Form
	if len(queryMap) == 0 {
		ok = false
	}
	return queryMap, ok
}

func (h *userHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var u Input
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		//internalServerError(w, r)
		return
	}

	valid := u.isValid("del")
	if valid != nil {
		log.Fatal(valid)
		return
	}

	date := ParseDateToString(w, u.Date)

	h.store.Lock()
	delete(h.store.m, date)
	h.store.Unlock()

	fmt.Println("DELETE , ", h.store.m)
	serializeJson(w, h.store.m)

}

func (h *userHandler) Update(w http.ResponseWriter, r *http.Request) {
	var u Input
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		//internalServerError(w, r)
		return
	}

	valid := u.isValid("upd")
	if valid != nil {
		log.Fatal(valid)
		return
	}

	date := ParseDateToString(w, u.Date)
	newEvent := EventInfo{
		EventId:   u.ID,
		EventName: u.Name,
	}

	h.store.Lock()
	h.store.m[date] = append(h.store.m[date], newEvent)
	h.store.Unlock()

	fmt.Println("Update", h.store.m)
	serializeJson(w, h.store.m)
}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	var u Input
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		//internalServerError(w, r)
		return
	}

	valid := u.isValid("crt")
	if valid != nil {
		log.Fatal(valid)
		return
	}

	date := ParseDateToString(w, u.Date)

	newEvent := EventInfo{
		EventId:   u.ID,
		EventName: u.Name,
	}

	h.store.Lock()
	h.store.m[date] = append(h.store.m[date], newEvent)
	h.store.Unlock()

	fmt.Println("Create", h.store.m)
	serializeJson(w, h.store.m)
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

func ParseDateToString(w http.ResponseWriter, date string) string {
	currDate, errorDate := time.Parse(dataFormat, date)
	if errorDate != nil {
		w.WriteHeader(400)
		return ""
	}

	year, day, month := currDate.Date()
	dayStr := strconv.Itoa(int(day))
	monthStr := strconv.Itoa(month)
	yearStr := strconv.Itoa(year)
	if len(dayStr) == 1 {
		dayStr = "0" + dayStr
	}
	if len(monthStr) == 1 {
		monthStr = "0" + monthStr
	}
	str := []string{dayStr, monthStr, yearStr}
	return strings.Join(str, "/")
}

func (h *userHandler) GetEventsForDay(w http.ResponseWriter, date string) {
	parsedDate := ParseDateToString(w, date)
	for key, _ := range h.store.m {
		if key == parsedDate {
			serializeJson(w, h.store.m[key])
			break
		}
	}
}

func (h *userHandler) GetEventsForMonth(w http.ResponseWriter, date string) {
	var memory [][]EventInfo
	parsedDate := ParseDate(w, date)
	for key, _ := range h.store.m {
		parsedKey := ParseDate(w, key)
		if parsedKey.Month == parsedDate.Month && parsedDate.Year == parsedKey.Year {
			memory = append(memory, h.store.m[key])
		}
	}
	serializeJson(w, memory)
}

func (h *userHandler) GetEventsForYear(w http.ResponseWriter, date string) {

	var memory [][]EventInfo
	parsedDate := ParseDate(w, date)

	for key, _ := range h.store.m {
		parsedKey := ParseDate(w, key)
		if parsedDate.Year == parsedKey.Year {
			memory = append(memory, h.store.m[key])
		}
	}

	serializeJson(w, memory)
}

func serializeJson(w http.ResponseWriter, input interface{}) {
	js, err := json.Marshal(input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
