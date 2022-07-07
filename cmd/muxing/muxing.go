package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", getParamName).Methods(http.MethodGet)
	router.HandleFunc("/bad", getBad).Methods(http.MethodGet)
	router.HandleFunc("/data", postData).Methods(http.MethodPost)
	router.HandleFunc("/headers", postHeaders).Methods(http.MethodPost)
	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

func getParamName(w http.ResponseWriter, r *http.Request) {

	data := mux.Vars(r)
	val := data["PARAM"]
	w.WriteHeader(200)
	json.NewEncoder(w).Encode("Hello," + val)

}

func getBad(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(500)
}

func postHeaders(w http.ResponseWriter, r *http.Request) {

	avalue, _ := strconv.Atoi(r.Header.Get("a"))
	bvalue, _ := strconv.Atoi(r.Header.Get("b"))

	w.Header().Add("a+b", strconv.Itoa(avalue+bvalue))
	w.WriteHeader(200)

}
func postData(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	json.NewEncoder(w).Encode("I got message:\n" + string(data))
	w.WriteHeader(200)
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
