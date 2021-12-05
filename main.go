package main

import (
	"log"
	"net/http"
	"strings"

	store "go-api/store"
)

var db = store.New()

type GetHandler struct{}

func (h *GetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		KeyId := strings.Split(r.URL.Path, "/")[2]
		Data := db.Get(KeyId)
		log.Println("KeyId : " + KeyId + " " + "Value : " + Data)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(Data))
	}
}

type AddHandler struct{}

func (h *AddHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		//val variable returns with the value as posted.
		val := r.FormValue("val")
		//fmt.Println(val)
		incrementId := db.Add(val)
		log.Println("Added Value : " + val + " " + "Added KeyId : " + incrementId)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("KeyId : " + incrementId))
	}
}

type FlushHandler struct{}

func (h *FlushHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPut {
		log.Println("Flush Handler Is Started")
		db.Flush()
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("All Data Are Flushed"))
	}
}

//Main method
func main() {
	log.Println("Api Is Started")
	http.Handle("/add", &AddHandler{})
	http.Handle("/get/", &GetHandler{})
	http.Handle("/flush", &FlushHandler{})
	http.ListenAndServe(":8888", nil)
	log.Println(":8888 Port")
}
