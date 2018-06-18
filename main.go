package main

import (
	"flag"
	"os"

	"github.com/apex/log"
	"github.com/krzyszko/loaddriver/config"
	"github.com/krzyszko/loaddriver/loader"
	"github.com/krzyszko/loaddriver/plan"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "", "path to config file")
	flag.Parse()
	con, err := config.GetConfig(configFile)
	if err != nil {
		log.Errorf("%s", err)
		os.Exit(100)
	}
	plan := &plan.Plan{}
	components, err := loader.ComponentsFromConfiguration(con)
	if err != nil {
		log.Errorf("%s", err)
	}
	for _, cmpt := range components {
		plan.AddComponent(cmpt)
	}

	plan.Run()
}
