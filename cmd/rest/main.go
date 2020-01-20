package main

import (
	"fmt"
	"github.com/dgparker/wing-chun/pkg/http/rest"
	"log"
	"os"
)

const (
	// PORT is the env value for wing-chun server port
	PORT = "WING_CHUN_PORT"
	// DEFAULTPORT is the default env value for wing-chun server port
	DEFAULTPORT = "9000"
)

func main() {
	port := os.Getenv(PORT)
	if port == "" {
		port = DEFAULTPORT
	}

	api := rest.New()

	api.Addr = fmt.Sprintf(":%s", port)
	log.Printf("server listening on port :%s", port)
	log.Fatal(api.ListenAndServe())
}
