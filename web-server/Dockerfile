FROM golang:1.22

WORKDIR /app

COPY go.mod *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /web-server

EXPOSE 8080
CMD [ "/web-server" ]
