FROM golang:1.17

LABEL  maintainer="daniehhu@stud.ntnu.no"

WORKDIR /go/src/app/cmd

COPY ./cmd /go/src/app/cmd
COPY ./constants /go/src/app/constants
COPY ./endpoints /go/src/app/endpoints
COPY ./globals /go/src/app/globals
COPY ./serviceKey /go/src/app/serviceKey
COPY ./go.mod /go/src/app/go.mod
COPY ./go.sum /go/src/app/go.sum


RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o server

CMD ["./cmd"]
