package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"
	"github.com/rajatjindal/example/pkg/api"
	"github.com/sirupsen/logrus"
)

var gorillamuxAdapter *gorillamux.GorillaMuxAdapter

func main() {
	s, err := api.New()
	if err != nil {
		logrus.Fatal(err)
	}

	serveUsingRegular(s.Router)
}

func serveUsingRegular(r *mux.Router) {
	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         ":8080",
		Handler:      r,
	}

	fmt.Printf("Starting HTTP server on %s\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		logrus.Fatalf("server.ListendAndServe() failed with %s", err)
	}
}

func serveUsingLambda(r *mux.Router) {
	gorillamuxAdapter = gorillamux.New(r)
	lambda.Start(Handler)
}

//Handler handles the events
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return gorillamuxAdapter.ProxyWithContext(ctx, req)
}
