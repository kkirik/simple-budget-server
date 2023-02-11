package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"simple-budget-backend/helpers"
	"simple-budget-backend/interfaces"
	"simple-budget-backend/users"

	"github.com/gorilla/mux"
)

func login(w http.ResponseWriter, r *http.Request) {
	// Read body
	body, err := ioutil.ReadAll(r.Body)

	helpers.HandleErr(err)
	// Handle Login
	var formattedBody interfaces.Login
	err = json.Unmarshal(body, &formattedBody)

	helpers.HandleErr(err)

	login := users.Login(formattedBody.Username, formattedBody.Password)

	// Prepare response
	if login["message"] == "all is fine" {
		resp := login
		json.NewEncoder(w).Encode(resp)
		// Handle error in else
	} else {
		resp := interfaces.ErrResponse{Message: "Wrong username or password"}
		json.NewEncoder(w).Encode(resp)
	}
}

func StartApi() {
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	fmt.Println("App is working on port :8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
