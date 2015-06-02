package server

import (
	"log"
	"net/http"
	"path"

	sql "github.com/aodin/aspect"
	"github.com/aodin/volta/config"
)

type Server struct {
	Conn sql.Connection
	Conf config.Config
}

type db struct {
	Conn sql.Connection
}

func (server *Server) ListenAndServe() error {
	log.Printf("server starting on address %s\n", server.Conf.Address())
	return http.ListenAndServe(server.Conf.Address(), nil)
}
func New(conf config.Config, conn sql.Connection) *Server {
	http.HandleFunc("/", root)
	http.HandleFunc("/signin", signin)
	http.HandleFunc("/create-account", createAccount)
	http.HandleFunc("/favicon.ico", favicon)
	return &Server{Conf: conf, Conn: conn}
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
	<a href="signin/">Sign In</a>
	<a href="create-account/">Create Account</a>
	`))
}

func favicon(w http.ResponseWriter, r *http.Request) {
	faviconName := path.Base(r.URL.Path)
	http.ServeFile(w, r, "/static/"+faviconName)
}
