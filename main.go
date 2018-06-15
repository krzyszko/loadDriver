package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/apex/log"
	"github.com/krzyszko/loaddriver/config"
	"github.com/krzyszko/loaddriver/plan"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "", "path to config file")
	flag.Parse()
	con, err := config.GetComponents(configFile)
	if err != nil {
		log.Errorf("%s", err)
		os.Exit(100)
	}
	for _, com := range con.Components {
		fmt.Printf("%s", config.GetComponentParams(com.Params))
	}
	plan := &plan.Plan{}
	plan.Run()
}
