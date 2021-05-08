package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/kelseyhightower/envconfig"
)

type envConfig struct {
	Port int `envconfig:"PORT" default:"8080" required:"true"`
}

func main() {
	var env envConfig
	if err := envconfig.Process("", &env); err != nil {
		log.Fatalf("failed to process env var: %s", err)
	}

	handler := new(Handler)
	myport := ":" + strconv.Itoa(env.Port)
	log.Printf("Server starting on port :%s\n", myport)
	if err := http.ListenAndServe(myport, handler); err != nil {
		log.Fatalf("failed to start server, %s", err.Error())
	}
}

/*

http://colorgrid.default.d2k.n3wscott.com/?url=http://colors.default.d2k.n3wscott.com&count=25

*/
