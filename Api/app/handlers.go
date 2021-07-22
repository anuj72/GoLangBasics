package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserResponse struct {
	ID    int
	Name  string
	Email string
}

func handlerFun(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	resp := []UserResponse{
		{ID: 123,
			Name:  "req.Name",
			Email: "req.Email"},
	}
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

}
