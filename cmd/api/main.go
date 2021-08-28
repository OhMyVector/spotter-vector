package main

import (
	"flag"

	"github.com/OhMyVector/spotter-vector/pkg/api"
	"github.com/OhMyVector/spotter-vector/pkg/core/config"
)

func main() {

	cfgPath := flag.String("p", "conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	checkErr(err)

	checkErr(api.Start(cfg))
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
