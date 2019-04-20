package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sp19-281-ace-traordinary/Backend/userapi/dao"
	"github.com/sp19-281-ace-traordinary/Backend/userapi/models"
	"github.com/sp19-281-ace-traordinary/Backend/userapi/utils"
)

var user []models.User

// RegisterUserHandler creta a user
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entered the function RegisterUserEndpoint")
	var user models.Registration
	_ = json.NewDecoder(r.Body).Decode(&user)
	//Need to Remove unecessary Comments
	log.Printf("Entered the function RegisterUserEndpoint")
	fmt.Println(user)
	user.Verificationcode = utils.GenerateVerificationTocken()
	res, Message := dao.RegisterUserDao(user)
	if res == false {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(Message))
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

//ConfirmRegistrationrHandler will return all users
func ConfirmRegistrationrHandler(w http.ResponseWriter, r *http.Request) {

}

//ResendConfirmHandler wsends confirmation mail again
func ResendConfirmHandler(w http.ResponseWriter, r *http.Request) {

}

//LoginUserHandler login specific user
func LoginUserHandler(w http.ResponseWriter, r *http.Request) {

}

//ForgotPasswordrHandler will return all users
func ForgotPasswordrHandler(w http.ResponseWriter, r *http.Request) {

}

//DeleteUserHandler Deletes the specific user from Database
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

}

//GetAllUsersHandler will return all users
func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	payload := dao.GetAllUsersDao()
	json.NewEncoder(w).Encode(payload)
}

/**
// GetUserEndpoint gets a user
func GetUserEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	payload := dao.GetAllUsers()
	for _, p := range payload {
		if p.Userid == params["userid"] {
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	json.NewEncoder(w).Encode("User not found")
}

// GetAllUserEndpoint gets all user
func GetAllUserEndpoint(w http.ResponseWriter, r *http.Request) {
	payload := dao.GetAllUsers()
	json.NewEncoder(w).Encode(payload)
}

// RegisterUserEndpoint creta a user
func RegisterUserEndpoint(w http.ResponseWriter, r *http.Request) {
	var user models.User
	fmt.Println("In RegisterUserEndpoint")
	_ = json.NewDecoder(r.Body).Decode(&user)
	dao.InsertOneValue(user)
	json.NewEncoder(w).Encode(user)
}

// DeleteUserEndpoint delets a user
func DeleteUserEndpoint(w http.ResponseWriter, r *http.Request) {
	var user models.User

	_ = json.NewDecoder(r.Body).Decode(&user)
	dao.DeleteUser(user)
}

/**func UpdateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	personID := mux.Vars(r)["id"]
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&person)
	dao.UpdatePerson(person, personID)

}**/