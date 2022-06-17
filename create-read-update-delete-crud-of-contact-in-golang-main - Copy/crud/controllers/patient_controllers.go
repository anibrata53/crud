package controllers

import (
	"crud/database"
	"crud/entity"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//GetAllPaitent get all Paitent data
func GetAllPaitent(w http.ResponseWriter, r *http.Request) {
	var Paitents []entity.Paitent
	database.Connector.Find(&Paitents)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Paitents)
}

// GetPaitentByID returns Paitent with specific ID
func GetPaitentByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var Paitent entity.Paitent
	database.Connector.First(&Paitent, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Paitent)
}

//CreatePaitent creates Paitent
func CreatePaitent(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var Paitent entity.Paitent
	json.Unmarshal(requestBody, &Paitent)

	database.Connector.Create(Paitent)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Paitent)
}

//UpdatePaitentByID updates Paitent with respective ID
func UpdatePaitentByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var Paitent entity.Paitent
	json.Unmarshal(requestBody, &Paitent)
	database.Connector.Save(&Paitent)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Paitent)
}

//DeletPaitentByID delete's Paitent with specific ID
func DeletPaitentByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var Paitent entity.Paitent
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&Paitent)
	w.WriteHeader(http.StatusNoContent)
}

// func Login(w http.ResponseWriter, r *http.Request) {
// requestBody, _ := ioutil.ReadAll(r.Body)
// var login entity.Login
// json.Unmarshal(requestBody, &login)

// database.Connector.Create(login)
// w.Header().Set("Content-Type", "application/json")
// w.WriteHeader(http.StatusCreated)
// json.NewEncoder(w).Encode(login)

// database.Connector.Find(&login)
// w.Header().Set("Content-Type", "application/json")
// w.WriteHeader(http.StatusOK)
// json.NewEncoder(w).Encode(login)
// vars := mux.Vars(r)
// key := vars["id"]

// 	var Paitent entity.Paitent
// 	database.Connector.First(&Login, key)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(Paitent)
// }

// func Login(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	key := vars["id"]

// 	var login entity.Login
// 	database.Connector.First(&login, key)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(login)
// }
type getdataa struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
	Acctype  string `json:"acctype"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	requestBody, _ := ioutil.ReadAll(r.Body)
	var getdata getdataa
	json.Unmarshal(requestBody, &getdata)

	vars := mux.Vars(r)
	key := vars["id"]

	var Paitent entity.Paitent
	database.Connector.First(&Paitent, key)
	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(Paitent.ID)
	fmt.Println(Paitent.ID)

	if Paitent.ID == getdata.ID {
		if Paitent.Password == getdata.Password {

			fmt.Println("Success")
			json.NewEncoder(w).Encode(Paitent)
			// json.NewEncoder(w).Encode(Paitent.Password)
			// json.NewEncoder(w).Encode(Paitent.Acctype)

		} else {
			fmt.Fprintf(w, "invalid password")
		}

	} else {
		fmt.Fprintf(w, "invalid id")
	}

}

//requestBody, _ := ioutil.ReadAll(r.Body)
// 	var getdata getdataa
// 	json.Unmarshal(requestBody, &getdata)

// 	vars := mux.Vars(r)
// 	key := vars["id"]

// 	var Paitent entity.Paitent
// 	database.Connector.First(&Paitent, key)
// 	w.Header().Set("Content-Type", "application/json")
// 	// json.NewEncoder(w).Encode(Paitent)
// 	fmt.Println(getdata.ID)

// 	if Paitent.ID == getdata.ID {

// 		if Paitent.Password == getdata.Password {

// 			// fmt.Println(w, Paitent.ID, Paitent.Acctype)
// 			fmt.Println("hoo")

// 		}
// 	}
// }
