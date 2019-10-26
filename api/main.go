package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	ms "github.com/damilarelana/goMicroservice/mathService"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// mathServiceAPIHomePage defines the landing page for the API
func mathServiceAPIHomePage(w http.ResponseWriter, r *http.Request) {
	dataHomePage := "Endpoint: homepage"
	io.WriteString(w, dataHomePage)
}

// custom404PageHandler defines custom 404 page
func custom404PageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")        // set the content header type
	w.WriteHeader(http.StatusNotFound)                 // this automatically generates a 404 status code
	data404Page := "This page does not exist ... 404!" // page content
	io.WriteString(w, data404Page)
}

func gRPCClient() MathServiceClient {
	// Connect to the MathsService
	conn, err := grpc.Dial("math-service:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatal(errors.Wrap(err, fmt.Sprintf("Dial failed")))
	}
	msClient := ms.NewMathServiceClient(conn)
	return msClient
}

// addHandler endpoint focus on the Add service to add x and y
func addHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json: charset=UFT-8") // set the content header type
	vars := mux.Vars(r)                                               // extract usable information from request object by parsing the inputs x and y

	x, err := strconv.ParseInt(vars["x"], 10, 64) // extract value of x from the variale request arguments
	if err != nil {
		log.Fatal(errors.Wrap(err, fmt.Sprintf("Invalid parameter x")))
	}

	y, err := strconv.ParseInt(vars["y"], 10, 64) // extract value of y from the variale request arguments
	if err != nil {
		log.Fatal(errors.Wrap(err, fmt.Sprintf("Invalid parameter y")))
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Minute) // initialize a call Add service
	defer cancel()

	req := &ms.AddRequest{ // initialize the request struct
		X: x,
		Y: y,
	}

	g := gRPCClient() // call the initialized client
	resp, err := g.Add(ctx, req)
	if err == nil {
		msg := fmt.Sprintf("Addition is %d:", resp.Result)
		json.NewEncoder(w).Encode(msg)
	} else {
		log.Fatal(errors.Wrap(err, fmt.Sprintf("Internal Server error")))
	}
}

// serviceRequestHandler() defines request handling service [used to aggregate all endpoints before running]
func serviceRequestHandlers() {
	muxRouter := mux.NewRouter().StrictSlash(true)                     // instantiate the gorillamux Router and enforce trailing slash rule i.e. `/path` === `/path/`
	muxRouter.NotFoundHandler = http.HandlerFunc(custom404PageHandler) // customer 404 Page handler scenario
	muxRouter.HandleFunc("/", mathServiceAPIHomePage)
	muxRouter.HandleFunc("/{x}/{y}", addHandler).Methods("GET") // responds to DELETE request to remove an article with a specific identifier
	fmt.Println("API Endpoint is up an running now at : 8080")
	log.Fatal(http.ListenAndServe(":8080", muxRouter)) // set the port where the http server listens and serves. changed `nil` to the instance muxRouter
}

func main() {
	go serviceRequestHandlers() // call and run the server as a goroutine

	// create an artificial pause "to ensure the main function goroutine does not cause the serviceRequestHandler goroutine to exit"
	var tempString string
	fmt.Scanln(&tempString)
}
