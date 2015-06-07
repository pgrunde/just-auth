package server

import (
	"log"
	"net/http"
	"path"
	"text/template"

	sql "github.com/aodin/aspect"
	"github.com/aodin/volta/config"
)

const staticURL = "/static/"

// Server starts the application
type Server struct {
	Conn sql.Connection
	Conf config.Config
}

// db calls methods that interact with the database
type db struct {
	Conn sql.Connection
}

// ListenAndServe starts the application on the config address
func (server *Server) ListenAndServe() error {
	log.Printf("server starting on address %s\n", server.Conf.Address())
	return http.ListenAndServe(server.Conf.Address(), nil)
}

// New spins off handler functions at given routes and returns the server
func New(conf config.Config, conn sql.Connection) *Server {
	http.HandleFunc("/", root)
	http.HandleFunc("/signin", signin)
	http.HandleFunc("/create-account", createAccount)
	http.HandleFunc("/favicon.ico", favicon)
	return &Server{Conf: conf, Conn: conn}
}

// root template
var rootTmpl = template.Must(
	template.New("index").Delims("<%", "%>").ParseFiles("./index.html"),
)

// root handler
func root(w http.ResponseWriter, r *http.Request) {
	attrs := struct {
		StaticURL string
	}{
		StaticURL: staticURL,
	}
	rootTmpl.ExecuteTemplate(w, "index", attrs)
}

func favicon(w http.ResponseWriter, r *http.Request) {
	faviconName := path.Base(r.URL.Path)
	http.ServeFile(w, r, "/static/"+faviconName)
}
