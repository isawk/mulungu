package mulungu

import (
	"net/http"
	"strings"

	"github.com/edgedagency/mulungu/middleware"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

//AppEngineServer representation of Google App Engine Server
var AppEngineServer *AppEngine

func init() {
	AppEngineServer = NewAppEngine()
}

//AppEngine Google appengine server configurations
type AppEngine struct {
	router *mux.Router
	chain  alice.Chain
}

//NewAppEngine create a new appengine server
func NewAppEngine() *AppEngine {
	return &AppEngine{router: mux.NewRouter(), chain: alice.New(middleware.Logging)}
}

//Start sets up http handler with register handlers
func (s *AppEngine) Start() {
	http.Handle("/", s.chain.Then(s.router))
}

//Middleware registers middlewares
func (s *AppEngine) Middleware(middlwares ...alice.Constructor) {
	s.chain = s.chain.Append(middlwares...)
}

//Handler can be used to register a handler, handlers process information based on a path signature
func (s *AppEngine) Handler(path string, h http.Handler) *mux.Route {
	return s.router.Handle(path, h)
}

//HandlerFunc can be used to register a handler, handlers process information based on a path signature
func (s *AppEngine) HandlerFunc(path string, f func(w http.ResponseWriter, r *http.Request)) *mux.Route {
	return s.router.HandleFunc(path, f)
}

//Context returns context from request
func (s *AppEngine) Context(r *http.Request) context.Context {
	return appengine.NewContext(r)
}

//AppEngineServiceURL this returns an app spot host
func AppEngineServiceURL(host, service, version string) string {
	var hostURL = strings.Join([]string{service, host}, "-dot-")
	//e.g. https://v10032017t163649-dot-application-dot-ibudo-console.appspot.com/
	if version != "" {
		hostURL = strings.Join([]string{version, service, host}, "-dot-")
	}

	return hostURL
}