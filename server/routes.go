package server

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/mberlanda/passione-trasferte/handlers"
	"github.com/mberlanda/passione-trasferte/middlewares"
)

const apiIdentifier = "https://bwf1cm8.eu.auth0.com/api/v2/" // "http://localhost:8080" //
const authDomain = "bwf1cm8.eu.auth0.com"

var jwtMiddleware = middlewares.NewJwtMiddleware(apiIdentifier, authDomain) // os.Getenv("AUTH0_DOMAIN"))

func GetRoutes() *mux.Router {
	r := mux.NewRouter()
	n := negroni.Classic()

	n.UseHandler(r)

	// This route is always accessible
	r.Handle("/api/public", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		message := "Hello from a public endpoint! You don't need to be authenticated to see this."
		handlers.ResponseJSON(message, w, http.StatusOK)
	}))

	// This route is only accessible if the user has a valid Access Token
	// We are chaining the jwtmiddleware middleware into the negroni handler function which will check
	// for a valid token.
	r.Handle("/api/private", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			message := "Hello from a private endpoint! You need to be authenticated to see this."
			handlers.ResponseJSON(message, w, http.StatusOK)
		}))))

	r.Handle("/api/private-scoped", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(middlewares.NewScopedMiddleware))))

	return r
}
