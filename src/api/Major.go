package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"social-golang/src/models"
	"social-golang/src/repository"
	"net/http"
)

func AddMajor(w http.ResponseWriter, req *http.Request) {

	majorRepository := repository.NewMajorRepository("Major")
	facultyRepository := repository.NewFacultyRepositoryMongo("Faculty")

	params := mux.Vars(req)
	var facultyName string
	facultyName = string(params["facultyName"])
	faculty, err2 := facultyRepository.FindByName(facultyName)
	if err2 != nil {
		fmt.Println(err2)
	}
	var name string
	name = string(params["name"])

	var p models.Major
	p.Name = name
	p.Faculty = faculty
	majorRepository.Save(&p)

}

func GetMajorByFaculty(w http.ResponseWriter, req *http.Request) {
	majorRepository := repository.NewMajorRepository("Major")

	params := mux.Vars(req)
	var facultyName = string(params["facultyName"])
	major, err2 := majorRepository.FindByFaculty(facultyName)
	if err2 != nil {
		fmt.Println(err2)
	}
	json.NewEncoder(w).Encode(major)

}

func GetMajorByFacultyEmail(w http.ResponseWriter, req *http.Request) {

	majorRepository := repository.NewMajorRepository("Major")
	userRepository := repository.NewUserRepository("User")
	params := mux.Vars(req)
	var facultyName = string(params["facultyName"])
	var email = string(params["email"])
	user, _ := userRepository.FindByEmail(email)
	var majors models.MajorPointer
	var isHave = false
	var tempMajor []string
	for i := 0; i < len(user.Subject); i++ {
		major, err := majorRepository.FindByName(user.Subject[i].Major.Name)
		if err != nil {
		}
		for j := 0; j < len(tempMajor); j++ {
			if majors[j].Name == major.Name {
				isHave = true
			} else {
				isHave = false
				continue
			}
		}
		if major.Faculty.Name == facultyName && !isHave {
			tempMajor = append(tempMajor, major.Name)
			majors = append(majors, major)
		}
	}
	json.NewEncoder(w).Encode(majors)
}
func DeleteMajor(w http.ResponseWriter, req *http.Request) {
	majorRepository := repository.NewMajorRepository("Major")
	params := mux.Vars(req)
	var majorName = string(params["majorname"])
	majorRepository.DeleteByName(majorName)
}

func GetMajorAll(w http.ResponseWriter, req *http.Request) {
	majorRepository := repository.NewMajorRepository("Major")
	major, err2 := majorRepository.FindAll()
	if err2 != nil {
		fmt.Println(err2)
	}
	json.NewEncoder(w).Encode(major)
}
