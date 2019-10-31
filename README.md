### goMicroservice

A gRPC based microservices implementation in Go, deployed in a kubernetes cluster and consisting of two microservices:

* `Maths` service
* `REST API` to expose the `Maths` service

The `Maths` microservice that:

* computes the `Addition` of two numbers
* computes the `Average` of an array of numbers
* computes the `Maximum` of an array of numbers
* computes the `Minimum` of an array of numbers
* computes the `Summation` of an array of numbers
* computes the `Sorted` form of an array of numbers

A `REST API` that responds to the client and routes them to the corresponding `Maths` compute

***

To start the application without deploying to `kubernetes`, we need to enter the downloaded project directory, and run the server and api in separate terminals respectively i.e.

* for the `server`, execute `go run server/main.go`
* for the `api`, execute `go run api/main.go`

Now test with integer `17`, integer `18` and array `[48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17]`, via the sample api requests:

```
127.0.0.1:8080/add/17/18
127.0.0.1:8080/average/[48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17]
127.0.0.1:8080/max/[48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17]
127.0.0.1:8080/min/[48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17]
127.0.0.1:8080/sum/[48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17]
127.0.0.1:8080/sort/[48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17]
```


***

To deploy and start the application in `kubernetes`, we need to make a few changes to ensure the `api` is able to communicate with the `math-service` server. These changes to be implemented from within the project directory are:

* start `kubernetes` by running `minikube start`
* ensure docker reuses the Docker deamon in minikube, by running `eval $(minikube docker-env)`
* update `api/main.go` by changing `localhost:9090` to `math-service:9090` on line 35
* rebuild `api` docker image by running `docker build -t math-service-api -f Dockerfile.api .`
* rebuild `server` docker image by running `docker build -t math-service-server -f Dockerfile.server .`
* clean-up previous `kubernetes` deployments, by running `kubectl delete all --all`
* deploy the `api` to kubernetes by running `kubectl create -f api.yaml`
* deploy the `server` to kubernetes by running `kubectl create -f server.yaml`
* copy the `api`'s new kubernetes ip address after running `minikube service api-service --url`

Now test with integer `17`, integer `18` and array `[48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17]`, via the sample api requests:

```
curl <ip-address:port>/add/17/18
curl <ip-address:port>/average/[48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17]
curl <ip-address:port>/max/[48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17]
curl <ip-address:port>/min/[48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17]
curl <ip-address:port>/sum/[48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17]
curl <ip-address:port>/sort/[48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17]
```
where `<http://ip-address:port>` represents the result of running `minikube service api-service --url`. For example, if `<http://ip-address:port>` was `http://192.168.99.100:32348`, then it means the first api request is `curl http://192.168.99.100:32348/add/17/18`. For the `array` parameters via a browser i.e. simply copy and paste 

*** 

In both scenarios of runing the application (i.e. with or without kubernetes), the test examples should give us results numbers similar to:

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
