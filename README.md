### goMicroservice

A gRPC based microservices implementation in Go, deployed in a kubernetes cluster and consisting of two microservices:

* `Maths` service
* `REST API` to expose the `Maths` service

***

The `Maths` microservice that:

* computes the `Addition` of two numbers
* computes the `Average` of an array of numbers
* computes the `Maximum` of an array of numbers
* computes the `Minimum` of an array of numbers
* computes the `Summation` of an array of numbers
* computes the `Sorted` form of an array of numbers

***

A `REST API` that responds to the client and routes them to the corresponding `Maths` compute

***

To start the application, you would need to enter the downloaded project directory, and then run the server and api in separate terminals respectively i.e.

* for the `server`, execute `go run server/main.go`
* for the `api`, execute `go run api/main.go`