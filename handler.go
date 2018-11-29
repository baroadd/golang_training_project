package main

import (
	"encoding/json"
	"log"
	"net/http"
	. "reserve-service/config"
	. "reserve-service/dao"
	. "reserve-service/models"
	"os"
	"encoding/csv"
	"strconv"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var config = Config{}
var dao = EventsDAO{}

func AllEvent(w http.ResponseWriter, r *http.Request) {
	log.Println("NEW REQ=>", r.Method, r.URL)
	setupResponse(&w, r)
	events, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, events)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	log.Println("NEW REQ=>", r.Method, r.URL, r.Body)
	setupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}
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
	setupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}
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
	setupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}
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

func ExportCSV(w http.ResponseWriter, r *http.Request){
	setupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}
	defer r.Body.Close()
	var event Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	file, err := os.Create("result.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for index, value := range event.User {
		err := writer.Write([]string{strconv.Itoa(index+1),value})
		if err != nil {
			log.Fatal(err)
		}
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
	r.HandleFunc("/events", CreateEvent).Methods("POST", "OPTIONS")
	r.HandleFunc("/events", UpdateEventEndPoint).Methods("PUT")
	r.HandleFunc("/events/delete", DeleteEventEndPoint).Methods("POST", "OPTIONS")
	r.HandleFunc("/event/report", ExportCSV).Methods("POST","OPTIONS")

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

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(*w).Header().Set("Access-Control-Max-Age", "86400")
}
