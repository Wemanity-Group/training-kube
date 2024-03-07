FROM golang:1.16 AS build

WORKDIR /app

COPY main.go go.mod /app/

RUN go mod download github.com/gorilla/mux
RUN go get github.com/gorilla/handlers

RUN go build -o hello-app

FROM debian:buster-slim

COPY --from=build /app/hello-app /usr/local/bin/

EXPOSE 8080

CMD ["hello-app"]
