package controller

import (
	"dev11/structs"
	"dev11/validator"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func ParseURL(r *http.Request) (url.Values, bool) {
	ok := true

	r.ParseForm()
	queryMap := r.Form
	if len(queryMap) == 0 {
		ok = false
	}
	return queryMap, ok
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var u structs.Input

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		//internalServerError(w, r)
		return
	}

	if !isValidJSON(w, u) {
		return
	}

	valid := validator.IsValid("del", u)
	if valid != nil {
		log.Fatal(valid)
		internalServerError(w)
		return
	}

	date := ParseDateToString(w, u.Date)

	h.Store.Lock()
	delete(h.Store.M, date)
	h.Store.Unlock()

	fmt.Println("DELETE , ", h.Store.M)
	serializeJson(w, h.Store.M)

}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var u structs.Input
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		//internalServerError(w, r)
		return
	}

	if !isValidJSON(w, u) {
		return
	}

	valid := validator.IsValid("upd", u)
	if valid != nil {
		log.Fatal(valid)
		internalServerError(w)
		return
	}

	date := ParseDateToString(w, u.Date)
	newEvent := structs.EventInfo{
		EventId:   u.ID,
		EventName: u.Name,
	}

	h.Store.Lock()
	h.Store.M[date] = append(h.Store.M[date], newEvent)
	h.Store.Unlock()

	fmt.Println("Update", h.Store.M)
	serializeJson(w, h.Store.M)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var u structs.Input

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		internalServerError(w)
		return
	}

	if !isValidJSON(w, u) {
		return
	}

	valid := validator.IsValid("crt", u)
	if valid != nil {
		log.Fatal(valid)
		internalServerError(w)
		return
	}

	date := ParseDateToString(w, u.Date)

	newEvent := structs.EventInfo{
		EventId:   u.ID,
		EventName: u.Name,
	}

	h.Store.Lock()
	h.Store.M[date] = append(h.Store.M[date], newEvent)
	h.Store.Unlock()

	fmt.Println("Create", h.Store.M)
	serializeJson(w, h.Store.M)
}

func ParseDate(w http.ResponseWriter, date string) structs.Date {
	currDate, errorDate := time.Parse(dataFormat, date)
	if errorDate != nil {
		w.WriteHeader(400)
		return structs.Date{}
	}

	year, day, month := currDate.Date()
	dateStruct := structs.Date{
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

func (h *UserHandler) GetEventsForDay(w http.ResponseWriter, date string) {
	parsedDate := ParseDateToString(w, date)
	for key, _ := range h.Store.M {
		if key == parsedDate {
			serializeJson(w, h.Store.M[key])
			break
		}
	}
}

func (h *UserHandler) GetEventsForMonth(w http.ResponseWriter, date string) {
	var memory [][]structs.EventInfo
	parsedDate := ParseDate(w, date)
	for key, _ := range h.Store.M {
		parsedKey := ParseDate(w, key)
		if parsedKey.Month == parsedDate.Month && parsedDate.Year == parsedKey.Year {
			memory = append(memory, h.Store.M[key])
		}
	}
	serializeJson(w, memory)
}

func (h *UserHandler) GetEventsForYear(w http.ResponseWriter, date string) {

	var memory [][]structs.EventInfo
	parsedDate := ParseDate(w, date)

	for key, _ := range h.Store.M {
		parsedKey := ParseDate(w, key)
		if parsedDate.Year == parsedKey.Year {
			memory = append(memory, h.Store.M[key])
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

func NotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}

func internalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func isValidJSON(w http.ResponseWriter, js structs.Input) bool {
	val, errInt := strconv.Atoi(js.ID)

	if errInt != nil || val <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Id field is invalid int."))
		return false
	}

	return true
}
