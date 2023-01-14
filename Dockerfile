FROM golang:latest 

WORKDIR /webApp

COPY go.mod ./
COPY go.sum ./

RUN go mod download 

COPY *.go ./internal  ./ 

RUN go build -o /web-application-golang

EXPOSE 8080

CMD [ "/web-application-golang" ]
