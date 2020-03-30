package route

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"social-golang/src/api"
	"social-golang/src/config"
	"log"
	"net/http"
	"os"
)

func Init() {
	r := mux.NewRouter()

	// Serve static file
	go func() {
		http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
		http.ListenAndServe(":"+config.ServerConfig.StaticPort, nil)
	}()

	// Middleware
	//r.Use(util.JwtAuthentication)

	// API r
	setRouter(r)

	listen(r, config.ServerConfig.Port)
}

func setRouter(r *mux.Router) {
	//Subject
	r.HandleFunc("/subject/{name}/{code}/{majorName}", api.AddSubject).Methods("GET")
	r.HandleFunc("/subject/{majorName}", api.GetSubjectByMajor).Methods("GET")
	r.HandleFunc("/subjectbycode/{code}", api.GetSubjectByCode).Methods("GET")
	r.HandleFunc("/subjects", api.GetSubjectAll).Methods("GET")
	r.HandleFunc("/subjectbyemail/{major}/{email}", api.GetSubjectByMajorEmail).Methods("GET")
	r.HandleFunc("/subjectfromuser/{email}", api.GetSubjectFromUser).Methods("GET")
	r.HandleFunc("/deletesubject/{code}", api.DeleteSubject).Methods("GET")
	r.HandleFunc("/subject", api.CreateSubject).Methods("POST")
	//Faculty
	r.HandleFunc("/faculty/{name}", api.AddFaculty).Methods("GET")
	r.HandleFunc("/facultyemail/{email}", api.GetFacultyByEmail).Methods("GET")
	r.HandleFunc("/faculties", api.GetFacultyAll).Methods("GET")
	r.HandleFunc("/deletefaculty/{facultyname}", api.DeleteFaculty).Methods("GET")
	//Major
	r.HandleFunc("/major/{name}/{facultyName}", api.AddMajor).Methods("GET")
	r.HandleFunc("/major/{facultyName}", api.GetMajorByFaculty).Methods("GET")
	r.HandleFunc("/majorbyemail/{facultyName}/{email}", api.GetMajorByFacultyEmail).Methods("GET")
	r.HandleFunc("/deletemajor/{majorname}", api.DeleteMajor).Methods("GET")
	r.HandleFunc("/majors", api.GetMajorAll).Methods("GET")
	// User Ping
	r.HandleFunc("/createuser", api.CreateUser).Methods("POST")
	//User
	r.HandleFunc("/users", api.GetUserAll).Methods("GET")
	r.HandleFunc("/deleteuser/{id}", api.DeleteUserById).Methods("GET")
	r.HandleFunc("/user/{Email}", api.GetUserByEmail).Methods("GET")
	r.HandleFunc("/follow/{email}/{code}", api.FollowSubject).Methods("GET")
	r.HandleFunc("/user/{firstName}/{lastName}/{Email}", api.AddUser).Methods("GET")
	r.HandleFunc("/unfollow/{email}/{code}", api.UnfollowSubject).Methods("GET")
	//Post
	r.HandleFunc("/postvdo/{text}/{email}/{code}/{vdoLink}", api.AddPost).Methods("GET")
	r.HandleFunc("/postfile/{text}/{email}/{code}/{name}/{token}", api.AddPost).Methods("GET")
	r.HandleFunc("/postfull/{text}/{email}/{code}/{vdoLink}/{name}/{token}", api.AddPost).Methods("GET")
	r.HandleFunc("/post/{text}/{email}/{code}", api.AddPost).Methods("POST")
	r.HandleFunc("/post", api.AddPost).Methods("POST")
	r.HandleFunc("/posts", api.GetPostAll).Methods("GET")
	r.HandleFunc("/post/{code}", api.GetPostByCode).Methods("GET")
	r.HandleFunc("/deletepost/{postid}", api.DeletePost).Methods("GET")
	r.HandleFunc("/getpost/{postid}", api.GetPostById).Methods("GET")
	//comment
	r.HandleFunc("/comments", api.GetCommentAll).Methods("GET")
	r.HandleFunc("/commenttext/{text}", api.AddComment).Methods("GET")
	r.HandleFunc("/comment/{postID}", api.AddComment).Methods("POST")
	r.HandleFunc("/deletecomment/{id}", api.DeleteCommentById).Methods("GET")
	r.HandleFunc("/deletecomment/{id}/{postid}", api.DeleteCommentByIdANDPostId).Methods("GET")
	//feedback
	r.HandleFunc("/feedback", api.AddFeedback).Methods("POST")
	r.HandleFunc("/deletefeedback/{id}", api.DeleteFeedback).Methods("GET")
	r.HandleFunc("/feedbacks", api.GetFeedbackAll).Methods("GET")
	//request
	r.HandleFunc("/request", api.AddRequest).Methods("POST")
	r.HandleFunc("/deleterequest/{id}", api.DeleteRequest).Methods("GET")
	r.HandleFunc("/requests", api.GetRequestAll).Methods("GET")
	//log.Fatal(http.ListenAndServe(getPort(), handlers.CORS(handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD"}), handlers.AllowedOrigins([]string{"*"}))(r)))

}
func listen(r *mux.Router, port string) {
	log.Printf("Server listening on port %v...\n", port)

	err := http.ListenAndServe(getPort(), handlers.CORS(handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD"}), handlers.AllowedOrigins([]string{"*"}))(r))

	if err != nil {
		log.Println("Serve server fail", err)
	}
}
func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		// port = "12345"
		port = "8000"
	}
	return ":" + port
}
