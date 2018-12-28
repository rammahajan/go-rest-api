FROM golang:latest
RUN mkdir /app
ADD ./src /app/
WORKDIR /app
RUN go get "github.com/gorilla/mux"
RUN go build -o main .
CMD ["/app/main"]