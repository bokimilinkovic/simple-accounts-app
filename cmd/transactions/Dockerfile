## We'll choose the incredibly lightweight
## Go alpine image to work with
FROM golang AS builder
ENV GO111MODULE=on

## We create an /app directory in which
## we'll put all of our project code
RUN mkdir /app
ADD . /app

WORKDIR /app
RUN apt update && apt install ca-certificates libgnutls30 -y

RUN go mod download

## We want to build our application's binary executable
RUN cd cmd/transactions/ && CGO_ENABLED=0 GOOS=linux go build -o transactions

## the lightweight scratch image we'll
## run our application within
FROM alpine:latest AS production
## We have to copy the output from our
## builder stage to our production stage
COPY --from=builder /app/cmd/transactions/ .
## we can then kick off our newly compiled
## binary exectuable!!
CMD ["./transactions"]