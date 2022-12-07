FROM golang:1.19-alpine
RUN mkdir /app
ADD go.mod go.sum . /app/

WORKDIR /app
RUN go build -o main.go .
CMD ["/app/main"]