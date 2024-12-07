package routes

import (
	"auth-api/controllers"
	"auth-api/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(r *mux.Router) {
	r.HandleFunc("/signup", controllers.SignUp).Methods("POST")
	r.HandleFunc("/signin", controllers.SignIn).Methods("POST")
	r.Handle("/protected", middlewares.AuthMiddleware(http.HandlerFunc(controllers.Protected))).Methods("GET")
	r.HandleFunc("/revoke", controllers.RevokeToken).Methods("POST")
	r.HandleFunc("/refresh", controllers.Refresh).Methods("POST")
}
