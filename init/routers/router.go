package routers

import (
	"net/http"
	"os"

	apihandlers "github.com/Mohammed-Aadil/ds-storage/api/handlers"

	"github.com/Mohammed-Aadil/ds-storage/config"
	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

// RequestHandlerFunc request handler func type
type RequestHandlerFunc func(http.ResponseWriter, *http.Request)

// AppRouter app router
type AppRouter struct {
	router       *mux.Router
	RouteHandler http.Handler
}

// Init initializes the app with predefined configuration
func Init() http.Handler {
	appRouter := &AppRouter{}
	appRouter.router = mux.NewRouter()
	appRouter.setRouters()
	return appRouter.RouteHandler
}

func setCorsOptions(handler http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedOrigins(config.AllowedOrigins),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"}),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "GET", "OPTIONS"}),
	)(handler)
}

func (a *AppRouter) setRouters() {
	a.Get("/", handleRequest(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is a catch-all route"))
	}))
	a.Post("/storage", handleRequest(apihandlers.Upload))
	a.Get("/storage/{id}", handleRequest(apihandlers.ListOne))
	a.Get("/storage/{id}/download", handleRequest(apihandlers.DownloadFiles))

	a.RouteHandler = setCorsOptions(handlers.LoggingHandler(os.Stdout, a.router))
}

// Get wraps the router for GET method
func (a *AppRouter) Get(path string, f http.HandlerFunc) {
	a.router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for GET method
func (a *AppRouter) Post(path string, f http.HandlerFunc) {
	a.router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for GET method
func (a *AppRouter) Put(path string, f http.HandlerFunc) {
	a.router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for GET method
func (a *AppRouter) Delete(path string, f http.HandlerFunc) {
	a.router.HandleFunc(path, f).Methods("DELETE")
}

// Option wraps the router for GET method
func (a *AppRouter) Option(path string, f http.HandlerFunc) {
	a.router.HandleFunc(path, f).Methods("OPTION")
}

func handleRequest(handler RequestHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}
