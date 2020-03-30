package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"social-golang/src/models"
	"social-golang/src/repository"
	"net/http"
)

func AddFavoriteSubject(w http.ResponseWriter, req *http.Request) {

	favoriteSubjectRepository := repository.NewFavoriteSubjectRepository("FavoriteSubject")
	userRepository := repository.NewUserRepository("User")
	subjectRepository := repository.NewSubjectRepository("Subject")

	params := mux.Vars(req)
	var email = string(params["email"])
	var code = string(params["code"])
	user, err2 := userRepository.FindByEmail(email)
	if err2 != nil {
		fmt.Println(err2)
	}

	subject, err3 := subjectRepository.FindByCode(code)
	if err3 != nil {
		fmt.Println(err3)
	}
	var p models.FavoriteSubject
	p.User = user
	p.Subject = subject
	favoriteSubjectRepository.Save(&p)

}

//Default data FavoriteSubject
func AddFavoriteSubjectDefault(email string, code string) {

	favoriteSubjectRepository := repository.NewFavoriteSubjectRepository("FavoriteSubject")
	subjectRepository := repository.NewSubjectRepository("Subject")
	userRepository := repository.NewUserRepository("User")

	user, err2 := userRepository.FindByEmail(email)
	if err2 != nil {
		fmt.Println(err2)
	}

	subject, err3 := subjectRepository.FindByCode(code)
	if err3 != nil {
		fmt.Println(err3)
	}
	var p models.FavoriteSubject
	p.User = user
	p.Subject = subject
	favoriteSubjectRepository.Save(&p)
}
func GetFavoriteSubjectByUserEmail(w http.ResponseWriter, req *http.Request) {
	//


	favoriteSubjectRepository := repository.NewFavoriteSubjectRepository("FavoriteSubject")

	params := mux.Vars(req)
	var email = string(params["email"])
	user, err2 := favoriteSubjectRepository.FindByEmail(email)
	if err2 != nil {
		fmt.Println(err2)
	}
	json.NewEncoder(w).Encode(user)

}
