package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tonglil/slack-here-reminder/function"
)

var (
	port string
)

func main() {
	function.MonitoredChannels()

	port = os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	http.HandleFunc("/here", function.HereHandler)
	log.Printf("listening on :%s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
