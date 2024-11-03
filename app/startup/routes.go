package startup

import (
	"earn-expense/app/controllers"
	"github.com/jinzhu/gorm"
	"net/http"
)

func (a *App) InitializeRoutes() {
	a.Get("/", a.HandleRequest(controllers.Home))

	// Routing for auth
	a.Post("/login", a.HandleRequest(controllers.Login))

	// Routing for handling the users
	a.Get("/users", a.HandleRequest(controllers.GetAllUsers))
	a.Post("/users", a.HandleRequest(controllers.CreateUser))
	a.Get("/users/{id:[0-9]+}", a.HandleRequest(controllers.GetUserById))
	a.Put("/users/{id:[0-9]+}", a.HandleRequest(controllers.UpdateUser))
	a.Delete("/users/{id:[0-9]+}", a.HandleRequest(controllers.DeleteUser))
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func (a *App) HandleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}
