package apiserver

import (
	"encoding/json"
	"git.array2d.com/cncf/tca/pkg/tca"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type kindHandler struct {
	tca *tca.Tca
}

func (k *kindHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	kind := mux.Vars(request)["kind"]
	id := mux.Vars(request)["id"]
	log.WithFields(
		log.Fields{
			"method": request.Method,
			"kind":   kind,
			"id":     id,
		}).Debugln("req")
	var values map[string]any
	var err error
	err = json.NewDecoder(request.Body).Decode(&values)
	if err != nil {
		log.WithFields(
			log.Fields{
				"method": request.Method,
				"kind":   kind,
				"id":     id,
				"decode": id,
			}).Errorln()
		return
	}
	code, stdouterr := k.tca.Method(kind, request.Method, id, values)
	writer.WriteHeader(code)
	writer.Write([]byte(stdouterr))
}
