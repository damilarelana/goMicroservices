FROM golang:1.12.7-buster

# Download and install the latest release of dep
# ADD https://github.com/golang/dep/releases/download/v0.5.4/dep-linux-amd64 /usr/local/bin/dep
# RUN chmod +x /usr/local/bin/dep

# create application directory in the container
RUN mkdir -p /go/src/server/

# copy the code into application directory in the container
ADD . /go/src/server/

# designate the working directory within the container
WORKDIR /go/src/server/

# Download and install dependencies e.g. github.com/gorilla/mux etc.
# RUN dep ensure -vendor-only

# Install dependencies
RUN go get -v ./...

# Expose the default port
EXPOSE 9090

RUN go build -o main ./...
CMD ["/go/src/server/main"]