package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var dataStore = make(map[string]string, 3)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleHome)
	mux.HandleFunc("/getData", getData)
	mux.HandleFunc("/getAll", getAll)
	mux.HandleFunc("/postData", postData)
	mux.HandleFunc("/deleteData", deleteData)

	server := http.Server{
		Addr:         ":9090",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
		Handler:      mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}

}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home"))
}

func getData(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		val, ok := r.URL.Query()["key"]

		if !ok {
			w.Write([]byte("provide key to be deleted"))
			return
		}
		data, ok := dataStore[val[0]]

		if !ok {
			w.Write([]byte("post the provided val first to get"))
			return
		}
		w.Write(([]byte(data)))
	} else {
		w.Write([]byte("Invalid Method"))
	}
}

func getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := json.Marshal(dataStore)
		checkErr(err)

		w.Write(data)
	} else {
		w.Write([]byte("Invalid Method"))
	}
}

func postData(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		key := r.FormValue("key")
		data := r.FormValue("data")

		if len(dataStore) > 2 {
			w.Write([]byte("delete a few items before adding, dataStore is full"))
			return
		}
		dataStore[key] = data
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func deleteData(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		val, ok := r.URL.Query()["key"]

		if !ok {
			w.Write([]byte("provide key to return value"))
			return
		}
		_, ok = dataStore[val[0]]

		if !ok {
			w.Write([]byte("provided value dose"))
			return
		}

		delete(dataStore, val[0])
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
