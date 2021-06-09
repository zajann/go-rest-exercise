package myapp

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	Location  string    `json:"location,omitempty"`
	CreatedAt time.Time `json:"create_at"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func getAllUserInfoHandelr(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GET /users/all")
}

func getUserInfoByIDHandelr(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GET /users/{id}")
}

func updateUserInfoHandelr(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "PUT /users")
}

func createUserInfoHandelr(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "POST /users")
}

func deleteUserInfoByIDHandelr(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "DELETE /users/{id}")
}

func authMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("AuthMiddleware Enter")

		auth := r.Header.Get("Auth")
		if auth == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("need auth"))
			return
		}
		f(w, r)
	})
}

func NewHandler() http.Handler {

	router := mux.NewRouter()

	router.HandleFunc("/", indexHandler).Methods("GET")
	router.HandleFunc("/users/all", getAllUserInfoHandelr).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", authMiddleware(getUserInfoByIDHandelr)).Methods("GET")
	router.HandleFunc("/users", createUserInfoHandelr).Methods("POST")
	router.HandleFunc("/users", updateUserInfoHandelr).Methods("PUT")
	router.HandleFunc("/users/{id:[0-9]+}", authMiddleware(deleteUserInfoByIDHandelr)).Methods("DELETE")

	return router
}
