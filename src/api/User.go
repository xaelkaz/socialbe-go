package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"net/http"
	"social-golang/src/models"
	"social-golang/src/repository"
	"social-golang/src/util"
)

func AddUser(w http.ResponseWriter, req *http.Request) {

	userRepository := repository.NewUserRepository("User")
	//get variable by path
	params := mux.Vars(req)
	var firstName, lastName, Email string
	firstName = string(params["firstName"])
	lastName = string(params["lastName"])
	Email = string(params["Email"])

	var p models.User
	p.Name = firstName + " " + lastName
	p.Email = Email
	userRepository.Save(&p)

}

func GetUserByEmail(w http.ResponseWriter, req *http.Request) {

	userRepository := repository.NewUserRepository("User")

	//get variable by path
	params := mux.Vars(req)
	var Email string
	Email = string(params["Email"])

	user, err2 := userRepository.FindByEmail(Email)
	if err2 != nil {
		fmt.Println(err2)
	}
	json.NewEncoder(w).Encode(user)

}

func FollowSubject(w http.ResponseWriter, req *http.Request) {

	userRepository := repository.NewUserRepository("User")
	subjectRepository := repository.NewSubjectRepository("Subject")

	//get variable by path
	params := mux.Vars(req)
	var email = string(params["email"])
	var codeSubject = string(params["code"])
	user, _ := userRepository.FindByEmail(email)
	subject, _ := subjectRepository.FindByCode(codeSubject)
	var status = 1
	for i := 0; i < len(user.Subject); i++ {
		fmt.Println("Loop")
		if user.Subject[i].Code == codeSubject {
			json.NewEncoder(w).Encode("false")
			status = 0
			break
		}
	}
	if status == 1 {
		fmt.Println(user.Subject)
		user.Subject = append(user.Subject, subject)
		fmt.Println(user.Subject)
		userRepository.Update(user)
	}
}
func CreateUser(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	// Unmarshal
	var msg models.User
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	userRepository := repository.NewUserRepository("User")

	userTemp, err := userRepository.FindByEmail(msg.Email)
	if userTemp == nil {
		var user models.User
		user.Email = msg.Email
		user.Name = msg.Name
		user.Picture = msg.Picture
		user.StudentId = msg.StudentId
		user.Major = msg.Major
		userRepository.Save(&user)
	} else {
		util.JSON(w, 500, util.T{
			"status":  false,
			"message": "Can not confirm account this email is already used!",
			"err":     err,
		})
		return
	}
}

func UnfollowSubject(w http.ResponseWriter, req *http.Request) {

	userRepository := repository.NewUserRepository("User")

	//get variable by path
	params := mux.Vars(req)
	var email = string(params["email"])
	var codeSubject = string(params["code"])
	user, _ := userRepository.FindByEmail(email)
	var subjects models.SubjectPointer
	for i := 0; i < len(user.Subject); i++ {
		if user.Subject[i].Code != codeSubject {
			subjects = append(subjects, user.Subject[i])
		}
	}
	user.Subject = subjects
	userRepository.Update(user)
}

func GetUserAll(w http.ResponseWriter, req *http.Request) {

	UserRepository := repository.NewUserRepository("User")
	user, err := UserRepository.FindAll()
	if err != nil {
		fmt.Println(err)
		util.JSON(w, 500, util.T{
			"status": 1,
			"error":  err.Error(),
		})
		return
	}
	util.JSON(w, 200, util.T{
		"user":   user,
		"status": 0,
	})
}

func DeleteUserById(w http.ResponseWriter, req *http.Request) {

	UserRepository := repository.NewUserRepository("User")
	params := mux.Vars(req)
	var userId = string(params["id"])
	userIdHex, _ := primitive.ObjectIDFromHex(userId)
	UserRepository.Delete(userIdHex)
}
