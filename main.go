package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const APIKey = "qwerklj1230dsa350123l2k1j4kl1j24"

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/test-1", RequestMiddleware(Test1)).Methods("POST")

	port := 8082
	fmt.Println("Playsee Backend Interview Test server is running at port", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), router)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	handleRequests()
}
