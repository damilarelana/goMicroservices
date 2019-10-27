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
func msAPIHomePage(w http.ResponseWriter, r *http.Request) {
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

// gRPCClient is used by the API to connect to the Microservice, when an API call is made to the API
func gRPCClient() ms.MathServiceClient {
	conn, err := grpc.Dial("math-service:9090", grpc.WithInsecure()) // Connect to the MathsService
	if err != nil {
		log.Fatal(errors.Wrap(err, fmt.Sprintf("Dial failed")))
	}
	msClient := ms.NewMathServiceClient(conn)
	return msClient
}

// strToInt takes in a url variable that is a string representation of an int,
func strToInt(str string) (num int64) {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Fatal(errors.Wrap(err, fmt.Sprintf("Invalid parameter conversion to int")))
	}
	return num
}

// strToArray takes in a url variable that is a string representation of an array,
func strToArray(str string) (array []float64) {
	byteData := []byte(str)
	err := json.Unmarshal(byteData, &array)
	if err != nil {
		log.Fatal(errors.Wrap(err, fmt.Sprintf("Unable to convert string to array")))
	}
	return array
}

// averageHandler endpoint focus on the Average service that finds the average value of an Array element
func averageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json: charset=UFT-8") // set the content header type
	vars := mux.Vars(r)                                               // extract usable information from request object by parsing the inputs x and y
	x := strToInt(vars["x"])                                          // extract value of x from the variable request arguments
	y := strToInt(vars["y"])                                          // extract value of y from the variale request arguments
	ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)   // initialize resources context
	defer cancel()
	req := &ms.AddRequest{ // initialize the request struct
		X: x,
		Y: y,
	}
	g := gRPCClient() // call the initialized client
	resp, err := g.Add(ctx, req)
	if err == nil {
		msg := fmt.Sprintf("Addition is %d:", resp.Addition)
		json.NewEncoder(w).Encode(msg)
	} else {
		log.Fatal(errors.Wrap(err, fmt.Sprintf("Internal Server error")))
	}
}

// addHandler endpoint focus on the Add service that sums up x and y, to give z
func addHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json: charset=UFT-8") // set the content header type
	vars := mux.Vars(r)                                               // extract usable information from request object by parsing the inputs x and y
	x := strToInt(vars["x"])                                          // extract value of x from the variable request arguments
	y := strToInt(vars["y"])                                          // extract value of y from the variale request arguments
	ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)   // initialize resources context
	defer cancel()
	req := &ms.AddRequest{ // initialize the request struct
		X: x,
		Y: y,
	}
	g := gRPCClient() // call the initialized client
	resp, err := g.Add(ctx, req)
	if err == nil {
		msg := fmt.Sprintf("Addition is %d:", resp.Addition)
		json.NewEncoder(w).Encode(msg)
	} else {
		log.Fatal(errors.Wrap(err, fmt.Sprintf("Internal Server error")))
	}
}

// serviceRequestHandler() defines request handling service [used to aggregate all endpoints before running]
func serviceRequestHandlers() {
	muxRouter := mux.NewRouter().StrictSlash(true)                     // instantiate the gorillamux Router and enforce trailing slash rule i.e. `/path` === `/path/`
	muxRouter.NotFoundHandler = http.HandlerFunc(custom404PageHandler) // customer 404 Page handler scenario
	muxRouter.HandleFunc("/", msAPIHomePage)
	muxRouter.HandleFunc("/{x}/{y}", addHandler).Methods("GET")     // the add service endpoint mapping
	muxRouter.HandleFunc("/{array}", averageHandler).Methods("GET") // the average service endpoint mapping
	fmt.Println("API is up an running now at : 8080")
	log.Fatal(http.ListenAndServe(":8080", muxRouter)) // set the port where the http server listens and serves the API from
}

func main() {
	go serviceRequestHandlers() // call and run the server as a goroutine

	// create an artificial pause "to ensure the main function goroutine does not cause the serviceRequestHandler goroutine to exit"
	var tempString string
	fmt.Scanln(&tempString)
}
