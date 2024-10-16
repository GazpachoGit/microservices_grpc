FROM golang:1.23 AS builder
## create a folder in the golang image
WORKDIR /usr/src/app
## copy the project to the golang image
COPY . .

## build binary with name 'payment'
RUN go build -o payment ./cmd/main.go

##create new empty image
FROM scratch
## copy the binary to empty image. There will be only the binary file(without any folder)
COPY --from=builder /usr/src/app/payment ./payment

##run the payment binary(when container is started)
CMD ["./payment"]

##docker build . -t payment:v1