package apiserver

import (
	"encoding/json"
	"git.array2d.com/cncf/tca/pkg/tca"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetkindHandler() (kh *kindHandler) {
	kh = &kindHandler{
		tca: tca.New(),
	}
	return
}

type kindHandler struct {
	tca *tca.Tca
}

func (kh *kindHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	mainKind := mux.Vars(request)["kind"]
	_kindids := request.URL.Query()
	kindids := make(map[string]string)
	for k, ids := range _kindids {
		if len(ids) > 0 {
			kindids[k] = ids[0]
		}
	}
	log.WithFields(
		log.Fields{
			"method":   request.Method,
			"mainkind": mainKind,
			"query":    request.URL.Query(),
		}).Debugln("req")

	var values map[string]any
	json.NewDecoder(request.Body).Decode(&values)
	code, stdouterr := kh.tca.Method(mainKind, request.Method, kindids, values)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	writer.Write([]byte(stdouterr))
}
