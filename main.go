package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rajatjindal/example/pkg/api"
	"github.com/sirupsen/logrus"
)

func main() {
	s, err := api.New()
	if err != nil {
		logrus.Fatal(err)
	}

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         ":8080",
		Handler:      s.Router,
	}

	fmt.Printf("Starting HTTP server on %s\n", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		logrus.Fatalf("server.ListendAndServe() failed with %s", err)
	}
}
