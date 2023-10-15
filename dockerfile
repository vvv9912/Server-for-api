FROM golang:alpine AS builder

WORKDIR /app


#ADD go.mod .

COPY . .

RUN go mod tidy
RUN go build -o serverBot2 


CMD ["./serverBot2"]


