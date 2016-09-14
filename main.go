package main

import (
	"github.com/gorilla/mux"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
        "github.com/docker/go-plugins-helpers/ipam"
	"github.com/vorchid63/vzipam/ipam"
)


type vzIpam struct {
    
     driver    ipam
}

func main() {

       // declare cli command arguments
	var flagDebug = cli.BoolFlag{
		Name:  "debug, d",
		Usage: "enable debugging",
	}
 
	app := cli.NewApp()
	app.Name = "vzipam"
	app.Usage = "Docker vzIpam Networking"
	app.Version = version
	app.Flags = []cli.Flag{
		flagDebug
	}
	app.Action = Run
	app.Run(os.Args)

}


// Run initializes the driver
func Run(ctx *cli.Context) {

	if ctx.Bool("debug") {
		log.SetLevel(log.DebugLevel)
	}

	ipam := vzIpam{}

	handler := driver.NewHandler(ipam)

	handler.ServeTCP("vzipam", ":8080")
}