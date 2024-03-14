package handler

import (
	"dev11/service"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func NewServer(store *service.Calendar) *Server {
	return &Server{store: store}
}
//обработка запроса создания ивента
func (ss *Server) HandlerCreateEvent(w http.ResponseWriter, r *http.Request) {
	_, date, mes, err := handlerDataPost(r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	id := ss.store.CreateEvent(date, mes)
	writeResponse(w, id)
}

//обработка запроса обнавления ивента
func (ss *Server) HandlerUpdateEvent(w http.ResponseWriter, r *http.Request) {
	id, date, mes, err := handlerDataPost(r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	event, err := ss.store.UpdateEvent(id, date, mes)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}

	writeResponse(w, event)
}

//обработка запроса удаления ивента
func (ss *Server) HandlerDeleteEvent(w http.ResponseWriter, r *http.Request) {
	id, _, _, err := handlerDataPost(r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	errDelete := ss.store.DeleteEvent(id)
	if errDelete != nil {
		http.Error(w, errDelete.Error(), 503)
		return
	}

	writeResponse(w, "Element deleted")
}

//обработка запроса ивента на день
func (ss *Server) HandlerEventsForDay(w http.ResponseWriter, r *http.Request) {

	date := handlerDataGet(r)
	fmt.Println(date)
	events, err := ss.store.EventsForDate(date, 1)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	if len(events) == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	writeResponse(w, events)
}

//обработка запроса ивента на неделю
func (ss *Server) HandlerEventsForWeek(w http.ResponseWriter, r *http.Request) {
	date := handlerDataGet(r)

	events, err := ss.store.EventsForDate(date, 7)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	if len(events) == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	writeResponse(w, events)
}
//обработка запроса ивента на месяц
func (ss *Server) HandlerEventsForMonth(w http.ResponseWriter, r *http.Request) {
	date := handlerDataGet(r)

	events, err := ss.store.EventsForDate(date, 30)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	if len(events) == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	writeResponse(w, events)
}

//обработчик запроса POST
func handlerDataPost(r *http.Request) (int, time.Time, string, error) {
	var id int
	var date time.Time
	var mes string

	idString := r.FormValue("id")
	if idString != "" {
		idInt, err := strconv.Atoi(idString)
		if err != nil {
			return 0, time.Time{}, "", errors.New("400: invalid int")
		}
		id = idInt
	}

	dateString := r.FormValue("date")
	if dateString != "" {
		dateString += "T00:00:00Z"
		dateTime, err := time.Parse(time.RFC3339, dateString)
		if err != nil {
			return 0, time.Time{}, "", errors.New("400: invalid date")
		}
		date = dateTime
	}

	mes = r.FormValue("mes")

	return id, date, mes, nil
}

// обработчик запроса GET
func handlerDataGet(r *http.Request) time.Time {
	dateF := r.FormValue("date") + "T00:00:00Z"
	date, err := time.Parse(time.RFC3339, dateF)
	if err != nil {
		fmt.Println(err)
	}
	return date
}

//Вспомогательная функция сериализация объектов JSON
func writeResponse(w http.ResponseWriter, v interface{}) {
	resultJSON := struct {
		Result interface{}
	}{Result: v}

	js, err := json.Marshal(&resultJSON)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
