FROM golang:1.10.1-alpine3.7
EXPOSE 8080
COPY ./webserver.go ./
COPY ./index.html ./
RUN go build -o ./webserver.go
USER nobody
ENTRYPOINT ["./webserver.go"]

