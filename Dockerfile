# syntax=docker/dockerfile:1

FROM golang:1.19-alpine


EXPOSE 8080

CMD [" /web-application/notes" ]