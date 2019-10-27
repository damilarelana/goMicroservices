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
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure()) // Connect to the MathsService
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

// sortHandler endpoint focus on the sum service that gives the summation of the Array elements
func sortHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json: charset=UFT-8") // set the content header type
	vars := mux.Vars(r)                                               // extract usable information from request object by parsing the inputs x and y
	inputArray := strToArray(vars["array"])                           // extract value of array from the variable request arguments
	ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)   // initialize resources context
	defer cancel()
	req := &ms.SortRequest{ // initialize the request struct
		Array: inputArray,
	}
	g := gRPCClient() // call the initialized client
	resp, err := g.Sort(ctx, req)
	if err == nil {
		msg := fmt.Sprintf("Sorted array is %f:", resp.SortedArray)
		json.NewEncoder(w).Encode(msg)
	} else {
		log.Fatal(errors.Wrap(err, fmt.Sprintf("Internal Server error: sortHandler")))
	}
}

// sumHandler endpoint focus on the sum service that gives the summation of the Array elements
func sumHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json: charset=UFT-8") // set the content header type
	vars := mux.Vars(r)                                               // extract usable information from request object by parsing the inputs x and y
	inputArray := strToArray(vars["array"])                           // extract value of array from the variable request arguments
	ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)   // initialize resources context
	defer cancel()
	req := &ms.SumRequest{ // initialize the request struct
		Array: inputArray,
	}
	g := gRPCClient() // call the initialized client
	resp, err := g.Sum(ctx, req)
	if err == nil {
		msg := fmt.Sprintf("Minimum element is %f:", resp.ArrayValuesSum)
		json.NewEncoder(w).Encode(msg)
	} else {
		log.Fatal(errors.Wrap(err, fmt.Sprintf("Internal Server error: sumHandler")))
	}
}

// minHandler endpoint focus on the Min service that finds the minimum value amongst Array elements
func minHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json: charset=UFT-8") // set the content header type
	vars := mux.Vars(r)                                               // extract usable information from request object by parsing the inputs x and y
	inputArray := strToArray(vars["array"])                           // extract value of array from the variable request arguments
	ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)   // initialize resources context
	defer cancel()
	req := &ms.MinRequest{ // initialize the request struct
		Array: inputArray,
	}
	g := gRPCClient() // call the initialized client
	resp, err := g.Min(ctx, req)
	if err == nil {
		msg := fmt.Sprintf("Minimum element is %f:", resp.Minimum)
		json.NewEncoder(w).Encode(msg)
	} else {
		log.Fatal(errors.Wrap(err, fmt.Sprintf("Internal Server error: minHandler")))
	}
}

// maxHandler endpoint focus on the Max service that finds the maximum value amongst Array elements
func maxHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json: charset=UFT-8") // set the content header type
	vars := mux.Vars(r)                                               // extract usable information from request object by parsing the inputs x and y
	inputArray := strToArray(vars["array"])                           // extract value of array from the variable request arguments
	ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)   // initialize resources context
	defer cancel()
	req := &ms.MaxRequest{ // initialize the request struct
		Array: inputArray,
	}
	g := gRPCClient() // call the initialized client
	resp, err := g.Max(ctx, req)
	if err == nil {
		msg := fmt.Sprintf("Maximum element is %f:", resp.Maximum)
		json.NewEncoder(w).Encode(msg)
	} else {
		log.Fatal(errors.Wrap(err, fmt.Sprintf("Internal Server error: maxHandler")))
	}
}

// averageHandler endpoint focus on the Average service that finds the average value amongst Array elements
func averageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json: charset=UFT-8") // set the content header type
	vars := mux.Vars(r)                                               // extract usable information from request object by parsing the inputs x and y
	inputArray := strToArray(vars["array"])                           // extract value of array from the variable request arguments
	ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)   // initialize resources context
	defer cancel()
	req := &ms.AverageRequest{ // initialize the request struct
		Array: inputArray,
	}
	g := gRPCClient() // call the initialized client
	resp, err := g.Average(ctx, req)
	if err == nil {
		msg := fmt.Sprintf("Average is %f:", resp.Average)
		json.NewEncoder(w).Encode(msg)
	} else {
		log.Fatal(errors.Wrap(err, fmt.Sprintf("Internal Server error: averageHandler")))
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
		log.Fatal(errors.Wrap(err, fmt.Sprintf("Internal Server error: addHandler")))
	}
}

func main() {
	muxRouter := mux.NewRouter().StrictSlash(true)                     // instantiate the gorillamux Router and enforce trailing slash rule i.e. `/path` === `/path/`
	muxRouter.NotFoundHandler = http.HandlerFunc(custom404PageHandler) // customer 404 Page handler scenario
	muxRouter.HandleFunc("/", msAPIHomePage)
	muxRouter.HandleFunc("/add/{x}/{y}", addHandler).Methods("GET")         // the add service endpoint mapping
	muxRouter.HandleFunc("/average/{array}", averageHandler).Methods("GET") // the average service endpoint mapping
	muxRouter.HandleFunc("/max/{array}", maxHandler).Methods("GET")         // the maximum service endpoint mapping
	muxRouter.HandleFunc("/min/{array}", minHandler).Methods("GET")         // the minimum service endpoint mapping
	muxRouter.HandleFunc("/sum/{array}", sumHandler).Methods("GET")         // the summation service endpoint mapping
	muxRouter.HandleFunc("/sort/{array}", sortHandler).Methods("GET")       // the sorting service endpoint mapping
	fmt.Println("API is up and running at http://127.0.0.1:8080")
	for {
		// log.Fatal(http.ListenAndServe(":8080", muxRouter)) // set the port where the http server listens and serves the API from
		log.Fatal(errors.Wrap(http.ListenAndServe(":8080", muxRouter), "Failed to start gRPC Server"))
	}
}
