package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name        string
	Age         int
	// Class       int
	// Address     string
	// Present     bool
	// Temperature float64
}

type errorResp struct {
	Statuscode int
	Message string
}

var (
	users = map[string]User{}
)

func main() { 

	http.HandleFunc("/createuser", adduser)
	http.HandleFunc("/users", getusers )

	fmt.Println("user are :",users)
	fmt.Println("Server started")
	log.Fatalf("Server start nahi hua, err : %v\n",  http.ListenAndServe(":8000",nil))

}

// create/Add new record of user in map
func adduser(w http.ResponseWriter , r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		errr := errorResp {
			Statuscode: http.StatusBadRequest,
			Message: "Bro/Sis get wala nahi chalega , dhyan do uspe , err = ",
			
		}
		json.NewEncoder(w).Encode(errr)
		return
	}

	user := User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errr := errorResp {
			Statuscode: http.StatusBadRequest,
			Message: "Bro/Sis jo payload aya hai wo sahi nahi hai , dhyan do uspe , err = "+ err.Error(),
		}
		json.NewEncoder(w).Encode(errr)
		return
	}

	users[user.Name] = user

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(user)


	return
}

// return user record from map 
func getusers(w http.ResponseWriter , r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errr := errorResp {
			Statuscode: http.StatusBadRequest,
			Message: "Bro/Sis payload nahi bana paya , err ="+ err.Error(),
		}
		json.NewEncoder(w).Encode(errr)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println("users are :",users)
	return

}