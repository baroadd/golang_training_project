package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	. "reserve-service/config"
	. "reserve-service/dao"
	. "reserve-service/models"
	"gopkg.in/mgo.v2/bson"
)

var config = Config{}
var dao = EventsDAO{}

func AllEvent (w http.ResponseWriter, r *http.Request) {
	events, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, events)
}

func CreateEvent (w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	var event Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	event.ID = bson.NewObjectId()
	if err := dao.Insert(event); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, event)
}

func FindEventEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	event, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Event ID")
		return
	}
	respondWithJSON(w, http.StatusOK, event)
}

func UpdateEventEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var event Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(event); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteEventEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var event Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(event); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func startServer() error {
	r := mux.NewRouter()
	r.HandleFunc("/events", AllEvent).Methods("GET")
	r.HandleFunc("/events/{id}", FindEventEndPoint).Methods("GET")
	r.HandleFunc("/events", CreateEvent).Methods("POST")
	r.HandleFunc("/events", UpdateEventEndPoint).Methods("PUT")
	r.HandleFunc("/events", DeleteEventEndPoint).Methods("DELETE")
	
	err := http.ListenAndServe(":3000", r)
	return err

}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}