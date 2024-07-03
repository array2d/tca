package cmd

import (
	apiserver2 "git.array2d.com/cncf/tca/pkg/apiserver"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"
)

var apiserver = &cobra.Command{
	Use:   "apiserver [port]",
	Short: "apiserver for your control request",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var addr string = ":4000"
		if len(args) > 0 {
			addr = args[0]
		}
		a := apiserver2.New()
		log.WithFields(
			log.Fields{
				"addr": addr,
			}).Infoln("httpserver starting")
		err := http.ListenAndServe(addr, a)
		if err != nil {
			log.WithFields(
				log.Fields{
					"err":  err,
					"addr": addr,
				}).Errorln("httpserver start fail")
		}
	},
}
