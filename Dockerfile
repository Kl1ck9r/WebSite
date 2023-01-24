FROM golang:latest 

WORKDIR /WebApplication

COPY go.mod ./WebApplication
COPY go.sum ./WebApplication

RUN go mod download 

COPY . ./WebApplication

EXPOSE 8080

CMD [ "./cmd/main.go" ]
