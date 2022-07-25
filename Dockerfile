#get a base image
FROM golang:1.17-buster

#MAINTAINER anaiya raisinghani <anaiya.raisinghani@mongodb.com>

WORKDIR /go/src/app
COPY . .

RUN go get -d -v
RUN go build -v
RUN go get -d "github.com/99designs/gqlgen@v0.17.13"
RUN go run github.com/99designs/gqlgen generate 

CMD ["./go-graphql"]
