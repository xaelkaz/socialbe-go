package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"social-golang/src/models"
	"social-golang/src/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"net/http"
	//"strings"
	"time"
)

func AddRequest(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	// Unmarshal
	var msg models.Request
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	requestRepository := repository.NewRequestRepository("Request")
	userRepository := repository.NewUserRepository("User")
	currentTime := time.Now()

	user, err2 := userRepository.FindByEmail(msg.User.Email)
	if err2 != nil {
		fmt.Println(err2)
	}
	var p models.Request
	p.SubjectCode = msg.SubjectCode
	p.SubjectName = msg.SubjectName
	p.Timestamp = currentTime.Format("3:4:5")
	p.Date = currentTime.Format("2006-01-02")
	p.User = user

	requestRepository.Save(&p)
}
func GetRequestAll(w http.ResponseWriter, req *http.Request) {

	requestRepository := repository.NewRequestRepository("Request")
	feedback, err2 := requestRepository.FindAll()
	if err2 != nil {
		fmt.Println(err2)
	}
	json.NewEncoder(w).Encode(feedback)

}
func DeleteRequest(w http.ResponseWriter, req *http.Request) {

	requestRepository := repository.NewRequestRepository("Request")
	params := mux.Vars(req)
	var requestId = string(params["requestid"])
	requestIdHex, err := primitive.ObjectIDFromHex(requestId)
	request, err := requestRepository.FindByID(requestIdHex)
	err = requestRepository.Delete(request)
	if err != nil {
	}
}
func GetRequestById(w http.ResponseWriter, req *http.Request) {

	requestRepository := repository.NewRequestRepository("Request")

	params := mux.Vars(req)
	var requestid = string(params["requestid"])
	requestIdHex, _ := primitive.ObjectIDFromHex(requestid)
	request, err2 := requestRepository.FindByID(requestIdHex)
	if err2 != nil {
		fmt.Println(err2)
	}
	json.NewEncoder(w).Encode(request)

}
