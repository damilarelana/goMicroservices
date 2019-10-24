### goMicroservice

A gRPC based microservices implementation in Go, deployed in a kubernetes cluster and consisting of two microservices:

* `Maths` service
* public `REST API`

The `Maths` microservice that:

* computes the `Addition` of two numbers
* computes the `Average` of an array of numbers
* computes the `Maximum` of an array of numbers
* computes the `Minimum` of an array of numbers
* computes the `Summation` of an array of numbers
* computes the `Sorted` form of an array of numbers

A `REST API` that responds to the client and routes them.