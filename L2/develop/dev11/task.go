package main

import (
	"fmt"
	"net/http"
	"regexp"
	"sync"
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

// renderJSON преобразует 'v' в формат JSON и записывает результат, в виде ответа, в w.
//func renderJSON(w http.ResponseWriter, v interface{}) {
//	js, err := json.Marshal(v)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	w.Header().Set("Content-Type", "application/json")
//	w.Write(js)
//}
//
//func main() {
//	mux := http.NewServeMux()
//
//	mux.HandleFunc("/task/", taskHandler())
//}
//package main

var (
	//POST
	createEvent = regexp.MustCompile(`^\/create_event[\/]?.+$`) // ^$ - полное совпадение должно быть
	updateEvent = regexp.MustCompile(`^\/update_event[\/]?.+$`) // \d+ - одна или более цифра доступная позже
	deleteEvent = regexp.MustCompile(`^\/delete_event[\/]?.+$`) // * - жабная функция
	//GET
	getEventForDay   = regexp.MustCompile(`^\/events_for_day[\/]$`)
	getEventForWeek  = regexp.MustCompile(`^\/events_for_week[\/]$`)
	getEventForMonth = regexp.MustCompile(`^\/events_for_month[\/]$`)
)

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type datastore struct {
	m map[string]user
	*sync.RWMutex
}

type userHandler struct {
	store *datastore
}

func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	////POST
	//createEvent = regexp.MustCompile(`^\/create_event[\/]?.+$`) // ^$ - полное совпадение должно быть
	//updateEvent = regexp.MustCompile(`^\/update_event[\/]?.+$`) // \d+ - одна или более цифра доступная позже
	//deleteEvent = regexp.MustCompile(`^\/delete_event[\/]?.+$`) // * - жабная функция
	////GET
	//getEventForDay   = regexp.MustCompile(`^\/events_for_day[\/]$`)
	//getEventForWeek  = regexp.MustCompile(`^\/events_for_week[\/]$`)
	//getEventForMonth = regexp.MustCompile(`^\/events_for_month[\/]$`)

	w.Header().Set("content-type", "application/json")
	r.ParseForm()
	queryMap := r.Form
	fmt.Println(queryMap, r.Method, createEvent.MatchString(r.URL.Path))
	switch {
	case r.Method == http.MethodGet && getEventForDay.MatchString(r.URL.Path):
		fmt.Println("Get1", queryMap)
		//h.List(w, r)
		return
	case r.Method == http.MethodGet && getEventForWeek.MatchString(r.URL.Path):
		fmt.Println("Get2", queryMap)
		//h.Get(w, r)
		return
	case r.Method == http.MethodGet && getEventForMonth.MatchString(r.URL.Path):
		fmt.Println("Get3", queryMap)
		//h.List(w, r)
		return
	case r.Method == http.MethodPost && createEvent.MatchString(r.URL.Path):
		fmt.Println("Post1", queryMap)
		//h.Create(w, r)
		return
	case r.Method == http.MethodPost && updateEvent.MatchString(r.URL.Path):
		fmt.Println("Post2", queryMap)
		//h.Create(w, r)
		return
	case r.Method == http.MethodPost && deleteEvent.MatchString(r.URL.Path):
		fmt.Println("Post3", queryMap)
		//h.Create(w, r)
		return
	default:
		//notFound(w, r)
		return
	}
}

//func (h *userHandler) List(w http.ResponseWriter, r *http.Request) {
//	h.store.RLock()
//	users := make([]user, 0, len(h.store.m))
//	for _, v := range h.store.m {
//		users = append(users, v)
//	}
//	h.store.RUnlock()
//	jsonBytes, err := json.Marshal(users)
//	if err != nil {
//		internalServerError(w, r)
//		return
//	}
//	w.WriteHeader(http.StatusOK)
//	w.Write(jsonBytes)
//}
//
//func (h *userHandler) Get(w http.ResponseWriter, r *http.Request) {
//	matches := getUserRe.FindStringSubmatch(r.URL.Path)
//	if len(matches) < 2 {
//		notFound(w, r)
//		return
//	}
//	h.store.RLock()
//	u, ok := h.store.m[matches[1]]
//	h.store.RUnlock()
//	if !ok {
//		w.WriteHeader(http.StatusNotFound)
//		w.Write([]byte("user not found"))
//		return
//	}
//	jsonBytes, err := json.Marshal(u)
//	if err != nil {
//		internalServerError(w, r)
//		return
//	}
//	w.WriteHeader(http.StatusOK)
//	w.Write(jsonBytes)
//}
//
//func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {
//	var u user
//	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
//		internalServerError(w, r)
//		return
//	}
//	h.store.Lock()
//	h.store.m[u.ID] = u
//	h.store.Unlock()
//	jsonBytes, err := json.Marshal(u)
//	if err != nil {
//		internalServerError(w, r)
//		return
//	}
//	w.WriteHeader(http.StatusOK)
//	w.Write(jsonBytes)
//}
//
//func internalServerError(w http.ResponseWriter, r *http.Request) {
//	w.WriteHeader(http.StatusInternalServerError)
//	w.Write([]byte("internal server error"))
//}
//
//func notFound(w http.ResponseWriter, r *http.Request) {
//	w.WriteHeader(http.StatusNotFound)
//	w.Write([]byte("not found"))
//}

func main() {
	mux := http.NewServeMux()
	userH := &userHandler{
		store: &datastore{
			m: map[string]user{
				"1": user{ID: "1", Name: "bob"},
			},
			RWMutex: &sync.RWMutex{},
		},
	}

	//POST
	//createEvent = regexp.MustCompile(`^\/create_event[\/]*$`) // ^$ - полное совпадение должно быть
	//updateEvent = regexp.MustCompile(`^\/update_event[\/]*$`) // \d+ - одна или более цифра доступная позже
	//deleteEvent = regexp.MustCompile(`^\/delete_event[\/]*$`) // * - жабная функция
	////GET
	//getEventForDay   = regexp.MustCompile(`^\/events_for_day[\/]*$`)
	//getEventForWeek  = regexp.MustCompile(`^\/events_for_week[\/]*$`)
	//getEventForMonth = regexp.MustCompile(`^\/events_for_month[\/]*$`)

	mux.Handle("/create_event/", userH)
	mux.Handle("/update_event/", userH)

	http.ListenAndServe("localhost:8080", mux)
}
