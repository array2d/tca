package main

import (
	"git.array2d.com/cncf/tca/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {
	// 执行命令

	if err := cmd.RootCmd.Execute(); err != nil {
		log.Errorln(err)
	}
}
