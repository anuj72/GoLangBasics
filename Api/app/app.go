package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()
	//define routes
	router.HandleFunc("/greet", handlerFun).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	router.HandleFunc("/customers", getAllCustomersPost).Methods(http.MethodPost)
	//starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func getAllCustomersPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "post req Received")
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}
