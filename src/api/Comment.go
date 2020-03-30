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
	"time"
)

func AddComment(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	// Unmarshal
	var msg models.Comment
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	//

	commentRepository := repository.NewCommentRepository("Comment")
	postRepository := repository.NewPostRepository("Post")
	userRepository := repository.NewUserRepository("User")
	currentTime := time.Now()
	//
	params := mux.Vars(req)
	var postID string
	postID = string(params["postID"])
	user, err := userRepository.FindByEmail(msg.User.Email)
	if err != nil {
		fmt.Println(err)
	}
	var p models.Comment
	p.Text = msg.Text
	p.ID = primitive.NewObjectID()
	p.Timestamp = currentTime.Format("3:4:5")
	p.Date = currentTime.Format("2006-01-02")
	p.User = user
	fmt.Println("bson:", p.ID)
	commentRepository.Save(&p)
	postIDHex, err := primitive.ObjectIDFromHex(postID)
	comment, err := commentRepository.FindByID(p.ID)
	post, err := postRepository.FindByID(postIDHex)
	fmt.Println(post.ID)
	post.Comment = append(post.Comment, comment)
	postRepository.Update(post)

}

func GetCommentAll(w http.ResponseWriter, req *http.Request) {

	commentRepository := repository.NewCommentRepository("Comment")
	comment, err2 := commentRepository.FindAll()
	if err2 != nil {
		fmt.Println(err2)
	}
	json.NewEncoder(w).Encode(comment)
}

func DeleteCommentById(w http.ResponseWriter, req *http.Request) {

	commentRepository := repository.NewCommentRepository("Comment")
	params := mux.Vars(req)
	var commentId = string(params["id"])
	commentIdHex, _ := primitive.ObjectIDFromHex(commentId)
	commentRepository.Delete(commentIdHex)
}

func DeleteCommentByIdANDPostId(w http.ResponseWriter, req *http.Request) {

	commentRepository := repository.NewCommentRepository("Comment")
	postRepository := repository.NewPostRepository("Post")
	params := mux.Vars(req)
	var commentId = string(params["id"])
	var postId = string(params["postid"])
	postIdHex, _ := primitive.ObjectIDFromHex(postId)
	commentIdHex, _ := primitive.ObjectIDFromHex(commentId)
	post, _ := postRepository.FindByID(postIdHex)
	var comments models.CommentPointer
	for i := 0; i < len(post.Comment); i++ {
		if post.Comment[i].ID != commentIdHex {
			comments = append(comments, post.Comment[i])
		}
	}
	post.Comment = comments
	postRepository.Update(post)
	commentRepository.Delete(commentIdHex)
}
