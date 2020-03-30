package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"social-golang/src/models"
	"social-golang/src/repository"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func AddSubject(w http.ResponseWriter, req *http.Request) {
	subjectRepository := repository.NewSubjectRepository("Subject")
	majorRepository := repository.NewMajorRepository("Major")

	//get variable by path
	params := mux.Vars(req)
	var name, code, majorName string
	code = string(params["code"])
	name = string(params["name"])
	majorName = string(params["majorName"])

	major, err2 := majorRepository.FindByName(majorName)
	if err2 != nil {
		fmt.Println(err2)
	}
	var p models.Subject
	p.Name = name
	p.Code = code
	p.Major = major
	subjectRepository.Save(&p)

}

func GetSubjectByMajor(w http.ResponseWriter, req *http.Request) {
	//
	subjectRepository := repository.NewSubjectRepository("Subject")
	//get variable by path
	params := mux.Vars(req)
	var majorName = string(params["majorName"])

	subject, err2 := subjectRepository.FindByMajor(majorName)
	if err2 != nil {
		fmt.Println(err2)
	}
	json.NewEncoder(w).Encode(subject)

}

func GetSubjectByCode(w http.ResponseWriter, req *http.Request) {
	//
	subjectRepository := repository.NewSubjectRepository("Subject")
	//get variable by path
	params := mux.Vars(req)
	var code = string(params["code"])
	if _, err := strconv.Atoi(code); err == nil || strings.Contains(code, "*") {
		subject, err2 := subjectRepository.FindByCodeEx(code)
		json.NewEncoder(w).Encode(subject)
		if err2 != nil {
			fmt.Println(err2)
		}
	} else {
		subject, err2 := subjectRepository.FindByNameEx(code)
		json.NewEncoder(w).Encode(subject)
		if err2 != nil {
			fmt.Println(err2)
		}
	}
}

func GetSubjectAll(w http.ResponseWriter, req *http.Request) {

	subjectRepository := repository.NewSubjectRepository("Subject")
	post, err2 := subjectRepository.FindAll()
	if err2 != nil {
		fmt.Println(err2)
	}
	json.NewEncoder(w).Encode(post)
}

func GetSubjectByMajorEmail(w http.ResponseWriter, req *http.Request) {
	//

	userRepository := repository.NewUserRepository("User")
	params := mux.Vars(req)
	var majorName = string(params["major"])
	var email = string(params["email"])

	user, _ := userRepository.FindByEmail(email)
	var subjects models.SubjectPointer
	for i := 0; i < len(user.Subject); i++ {
		if user.Subject[i].Major.Name == majorName {
			subjects = append(subjects, user.Subject[i])
		}
	}
	json.NewEncoder(w).Encode(subjects)

}

func GetSubjectFromUser(w http.ResponseWriter, req *http.Request) {

	userRepository := repository.NewUserRepository("User")
	params := mux.Vars(req)
	var email = string(params["email"])

	user, _ := userRepository.FindByEmail(email)
	var subjects models.SubjectPointer
	for i := 0; i < len(user.Subject); i++ {
		subjects = append(subjects, user.Subject[i])
	}
	json.NewEncoder(w).Encode(subjects)
}
func DeleteSubject(w http.ResponseWriter, req *http.Request) {

	subjectRepository := repository.NewSubjectRepository("Subject")
	params := mux.Vars(req)
	var code = string(params["code"])
	subjectRepository.DeleteByCode(code)
}

func CreateSubject(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	// Unmarshal
	var msg models.Subject
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	subjectRepository := repository.NewSubjectRepository("Subject")
	majorRepository := repository.NewMajorRepository("Major")

	major, err2 := majorRepository.FindByName(msg.Major.Name)
	if err2 != nil {
		fmt.Println(err2)
	}
	var p models.Subject
	p.Name = msg.Name
	p.Code = msg.Code
	p.Picture = msg.Picture
	p.Major = major
	subjectRepository.Save(&p)
}
