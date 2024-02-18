package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func dbinit() *gorm.DB {
	db, err := Connect()
	if err != nil {
		log.Fatal("error ", err)
	}
	db.AutoMigrate(&User{})
	return db

}

func JsonResponse(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("Failed to parse json ", err)
		w.WriteHeader(http.StatusBadGateway)
	}
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func ErrResponse(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding the 500 error :", code)
	}
	type err_res struct {
		Error string `json:"error"`
	}
	JsonResponse(w, code, err_res{
		Error: msg,
	})

}
func Getallusers(w http.ResponseWriter, r *http.Request) {
	gormdb := dbinit()
	var user []User
	users := gormdb.Find(&user)
	if users.Error != nil {
		ErrResponse(w, http.StatusInternalServerError, fmt.Sprintln("Error fetching the data from database:", users.Error))
		return
	}

	JsonResponse(w, http.StatusOK, user)
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	gormdb := dbinit()
	user := &User{}
	body := r.Body
	data, err := io.ReadAll(body)
	if err != nil {
		log.Println("Error parsing the request body:", err)
	}

	err = json.Unmarshal(data, user)
	if err != nil {
		ErrResponse(w, http.StatusBadRequest, fmt.Sprint("error during reading the body of the request:", err))
		return
	}

	Createduser := gormdb.Create(&user)
	rows := Createduser.RowsAffected
	if Createduser.Error != nil {
		ErrResponse(w, http.StatusBadGateway, fmt.Sprint("error craeting the user:", err))
	}
	fmt.Println("Added a user \n >>Rows affected: ", rows)

}
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	gormdb := dbinit()
	id := mux.Vars(r)
	ID, err := strconv.Atoi(id["id"])
	if err != nil {
		ErrResponse(w, http.StatusBadRequest, fmt.Sprint("error in conversion of id:", err))
	}
	var user User
	founduser := gormdb.Find(&user, "id = ?", ID)
	if user.ID == 0 {
		ErrResponse(w, http.StatusOK, fmt.Sprint("No such user for id:", ID))
	}
	if founduser.Error != nil {
		ErrResponse(w, http.StatusBadRequest, fmt.Sprint("Could not find user:", founduser.Error))
		return
	}

	JsonResponse(w, http.StatusOK, user)

}
func Delete(w http.ResponseWriter, r *http.Request) {
	gormdb := dbinit()
	id := mux.Vars(r)
	ID, err := strconv.Atoi(id["id"])
	if err != nil {
		ErrResponse(w, http.StatusBadRequest, fmt.Sprint("error in conversion of id:", err))
	}
	del := gormdb.Delete(&User{}, ID)
	if del.Error != nil {
		ErrResponse(w, http.StatusBadGateway, fmt.Sprint("Could not delete user:", del.Error))
		return
	}
	rows := del.RowsAffected
	fmt.Println("Deleted a user \n >>Rows affected: ", rows)
	JsonResponse(w, http.StatusOK, ID)

}
func Update(w http.ResponseWriter, r *http.Request) {
	gormdb := dbinit()
	id := mux.Vars(r)
	ID, err := strconv.Atoi(id["id"])
	if err != nil {
		ErrResponse(w, http.StatusBadRequest, fmt.Sprint("error in conversion of id:", err))
	}
	var user User
	users := gormdb.Find(&user, ID)
	if users.Error != nil {
		ErrResponse(w, http.StatusInternalServerError, fmt.Sprintln("Error fetching the data from database:", users.Error))
		return
	}

	body := r.Body
	data, err := io.ReadAll(body)
	if err != nil {
		log.Println("Error parsing the request body:", err)
	}

	err = json.Unmarshal(data, &user)
	if err != nil {
		ErrResponse(w, http.StatusBadRequest, fmt.Sprint("error during reading the body of the request:", err))
		return
	}
	updated := gormdb.Save(&user)
	if updated.Error != nil {
		ErrResponse(w, http.StatusBadGateway, fmt.Sprint("Could not update user:", updated.Error))
		return

	}
	rows := updated.RowsAffected
	fmt.Println("Updated a user \n >>Rows affected: ", rows)
	JsonResponse(w, http.StatusOK, ID)

}
