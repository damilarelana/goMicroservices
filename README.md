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

***

*Example*

Let's test with integer `17`, integer `18` and array `[48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17]`, via the sample requests:

```
127.0.0.1:8080/add/17/18
127.0.0.1:8080/average/[48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17]
127.0.0.1:8080/max/[48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17]
127.0.0.1:8080/min/[48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17]
127.0.0.1:8080/sum/[48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17]
127.0.0.1:8080/sort/[48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17]
```

This gives us the following results:

```
Addition is:  35
================
Average is:  55.8125
================
Maximum element is:  97
================
Minimum element is:  9
================
Sum total of the elements is:  893
================
Sorted array is:  [9 17 19 27 34 37 48 57 63 68 70 82 83 86 96 97]
================
```
