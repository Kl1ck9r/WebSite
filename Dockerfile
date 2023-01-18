FROM golang:latest 

WORKDIR /WebApplication

COPY go.mod ./
COPY go.sum ./

RUN go mod download 

COPY *.go ./internal  ./ 

EXPOSE 8080

CMD [ "/web-application-golang" ]
