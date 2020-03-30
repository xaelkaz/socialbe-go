package route

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"social-golang/src/api"
	"social-golang/src/config"
	"social-golang/src/util"
)

func Init() {
	r := mux.NewRouter()

	// Serve static file
	go func() {
		http.Handle("/api/static/", http.StripPrefix("/api/static", http.FileServer(http.Dir("static"))))
		http.ListenAndServe(":"+config.ServerConfig.StaticPort, nil)
	}()

	// Middleware
	r.Use(util.JwtAuthentication)

	// API r
	setRouter(r)

	listen(r, config.ServerConfig.Port)
}

func setRouter(r *mux.Router) {

	//Login
	r.HandleFunc("/api/auth/login", api.LoginUser).Methods("POST")

	//Subject
	r.HandleFunc("/api/subject/{name}/{code}/{majorName}", api.AddSubject).Methods("GET")
	r.HandleFunc("/api/subject/{majorName}", api.GetSubjectByMajor).Methods("GET")
	r.HandleFunc("/api/subjectbycode/{code}", api.GetSubjectByCode).Methods("GET")
	r.HandleFunc("/api/subjects", api.GetSubjectAll).Methods("GET")
	r.HandleFunc("/api/subjectbyemail/{major}/{email}", api.GetSubjectByMajorEmail).Methods("GET")
	r.HandleFunc("/api/subjectfromuser/{email}", api.GetSubjectFromUser).Methods("GET")
	r.HandleFunc("/api/deletesubject/{code}", api.DeleteSubject).Methods("GET")
	r.HandleFunc("/api/subject", api.CreateSubject).Methods("POST")
	//Faculty
	r.HandleFunc("/api/faculty/{name}", api.AddFaculty).Methods("GET")
	r.HandleFunc("/api/facultyemail/{email}", api.GetFacultyByEmail).Methods("GET")
	r.HandleFunc("/api/faculties", api.GetFacultyAll).Methods("GET")
	r.HandleFunc("/api/deletefaculty/{facultyname}", api.DeleteFaculty).Methods("GET")
	//Major
	r.HandleFunc("/api/major/{name}/{facultyName}", api.AddMajor).Methods("GET")
	r.HandleFunc("/api/major/{facultyName}", api.GetMajorByFaculty).Methods("GET")
	r.HandleFunc("/api/majorbyemail/{facultyName}/{email}", api.GetMajorByFacultyEmail).Methods("GET")
	r.HandleFunc("/api/deletemajor/{majorname}", api.DeleteMajor).Methods("GET")
	r.HandleFunc("/api/majors", api.GetMajorAll).Methods("GET")
	// User Ping
	r.HandleFunc("/api/auth/createuser", api.CreateUser).Methods("POST")
	//User
	r.HandleFunc("/api/users", api.GetUserAll).Methods("GET")
	r.HandleFunc("/api/deleteuser/{id}", api.DeleteUserById).Methods("GET")
	r.HandleFunc("/api/user/{Email}", api.GetUserByEmail).Methods("GET")
	r.HandleFunc("/api/follow/{email}/{code}", api.FollowSubject).Methods("GET")
	r.HandleFunc("/api/user/{firstName}/{lastName}/{Email}", api.AddUser).Methods("GET")
	r.HandleFunc("/api/unfollow/{email}/{code}", api.UnfollowSubject).Methods("GET")
	//Post
	r.HandleFunc("/api/postvdo/{text}/{email}/{code}/{vdoLink}", api.AddPost).Methods("GET")
	r.HandleFunc("/api/postfile/{text}/{email}/{code}/{name}/{token}", api.AddPost).Methods("GET")
	r.HandleFunc("/api/postfull/{text}/{email}/{code}/{vdoLink}/{name}/{token}", api.AddPost).Methods("GET")
	r.HandleFunc("/api/post/{text}/{email}/{code}", api.AddPost).Methods("POST")
	r.HandleFunc("/api/post", api.AddPost).Methods("POST")
	r.HandleFunc("/api/posts", api.GetPostAll).Methods("GET")
	r.HandleFunc("/api/post/{code}", api.GetPostByCode).Methods("GET")
	r.HandleFunc("/api/deletepost/{postid}", api.DeletePost).Methods("GET")
	r.HandleFunc("/api/getpost/{postid}", api.GetPostById).Methods("GET")
	//comment
	r.HandleFunc("/api/comments", api.GetCommentAll).Methods("GET")
	r.HandleFunc("/api/commenttext/{text}", api.AddComment).Methods("GET")
	r.HandleFunc("/api/comment/{postID}", api.AddComment).Methods("POST")
	r.HandleFunc("/api/nestedcomment/{postID}/{replyID}", api.AddNestedComment).Methods("POST")

	r.HandleFunc("/api/deletecomment/{id}", api.DeleteCommentById).Methods("GET")
	r.HandleFunc("/api/deletecomment/{id}/{postid}", api.DeleteCommentByIdANDPostId).Methods("GET")
	//feedback
	r.HandleFunc("/api/feedback", api.AddFeedback).Methods("POST")
	r.HandleFunc("/api/deletefeedback/{id}", api.DeleteFeedback).Methods("GET")
	r.HandleFunc("/api/feedbacks", api.GetFeedbackAll).Methods("GET")
	//request
	r.HandleFunc("/api/request", api.AddRequest).Methods("POST")
	r.HandleFunc("/api/deleterequest/{id}", api.DeleteRequest).Methods("GET")
	r.HandleFunc("/api/requests", api.GetRequestAll).Methods("GET")
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
