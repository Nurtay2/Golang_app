# Name this stage as 'builder'
FROM golang:1.21.0 as builder
WORKDIR /usr/src
COPY . .
RUN go mod download

EXPOSE 8080

CMD ["go", "run", "/usr/src/cmd/ufc", "."]