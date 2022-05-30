package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"golang.api/controllers"
	"golang.api/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalities", controllers.AddPersonality).Methods("Post")
	r.HandleFunc("/personalities", controllers.FindAllPersonalities).Methods("Get")
	r.HandleFunc("/personalities/{id}", controllers.FindPersonalityById).Methods("Get")
	r.HandleFunc("/personalities/{id}", controllers.RemovePersonality).Methods("Delete")
	r.HandleFunc("/personalities/{id}", controllers.EditPersonality).Methods("Put")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))

}
