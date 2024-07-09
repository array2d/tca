package apiserver

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Apiserver struct {
	Addr   string
	Prefix string
	router *mux.Router
}

func New() (a *Apiserver) {
	a = &Apiserver{
		router: NewRouter(),
	}
	a.Route("/{kind}", GetkindHandler())
	return
}
func (a *Apiserver) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	a.router.ServeHTTP(writer, request)
}
func (a *Apiserver) Route(path string, handler http.Handler) (err error) {
	err = a.router.PathPrefix(a.Prefix).Path(path).Handler(handler).GetError()
	if err != nil {
		log.WithFields(
			log.Fields{
				"err":  err,
				"path": path,
			}).Errorln("router register failed")
	} else {
		log.WithFields(
			log.Fields{
				"path": path,
			}).Infoln("router register")
	}
	return
}
